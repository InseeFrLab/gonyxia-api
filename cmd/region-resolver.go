package cmd

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inseefrlab/onyxia-api/internal/configuration"
)

func RegionResolver() gin.HandlerFunc {
	if len(configuration.Config.Regions) == 0 {
		panic("No region configured")
	}
	defaultRegion := configuration.Config.Regions[0]
	return func(c *gin.Context) {
		requestContext := GetRequestContext(c)
		headerRegion := c.GetHeader("ONYXIA-REGION")
		if headerRegion == "" {
			requestContext.Region = defaultRegion
		} else {
			var foundRegion configuration.Region
			for _, region := range configuration.Config.Regions {
				if region.ID == headerRegion {
					foundRegion = region
					requestContext.Region = foundRegion
				}
			}
			if foundRegion.ID == "" {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Requested region not found"})
			}
		}
		SetRequestContext(c, requestContext)
		c.Next()
	}
}
