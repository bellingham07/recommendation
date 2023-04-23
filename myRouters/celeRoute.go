package myRouters

import (
	"github.com/gin-gonic/gin"
	"recommendation/controller"
	"recommendation/middleware"
)

type CeleRoute struct {
}

func (*CeleRoute) InitCeleRoute(g *gin.RouterGroup) {
	// celebrity路由
	cg := g.Group("/celebrity")
	{
		cg.POST("/register", controller.CeleRegister)
		cg.POST("/login", controller.CeleLogin)
		cg.POST("/info", middleware.AuthMiddlewareForCele(), controller.InfoForCele)
		cg.GET("/find", controller.GetUserInfo)
		cg.POST("/updateInfo", controller.UpdateInfo)
		cg.GET("/findAll", controller.GetAllUser) //获取全部商家
		cg.GET("/allGoods", controller.GetAllGoods)
	}
}
