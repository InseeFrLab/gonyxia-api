package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	oidc "github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
	cmd "github.com/inseefrlab/onyxia-admin/cmd"
	_ "github.com/inseefrlab/onyxia-admin/docs"
	"github.com/inseefrlab/onyxia-admin/internal/kubernetes"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

// gin-swagger middleware
// swagger embed files

func main() {
	loadConfiguration()
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	zap.ReplaceGlobals(zap.Must(zap.NewProduction()))

	if config.Authentication.IssuerURI != "" {
		fmt.Printf("Using authentication with issuer %s", config.Authentication.IssuerURI)
		fmt.Println()
		client := &http.Client{
			Timeout: time.Duration(6000) * time.Second,
		}
		ctx := oidc.ClientContext(context.Background(), client)
		provider, _ := oidc.NewProvider(ctx, config.Authentication.IssuerURI)
		oidcConfig := &oidc.Config{}
		if config.Authentication.Audience != "" {
			oidcConfig.ClientID = config.Authentication.Audience
		} else {
			zap.L().Warn("Token audience validation disabled")
			oidcConfig.SkipClientIDCheck = true
		}
		verifier := provider.Verifier(oidcConfig)
		r.Use(cmd.AuthMiddleware(ctx, verifier))
	}

	kubernetes.InitClient()

	cmd.RegisterHandlers(r)
	r.Run()
}
