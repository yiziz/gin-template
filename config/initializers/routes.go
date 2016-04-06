package initializers

import (
	"github.com/yiziz/gin-template/config/routes"
	"github.com/gin-gonic/gin"
)

// InitializeRoutes adds routes to appRouter
func InitializeRoutes(appRouter *gin.Engine) {
	routes.AddRoutesTo(appRouter)
}
