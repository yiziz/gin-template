package initializers

import "github.com/gin-gonic/gin"

// Initialize setups the app
func Initialize(appRouter *gin.Engine, appEnv string) {
	InitializeDatabase(appEnv)
	InitializeMiddleware(appRouter)
	InitializeRoutes(appRouter)
}
