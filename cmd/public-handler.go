package cmd

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingExample godoc
// @Summary Get your public IP address.
// @Schemes
// @Description Get the public IP (as seen by this app).
// @Tags public
// @Produce json
// @Success 200 {string} Helloworld
// @Router /public/ip [get]
func ip(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ip": c.RemoteIP(),
	})
}

func registerPublicHandlers(r *gin.Engine) {
	r.GET("/public/ip", ip)
}
