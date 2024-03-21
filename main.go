package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	oidc "github.com/coreos/go-oidc/v3/oidc"
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

	if configuration.Config.Authentication.IssuerURI != "" {
		fmt.Printf("Using authentication with issuer %s", configuration.Config.Authentication.IssuerURI)
		fmt.Println()
		client := &http.Client{
			Timeout: time.Duration(6000) * time.Second,
		}
		ctx := oidc.ClientContext(context.Background(), client)
		provider, _ := oidc.NewProvider(ctx, configuration.Config.Authentication.IssuerURI)
		oidcConfig := &oidc.Config{}
		if configuration.Config.Authentication.Audience != "" {
			oidcConfig.ClientID = configuration.Config.Authentication.Audience
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
	r.Run()
}
