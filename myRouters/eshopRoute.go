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
		eg.POST("/register", controller.EshopRegister) //注册
		eg.POST("/login", controller.EshopLogin)       //登录
		eg.POST("/saveGood", controller.SaveGood)      //新增商品
		eg.POST("/updateGood", controller.UpdateGood)  //修改商品
		eg.POST("/status", controller.ChangeStatus)
		eg.POST("/update", controller.UpdateEshop)
		eg.POST("/upload", controller.EUpdateAvatar)   //修改头像
		eg.POST("/uploadGood", controller.SaveGoodImg) //上传商品图片
		eg.POST("/delete", controller.Delete)

		eg.GET("/info", middleware.AuthMiddleware(), controller.Info) //用中间件保护我们用户信息结构 用户信息
		eg.GET("/getAllGoods", controller.GetAllGoodsById)            //获取全部商品
		eg.GET("/findAll", controller.GetAll)                         //获取全部网红
		eg.GET("/find", controller.GetEshopInfo)
		eg.GET("/getContract", controller.GetAllContract)
	}
}
