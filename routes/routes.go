package routes

import (
	"github.com/gin-gonic/gin"
	"recommendation/controller"
	"recommendation/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	//eshop路由
	r.POST("/eshop/register", controller.EshopRegister)                //注册
	r.POST("/eshop/login", controller.EshopLogin)                      //登录
	r.GET("/eshop/info", middleware.AuthMiddleware(), controller.Info) //用中间件保护我们用户信息结构 用户信息

	//celebrity路由
	r.POST("celebrity/register", controller.CeleRegister)
	r.POST("celebrity/login", controller.CeleLogin)
	r.POST("celebrity/info", middleware.AuthMiddlewareForCele(), controller.Info)
	// 返回值
	return r
}
