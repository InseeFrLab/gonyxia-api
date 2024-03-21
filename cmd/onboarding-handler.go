package cmd

import (
	"net/http"

	"github.com/inseefrlab/onyxia-api/internal/helm"

	"github.com/gin-gonic/gin"
)

// @Summary Init a namespace for a user or a group
// @Schemes
// @Description Create or replace the namespace of the user or the namespace of a group if the user is in the requested group and the according rbac policies. with the group prefix / user prefix of the region
// @Tags Onboarding
// @Produce json
// @Success 200
// @Router /onboarding [post]
func onboarding(c *gin.Context) {
	myServices := MyServices{}
	for _, release := range helm.ListReleases() {
		myServices.Apps = append(myServices.Apps, App{ID: release.Name, Chart: release.Chart.Name()})
	}
	c.JSON(http.StatusOK, myServices)
}

func registerOnboardingHandlers(r *gin.RouterGroup) {
	r.POST("/onboarding", onboarding)
}
