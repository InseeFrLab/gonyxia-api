package cmd

import (
	"fmt"
	"net/http"

	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/inseefrlab/onyxia-api/internal/configuration"
)

type BuildInfo struct {
	Commit     string `json:"commit,omitempty"`
	CommitDate string `json:"commitDate,omitempty"`
}
type PublicConfiguration struct {
	Build   BuildInfo   `json:"build"`
	Regions interface{} `json:"regions"`
}

// PingExample godoc
// @Summary Get your public IP address.
// @Schemes
// @Description Get the public IP (as seen by this app).
// @Tags public
// @Produce json
// @Success 200 {string} Helloworld
// @Router /public/ip [get]
func ip(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ip": c.RemoteIP(),
	})
}

func healthcheck(c *gin.Context) {
	c.Status(http.StatusOK)
}

func getConfiguration(c *gin.Context) {
	info, _ := debug.ReadBuildInfo()
	var buildInfo = BuildInfo{}
	for _, kv := range info.Settings {
		fmt.Println(kv.Key)
		if kv.Key == "vcs.time" {
			buildInfo.CommitDate = kv.Value
		}
		if kv.Key == "vcs.revision" {
			buildInfo.Commit = kv.Value
		}
	}
	c.JSON(http.StatusOK, PublicConfiguration{
		Build:   buildInfo,
		Regions: configuration.Config.Regions,
	})
}

func registerPublicHandlers(r *gin.RouterGroup) {
	r.GET("/ip", ip)
	r.GET("/healthcheck", healthcheck)
	r.GET("/configuration", getConfiguration)
}
