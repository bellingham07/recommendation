package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"recommendation/common"
	"recommendation/model"
	"recommendation/response"
	"strings"
	"time"
)

func SaveCToE(ctx *gin.Context) {
	db := common.GetDB()

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
	fmt.Println("contract.Eshop", contract.Eshop)

	//获取当前用户id celebrity Id
	tokenString := ctx.GetHeader("Authorization")
	if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
		response.Fail(ctx, nil)
		ctx.Abort() //阻止调用后续的处理函数
		return
	}
	tokenString = tokenString[6:]

	token, claims, err := common.ParseToken(tokenString)
	if err != nil || !token.Valid {
		response.Fail(ctx, nil)
		ctx.Abort()
		return
	}
	contract.Celebrity = claims.UserId

	contract.Status = "1"
	contract.CreateBy = claims.UserId
	contract.CreateTime = time.Now()
	contract.StartTime = time.Now()
	contract.EndTime = contract.CreateTime.AddDate(3, 0, 0)

	tx := db.Save(&contract)
	if tx.RowsAffected == 0 {
		panic(tx.Error)
	}
	response.Success(ctx, nil)
}

func GetCContract(ctx *gin.Context) {
	db := common.GetDB()
	var contracts []model.TbContract

	tokenString := ctx.GetHeader("Authorization")
	if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
		response.Fail(ctx, nil)
		ctx.Abort()
		return
	}
	tokenString = tokenString[6:]

	token, claims, err := common.ParseToken(tokenString)
	if err != nil || !token.Valid {
		response.Fail(ctx, nil)
		ctx.Abort()
		return
	}

	db.Where("create_by=?", claims.UserId).Find(&contracts)
	fmt.Println("contract", contracts)
	response.Success(ctx, gin.H{"data": contracts})
}
