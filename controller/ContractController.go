package controller

import (
	"github.com/gin-gonic/gin"
	"recommendation/common"
	"recommendation/database"
	"recommendation/model"
	"recommendation/response"
	"time"
)

// SaveCToE 网红向商家发起合作
func SaveCToE(ctx *gin.Context) {
	db := database.GetDB()

	var contract model.TbContract

	//雪花算法生成id
	node, err1 := common.NewWorker(1)
	if err1 != nil {
		panic(err1)
	}
	contract.Id = node.GetId()

	// 获得eshop id
	err := ctx.ShouldBind(&contract)
	if err != nil {
		panic(err)
	}

	//获取当前用户id celebrity Id
	contract.Celebrity = common.GetId(ctx)

	contract.Status = "1"
	contract.CreateBy = common.GetId(ctx)
	contract.CreateTime = time.Now()
	contract.StartTime = time.Now()
	contract.EndTime = contract.CreateTime.AddDate(3, 0, 0)

	tx := db.Save(&contract)
	if tx.RowsAffected == 0 {
		panic(tx.Error)
	}
	response.Success(ctx, nil)
}

func GetContract(ctx *gin.Context) {
	db := database.GetDB()

	var contracts []model.TbContract

	db.Where("create_by=?", common.GetId(ctx)).Find(&contracts)
	response.Success(ctx, gin.H{"data": contracts})
}

// SaveEToC 商家向网红发起合作
func SaveEToC(c *gin.Context) {
	db := database.GetDB()
	var contract model.TbContract

	//雪花算法生成id
	node, err1 := common.NewWorker(1)
	if err1 != nil {
		panic(err1)
	}
	contract.Id = node.GetId()

	// 获得eshop id
	err := c.ShouldBind(&contract)
	if err != nil {
		panic(err)
	}

	//获取当前用户id celebrity Id
	contract.Eshop = common.GetId(c)

	contract.Status = "1"
	contract.CreateBy = common.GetId(c)
	contract.CreateTime = time.Now()
	contract.StartTime = time.Now()
	contract.EndTime = contract.CreateTime.AddDate(3, 0, 0)

	tx := db.Save(&contract)
	if tx.RowsAffected == 0 {
		panic(tx.Error)
	}
	response.Success(c, nil)
}
