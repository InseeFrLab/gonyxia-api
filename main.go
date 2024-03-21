package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	oidc "github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	cmd "github.com/inseefrlab/onyxia-api/cmd"
	_ "github.com/inseefrlab/onyxia-api/docs"
	configuration "github.com/inseefrlab/onyxia-api/internal/configuration"
	"github.com/inseefrlab/onyxia-api/internal/kubernetes"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

var (
	Version = 2
)

// gin-swagger middleware
// swagger embed files

func main() {
	configuration.LoadConfiguration()
	r := gin.Default()
	baseRoutes := r.Group(configuration.Config.RootPath)
	baseRoutes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	privateRoutes := baseRoutes.Group("/")
	publicRoutes := baseRoutes.Group("/public")

	zap.ReplaceGlobals(zap.Must(zap.NewProduction()))

	if strings.EqualFold(configuration.Config.Authentication.Mode, "openidconnect") {
		fmt.Printf("Using OIDC authentication with issuer %s", configuration.Config.OIDC.IssuerURI)
		fmt.Println()
		client := &http.Client{
			Timeout: time.Duration(6000) * time.Second,
		}
		ctx := oidc.ClientContext(context.Background(), client)
		provider, _ := oidc.NewProvider(ctx, configuration.Config.OIDC.IssuerURI)
		oidcConfig := &oidc.Config{}
		if configuration.Config.OIDC.Audience != "" {
			oidcConfig.ClientID = configuration.Config.OIDC.Audience
		} else {
			zap.L().Warn("Token audience validation disabled")
			oidcConfig.SkipClientIDCheck = true
		}
		verifier := provider.Verifier(oidcConfig)
		privateRoutes.Use(cmd.AuthMiddleware(ctx, verifier))
	}

	kubernetes.InitClient()

	cmd.RegisterPrivateHandlers(privateRoutes)
	cmd.RegisterPublicHandlers(publicRoutes)

	if configuration.Config.Security.CORS.AllowedOrigins != "" {
		r.Use(cors.New(cors.Config{
			AllowOrigins:     []string{configuration.Config.Security.CORS.AllowedOrigins},
			AllowMethods:     []string{"*"},
			AllowHeaders:     []string{"Origin"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))
	}
	r.Run()
}
