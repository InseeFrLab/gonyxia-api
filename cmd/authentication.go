package cmd

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
	"github.com/inseefrlab/onyxia-api/internal/configuration"
	pkg "github.com/inseefrlab/onyxia-api/pkg"
)

type Claims struct {
	Email  string   `json:"email,omitempty"`
	Name   string   `json:"name,omitempty"`
	Groups []string `json:"groups,omitempty"`
	ID     string   `json:"preferred_username,omitempty"`
}

func AuthMiddleware(ctx context.Context, verifier *oidc.IDTokenVerifier) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := strings.TrimPrefix(c.Request.Header.Get("Authorization"), "Bearer ")
		token, err := verifier.Verify(ctx, tokenHeader)
		if err != nil {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		var IDTokenClaims Claims // ID Token payload is just JSON.
		if err := token.Claims(&IDTokenClaims); err != nil {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		var allClaims map[string]interface{}
		token.Claims(&allClaims)
		c.Set("claims", IDTokenClaims)
		region, _ := c.Get("region")
		fmt.Println(region.(configuration.Region).ID)
		c.Set("user", pkg.UserInfo{
			Email:    IDTokenClaims.Email,
			ID:       IDTokenClaims.ID,
			Name:     IDTokenClaims.Name,
			Groups:   IDTokenClaims.Groups,
			IP:       c.RemoteIP(),
			Projects: []pkg.Project{{Name: "todo"}},
		})

		c.Next()
	}
}
