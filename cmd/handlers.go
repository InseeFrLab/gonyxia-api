package cmd

import "github.com/gin-gonic/gin"

func RegisterHandlers(r *gin.Engine) {
	registerUserHandlers(r)
	registerMyLabHandlers(r)
}
