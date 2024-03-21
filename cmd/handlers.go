package cmd

import "github.com/gin-gonic/gin"

func RegisterPublicHandlers(r *gin.RouterGroup) {
	registerPublicHandlers(r)
}

func RegisterPrivateHandlers(r *gin.RouterGroup) {
	registerUserHandlers(r)
	registerMyLabHandlers(r)
}
