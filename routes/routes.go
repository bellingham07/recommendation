package routes

import (
	"github.com/gin-gonic/gin"
	"recommendation/controller"
	"recommendation/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	r.POST("/api/auth/register", controller.Register) //注册
	//r.POST("/api/auth/login", controller.Login)                           //登录
	//r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info) //用中间件保护我们用户信息结构 用户信息
	return r
}
