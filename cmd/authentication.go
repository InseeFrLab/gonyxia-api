package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	Email  string   `json:"email,omitempty"`
	Name   string   `json:"name,omitempty"`
	Groups []string `json:"groups,omitempty"`
	ID     string   `json:"preferred_username,omitempty"`
}

func AuthMiddleware(ctx context.Context, verifier *oidc.IDTokenVerifier) gin.HandlerFunc {
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
			c.Set("claims", IDTokenClaims)
		}

		c.Next()
	}
}
