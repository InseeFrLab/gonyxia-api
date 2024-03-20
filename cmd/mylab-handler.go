package cmd

import (
	"net/http"

	"github.com/inseefrlab/onyxia-admin/internal/helm"

	"github.com/gin-gonic/gin"
)

type App struct {
	ID    string `json:"id"`
	Chart string `json:"chart"`
}
type MyServices struct {
	Apps []App `apps:"apps"`
}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /my-lab/services [get]
func myServices(c *gin.Context) {
	myServices := MyServices{}
	for _, release := range helm.ListReleases() {
		myServices.Apps = append(myServices.Apps, App{ID: release.Name, Chart: release.Chart.Name()})
	}
	c.JSON(http.StatusOK, myServices)
}

func registerMyLabHandlers(r *gin.Engine) {
	r.GET("/my-lab/services", myServices)
}
