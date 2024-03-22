package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/inseefrlab/onyxia-api/internal/configuration"
	"github.com/inseefrlab/onyxia-api/pkg"
)

type RequestContext struct {
	User    pkg.UserInfo
	Project pkg.Project
	Region  configuration.Region
}

var requestContextKey = "requestContext"

func GetRequestContext(c *gin.Context) RequestContext {
	context, exists := c.Get(requestContextKey)
	if !exists {
		context = RequestContext{}
		c.Set(requestContextKey, context)
	}
	return context.(RequestContext)
}

func SetRequestContext(c *gin.Context, newContext RequestContext) {
	c.Set(requestContextKey, newContext)
}
