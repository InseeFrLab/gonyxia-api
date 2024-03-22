package cmd

import (
	"fmt"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/inseefrlab/onyxia-api/internal/kubernetes"
)

type OnboardingRequest struct {
	Group string `json:"group" `
}

// @Summary Init a namespace for a user or a group
// @Schemes
// @Description Create or replace the namespace of the user or the namespace of a group if the user is in the requested group and the according rbac policies. with the group prefix / user prefix of the region
// @Tags Onboarding
// @Consume json
// @Produce json
// @Success 200
// @Router /onboarding [post]
func onboarding(c *gin.Context) {
	var onboardingRequest OnboardingRequest
	if err := c.BindJSON(&onboardingRequest); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	requestContext := GetRequestContext(c)
	if requestContext.Region.Services.SingleNamespace {
		c.AbortWithStatusJSON(400, gin.H{"error": "Instance is in single namespace mode"})
		return
	}

	if !requestContext.Region.Services.AllowNamespaceCreation {
		c.AbortWithStatusJSON(400, gin.H{"error": "This instance does not allow namespace creation"})
		return
	}

	if onboardingRequest.Group != "" && !slices.Contains(requestContext.User.Groups, onboardingRequest.Group) {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": fmt.Sprintf("User %s does not belong to group %s", requestContext.User.ID, onboardingRequest.Group)})
		return
	}

	var namespace string
	if onboardingRequest.Group == "" {
		namespace = fmt.Sprintf("%s%s", requestContext.Region.Services.NamespacePrefix, requestContext.User.ID)
	} else {
		namespace = fmt.Sprintf("%s%s", requestContext.Region.Services.GroupNamespacePrefix, onboardingRequest.Group)
	}
	kubernetes.InitNamespace(namespace)
}

func registerOnboardingHandlers(r *gin.RouterGroup) {
	r.POST("/onboarding", onboarding)
}
