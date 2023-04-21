package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"recommendation/common"
	"recommendation/model"
)

func GetAll(ctx *gin.Context) {
	db := common.GetDB()
	var goods []model.TbGood
	db.Find(&goods)
	for _, good := range goods {
		fmt.Println(good)
	}
}

func SaveGood(ctx *gin.Context) {
	var good model.TbGood
	err := ctx.ShouldBind(&good)
	if err != nil {
		panic(err)
	}
	fmt.Println(good)
}
