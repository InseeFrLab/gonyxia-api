package cmd

import (
	"io"
	"net/http"

	"github.com/inseefrlab/onyxia-api/internal/helm"
	"github.com/inseefrlab/onyxia-api/internal/kubernetes"

	"github.com/gin-gonic/gin"
	eventsv1 "k8s.io/api/events/v1"
)

var namespace = "user-f2wbnp"

type App struct {
	ID    string `json:"id"`
	Chart string `json:"chart"`
}
type MyServices struct {
	Apps []App `apps:"apps"`
}

// @Summary List the services installed in a namespace.
// @Schemes
// @Description
// @Tags My lab
// @Produce json
// @Success 200
// @Router /my-lab/services [get]
func myServices(c *gin.Context) {
	myServices := MyServices{}
	for _, release := range helm.ListReleases() {
		myServices.Apps = append(myServices.Apps, App{ID: release.Name, Chart: release.Chart.Name()})
	}
	c.JSON(http.StatusOK, myServices)
}

func events(c *gin.Context) {
	c.Stream(func(w io.Writer) bool {
		for event := range kubernetes.GetEvents(namespace).ResultChan() {
			item := event.Object.(*eventsv1.Event)
			c.SSEvent("message", item)
		}
		return false
	})
}

// @Summary List the quotas in a namespace.
// @Schemes
// @Description
// @Tags My lab
// @Produce json
// @Success 200
// @Router /my-lab/quota [get]
func quotas(c *gin.Context) {
	resourceQuotas := kubernetes.GetOnyxiaResourceQuota(namespace)
	c.JSON(http.StatusOK, resourceQuotas.Status)
}

type model struct {
    Namespace string `json:"namespace" example:"default" format:"string"` // Namespace of the quota
    QuotaName string `json:"quotaName" example:"name" format:"string"`    // Name of the quota
    NewLimit  int    `json:"newLimit" example:"100" format:"int"`          // New limit for the quota
}


// @Summary Change the quotas for a namespace.
// @Schemes
// @Description
// @Tags My lab
// @Success 200
// @Accept json
// @Produce json
// @Param newQuota body model true "Modify quotas"
// @Router /my-lab/quota [post]
func updateOnyxiaQuota(c *gin.Context) {
	//	kubernetes.PostOnyxiaResourceQuotas(namespace)
}

func registerMyLabHandlers(r *gin.RouterGroup) {
	r.GET("/my-lab/services", myServices)
	r.GET("/my-lab/events", events)
	r.GET("/my-lab/quota", quotas)
	r.POST("/my-lab/quota", updateOnyxiaQuota)
}
