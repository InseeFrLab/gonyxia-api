package cmd

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Project struct {
	ID          string `json:"id"`
	Namespace   string `json:"namespace"`
	Name        string `json:"name"`
	VaultTopDir string `json:"vaultTopDir"`
}

type UserInfo struct {
	Email    string    `json:"email,omitempty"`
	ID       string    `json:"idep,omitempty"`
	Name     string    `json:"nomComplet,omitempty"`
	IP       string    `json:"ip,omitempty"`
	Groups   []string  `json:"groups,omitempty"`
	Projects []Project `json:"projects"`
}

func RegisterUserHandlers(r *gin.Engine) {
	r.GET("/user/info", func(c *gin.Context) {

		claims, _ := c.Get("claims")
		typedClaims := claims.(Claims)
		userInfo := UserInfo{
			Email:    typedClaims.Email,
			ID:       typedClaims.ID,
			Name:     typedClaims.Name,
			Groups:   typedClaims.Groups,
			IP:       c.RemoteIP(),
			Projects: []Project{{Name: "todo"}},
		}
		c.JSON(http.StatusOK, userInfo)
	})
}
