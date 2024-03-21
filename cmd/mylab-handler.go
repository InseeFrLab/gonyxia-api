package cmd

import (
	"io"
	"net/http"

	"github.com/inseefrlab/onyxia-admin/internal/helm"
	"github.com/inseefrlab/onyxia-admin/internal/kubernetes"

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

type Quotas struct {
	Spec  []Quota `spec:"spec"`
	Usage []Quota `usage:"usage"`
}

type Quota struct {
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

func quotas(c *gin.Context) {
	/*quotas := Quotas{}
	for {
		quotas.Quota = append
	}*/
}

func registerMyLabHandlers(r *gin.RouterGroup) {
	r.GET("/my-lab/services", myServices)
	r.GET("/my-lab/events", events)
	r.GET("/my-lab/quota", quotas)
}
