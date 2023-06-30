package initial

import (
	"github.com/gin-gonic/gin"
	"recommendation/middle"
	"recommendation/myRouters"
)

func Routers() *gin.Engine {
	router := gin.New()
	myRouter := new(myRouters.SystemGroup)
	// 配置跨域
	router.Use(middle.CORSMiddleware())
	groupRegistry := router.Group("/")
	{
		myRouter.CeleRoute.InitCeleRoute(groupRegistry)
		myRouter.EshopRoute.InitEshopRoute(groupRegistry)
	}
	return router
}
