package cmd

import (
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
	Build             BuildInfo          `json:"build"`
	Regions           interface{}        `json:"regions"`
	OIDCConfiguration *OIDCConfiguration `json:"oidcConfiguration,omitempty"`
}
type OIDCConfiguration struct {
	IssuerURI        string `json:"issuerURI,omitempty"`
	ClientID         string `json:"clientID,omitempty"`
	ExtraQueryParams string `json:"extraQueryParams,omitempty"`
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
		if kv.Key == "vcs.time" {
			buildInfo.CommitDate = kv.Value
		}
		if kv.Key == "vcs.revision" {
			buildInfo.Commit = kv.Value
		}
	}
	var publicConfiguration = PublicConfiguration{
		Build:   buildInfo,
		Regions: configuration.Config.Regions,
	}
	if configuration.IsAuthenticationEnabled() {
		publicConfiguration.OIDCConfiguration = &OIDCConfiguration{
			IssuerURI:        configuration.Config.OIDC.IssuerURI,
			ClientID:         configuration.Config.OIDC.ClientID,
			ExtraQueryParams: configuration.Config.OIDC.ExtraQueryParams,
		}
	}
	c.JSON(http.StatusOK, publicConfiguration)
}

func registerPublicHandlers(r *gin.RouterGroup) {
	r.GET("/ip", ip)
	r.GET("/healthcheck", healthcheck)
	r.GET("/configuration", getConfiguration)
}
