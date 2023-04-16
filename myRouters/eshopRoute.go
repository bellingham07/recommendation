package myRouters

import (
	"github.com/gin-gonic/gin"
	"recommendation/controller"
	"recommendation/middleware"
)

type EshopRoute struct {
}

func (*EshopRoute) InitEshopRoute(g *gin.RouterGroup) {
	// eshop 路由
	eg := g.Group("/eshop")
	{
		eg.POST("/register", controller.EshopRegister)                //注册
		eg.POST("/login", controller.EshopLogin)                      //登录
		eg.GET("/info", middleware.AuthMiddleware(), controller.Info) //用中间件保护我们用户信息结构 用户信息
	}
}
