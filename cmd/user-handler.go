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

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func userInfo(c *gin.Context) {
	claims, _ := c.Get("claims")
	var userInfo UserInfo
	if claims == nil {
		userInfo = UserInfo{
			Email:    "johndoe@example.com",
			ID:       "johndoe",
			Name:     "John Doe",
			Groups:   []string{},
			IP:       c.RemoteIP(),
			Projects: []Project{{Name: "todo"}},
		}

	} else {
		typedClaims := claims.(Claims)
		userInfo = UserInfo{
			Email:    typedClaims.Email,
			ID:       typedClaims.ID,
			Name:     typedClaims.Name,
			Groups:   typedClaims.Groups,
			IP:       c.RemoteIP(),
			Projects: []Project{{Name: "todo"}},
		}
	}
	c.JSON(http.StatusOK, userInfo)
}

func registerUserHandlers(r *gin.Engine) {
	r.GET("/user/info", userInfo)
}
