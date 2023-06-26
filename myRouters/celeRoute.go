package myRouters

import (
	"github.com/gin-gonic/gin"
	"recommendation/controller"
	"recommendation/middle"
)

type CeleRoute struct {
}

func (*CeleRoute) InitCeleRoute(g *gin.RouterGroup) {
	// celebrity路由
	cg := g.Group("/celebrity")
	{
		cg.POST("/register", controller.CeleRegister)
		cg.POST("/login", controller.CeleLogin)
		cg.POST("/info", middle.AuthMiddlewareForCele(), controller.InfoForCele)
		cg.POST("/updateInfo", controller.UpdateInfo)
		cg.POST("/upload", controller.UpdateAvatar)         //修改头像
		cg.POST("/save", controller.SaveCToE)               //创建合约
		cg.POST("/order", controller.Save)                  //保存订单
		cg.POST("/addr", controller.Address)                //添加地址
		cg.POST("/upAddr", controller.UpdateAddr)           //修改地址
		cg.POST("/deleteAddr", controller.DeleteAddr)       //删除地址
		cg.POST("/isLiked", controller.IsLiked)             //获取是否点赞
		cg.POST("/like", controller.Like)                   //获取是否点赞
		cg.POST("/send/mail/code", controller.SendMailCode) // 发送验证码

		cg.GET("/address", controller.GetAddrById) //获取地址
		cg.GET("/find", controller.GetUserInfo)
		cg.GET("/findAll", controller.GetAllUser) //获取全部商家
		cg.GET("/allGoods", controller.GetAllGoods)
		cg.GET("/contract", controller.GetContract)
	}
}
