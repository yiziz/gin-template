package routes

import (
	"github.com/bamzi/jobrunner"
	"github.com/gin-gonic/gin"
)

// special routes for jobrunner

// JobJSON returns jobrunner JSON
func JobJSON(c *gin.Context) {
	// returns a map[string]interface{} that can be marshalled as JSON
	c.JSON(200, jobrunner.StatusJson())
}

// JobHTML returns jobrunner monitoring page
func JobHTML(c *gin.Context) {
	// Returns the template data pre-parsed
	c.HTML(200, "", jobrunner.StatusPage())

}

func init() {
	routeFuncList = append(routeFuncList, func(rootRouter *gin.Engine, rootPath string) {
		// Resource to return the JSON data
		rootRouter.GET("/jobrunner/json", JobJSON)

		// Load template file location relative to the current working directory
		rootRouter.LoadHTMLGlob("app/views/jobrunner/status.html")

		// Returns html page at given endpoint based on the loaded
		// template from above
		rootRouter.GET("/jobrunner/html", JobHTML)
	})
}
