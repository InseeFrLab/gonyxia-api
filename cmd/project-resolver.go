package cmd

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inseefrlab/onyxia-api/pkg"
)

func ProjectResolver() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestContext := GetRequestContext(c)
		if len(requestContext.User.Projects) == 0 {
			c.AbortWithStatusJSON(500, gin.H{"error": fmt.Sprintf("User %s has no project", requestContext.User.ID)})
			return
		}

		headerProject := c.GetHeader("ONYXIA-PROJECT")
		if headerProject == "" {
			requestContext.Project = requestContext.User.Projects[0]
		} else {
			var foundProject pkg.Project
			for _, project := range requestContext.User.Projects {
				if project.ID == headerProject {
					foundProject = project
					requestContext.Project = project
				}
			}
			if foundProject.ID == "" {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Requested project not found"})
			}
		}
		SetRequestContext(c, requestContext)
		c.Next()
	}
}
