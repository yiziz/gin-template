package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/yiziz/gin-template/app/controllers"
)

func init() {
	routeFuncList = append(routeFuncList, func(rootRouter *gin.Engine, rootPath string) {
		router := rootRouter.Group(rootPath + "/")
		{
			router.GET("", controllers.Home)
		}
		// can't have wildcards and static name under same path
		// see https://github.com/gin-gonic/gin/issues/388
		// so using /user/:id instead of /users/:id
		router = rootRouter.Group(rootPath + "/")
		{
		}
	})
}
