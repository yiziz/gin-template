package routes

import "github.com/gin-gonic/gin"

// RouteFuncType is a func is called to add routes to rootRouter
type RouteFuncType func(rootRouter *gin.Engine, rootPath string)

var routeFuncList []RouteFuncType

// AddRoutesTo adds routes form routeFuncList to appRouter
func AddRoutesTo(appRouter *gin.Engine) {
	for _, routeFunc := range routeFuncList {
		routeFunc(appRouter, "/api/")
	}
}
