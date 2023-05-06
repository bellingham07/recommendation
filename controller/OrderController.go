package controller

import (
	"github.com/gin-gonic/gin"
	"recommendation/common"
	"recommendation/model"
)

func Save(c *gin.Context) {
	db := common.GetDB()

	var good model.TbGood
	err := c.ShouldBind(&good)
	if err != nil {
		panic(err)
	}

	// 获取发货人姓名
	var eshop model.TbEshop
	db.Where("id=?", good.Eshop).Find(&eshop)

	// 收货人姓名
	var cele model.TbCelebrity
	db.Where("id=?", common.GetId(c)).Find(&cele)

	// 将商品信息存入订单信息中
	order := model.TbOrder{
		Eshop:     good.Eshop,
		Good:      good.Name,
		Consignor: eshop.Name, // 发货人
		Consignee: cele.Name,  // 收货人
	}
	// 生成订单id
	node, err := common.NewWorker(1)
	order.Id = node.GetId()

	order.Celebrity = common.GetId(c)

	tx := db.Create(&order)
	if tx.Error != nil {
		panic(tx.Error)
	}
}

func ESave(c *gin.Context) {
	db := common.GetDB()

	var good model.TbGood
	err := c.ShouldBind(&good)
	if err != nil {
		panic(err)
	}

	// 获取收货人姓名
	var eshop model.TbEshop
	db.Where("id=?", common.GetId(c)).Find(&eshop)

	// 发货人姓名
	var eshop1 model.TbEshop
	db.Where("id=?", good.Eshop).Find(&eshop1)

	// 将商品信息存入订单信息中
	order := model.TbOrder{
		Eshop:     good.Eshop,
		Good:      good.Name,
		Consignor: eshop1.Name, // 发货人
		Consignee: eshop.Name,  // 收货人
	}
	// 生成订单id
	node, err := common.NewWorker(1)
	order.Id = node.GetId()

	order.Celebrity = common.GetId(c)

	tx := db.Create(&order)
	if tx.Error != nil {
		panic(tx.Error)
	}
}
