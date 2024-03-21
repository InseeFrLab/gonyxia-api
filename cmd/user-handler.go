package cmd

import (
	"net/http"

	"github.com/gin-gonic/gin"
	pkg "github.com/inseefrlab/onyxia-api/pkg"
)

// @Summary Get user info
// @Schemes
// @Description Get user info and projects
// @Tags user
// @Produce json
// @Success 200
// @Router /user/info [get]
func userInfo(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, user.(pkg.UserInfo))
}

func registerUserHandlers(r *gin.RouterGroup) {
	r.GET("/user/info", userInfo)
}
