package initial

import (
	"github.com/gin-gonic/gin"
	"recommendation/middleware"
	"recommendation/myRouters"
)

func Routers() *gin.Engine {
	router := gin.New()

	myRouter := new(myRouters.SystemGroup)

	// 配置跨域
	router.Use(middleware.CORSMiddleware())
	groupRegistry := router.Group("/")
	{
		myRouter.CeleRoute.InitCeleRoute(groupRegistry)
		myRouter.EshopRoute.InitEshopRoute(groupRegistry)
	}
	return router
}
