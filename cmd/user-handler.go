package cmd

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get user info
// @Schemes
// @Description Get user info and projects
// @Tags user
// @Produce json
// @Success 200
// @Router /user/info [get]
func userInfo(c *gin.Context) {
	requestContext := GetRequestContext(c)
	fmt.Printf("Project %s", requestContext.Project.ID)
	fmt.Println()
	fmt.Printf("Region %s", requestContext.Region.ID)
	fmt.Println()
	c.JSON(http.StatusOK, requestContext.User)
}

func registerUserHandlers(r *gin.RouterGroup) {
	r.GET("/user/info", userInfo)
}
