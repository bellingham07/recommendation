package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"recommendation/common"
	"recommendation/model"
	"strings"
)

func Save(c *gin.Context) {
	// 获取用户id
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "mag": "权限不足"})
		c.Abort()
		return
	}

	tokenString = tokenString[6:]
	token, claims, err := common.ParseToken(tokenString)

	if err != nil || !token.Valid {
		fmt.Println("err", err)
		fmt.Println("token", token.Valid)
		c.JSON(http.StatusUnauthorized, gin.H{"code": 402, "mag": "权限不足"})
		c.Abort()
		return
	}

	db := common.GetDB()

	var good model.TbGood
	err1 := c.ShouldBind(&good)
	if err1 != nil {
		panic(err)
	}

	// 获取发货人姓名
	var eshop model.TbEshop
	db.Where("id=?", good.Eshop).Find(&eshop)

	// 收货人姓名
	var cele model.TbCelebrity
	db.Where("id=?", claims.UserId).Find(&cele)

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

	order.Celebrity = claims.UserId

	tx := db.Debug().Create(&order)
	if tx.Error != nil {
		panic(tx.Error)
	}
}
