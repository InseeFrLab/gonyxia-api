package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	oidc "github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine/log"
)

type Claims struct {
	Sub  string `json:"sub,omitempty"`
	Name string `json:"name,omitempty"`
}

func authMiddleware(ctx context.Context, verifier *oidc.IDTokenVerifier) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !strings.HasPrefix(c.Request.URL.Path, "/public") {
			tokenHeader := strings.TrimPrefix(c.Request.Header.Get("Authorization"), "Bearer ")
			fmt.Printf("Authorization %s", tokenHeader)
			fmt.Println()
			token, err := verifier.Verify(ctx, tokenHeader)
			if err != nil {
				panic(err)
			}
			var IDTokenClaims Claims // ID Token payload is just JSON.
			if err := token.Claims(&IDTokenClaims); err != nil {
				panic(err)
			}
			fmt.Println(IDTokenClaims)
		}

		c.Next()
	}
}

func main() {
	loadConfiguration()
	r := gin.Default()

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
			log.Warningf(ctx, "Token audience validation disabled")
			oidcConfig.SkipClientIDCheck = true
		}
		verifier := provider.Verifier(oidcConfig)
		r.Use(authMiddleware(ctx, verifier))
	}
	r.GET("/ping", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
