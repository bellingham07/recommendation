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
		cg.POST("/updateInfo", controller.UpdateInfo)
		cg.POST("/upload", controller.UpdateAvatar) //修改头像
		cg.POST("/save", controller.SaveCToE)       //创建订单
		cg.POST("/order", controller.Save)          //保存订单

		cg.GET("/find", controller.GetUserInfo)
		cg.GET("/findAll", controller.GetAllUser) //获取全部商家
		cg.GET("/allGoods", controller.GetAllGoods)
		cg.GET("/getContract", controller.GetCContract)
	}
}
