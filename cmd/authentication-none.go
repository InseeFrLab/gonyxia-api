package cmd

import (
	"github.com/gin-gonic/gin"
	pkg "github.com/inseefrlab/onyxia-api/pkg"
)

func NoAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestContext := GetRequestContext(c)
		requestContext.User = pkg.UserInfo{
			Email:    "johndoe@example.com",
			ID:       "johndoe",
			Name:     "John Doe",
			Groups:   []string{},
			IP:       c.RemoteIP(),
			Projects: []pkg.Project{{Name: "todo", ID: "todooo"}},
		}
		SetRequestContext(c, requestContext)
		c.Next()
	}
}
