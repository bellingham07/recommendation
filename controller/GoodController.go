package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"recommendation/common"
	"recommendation/model"
	"recommendation/ossUtils"
	"recommendation/response"
	"strconv"
)

func GetAllGoodsById(ctx *gin.Context) {
	db := common.GetDB()

	id := common.GetId(ctx)

	var goods []model.TbGood
	db.Debug().Where("eshop=?", id).Find(&goods)
	response.Success(ctx, gin.H{"data": goods})
}

func GetAllGoods(ctx *gin.Context) {
	db := common.GetDB()
	var goods []model.TbGood
	db.Find(&goods)
	response.Success(ctx, gin.H{"data": goods})
}

func SaveGood(ctx *gin.Context) {
	db := common.GetDB()

	var good model.TbGood
	err1 := ctx.ShouldBind(&good)
	if err1 != nil {
		panic(err1)
	}

	//use snowflake generate a new id
	node, err := common.NewWorker(1)
	if err != nil {
		panic(err)
	}

	good.Id = node.GetId()
	good.Eshop = common.GetId(ctx)
	good.Status = 1

	if common.IsGoodExist(good.Name) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "good is existed")
		return
	}

	db.Save(&good)
	response.Success(ctx, nil)
}

func SaveGoodImg(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")
	id := ctx.PostForm("id")

	url := ossUtils.OssUtils(file, id)

	db := common.GetDB()
	tx := db.Table("tb_good").Where("id=?", id).Update("img", url)
	if tx.Error != nil {
		panic(tx.Error)
	}
	response.Success(ctx, gin.H{"url": url})
}

func ChangeStatus(ctx *gin.Context) {
	db := common.GetDB()

	//get params
	status := ctx.PostForm("status")
	id, _ := strconv.Atoi(ctx.PostForm("id"))
	tx := db.Table("tb_good").Where("id=?", id).Update("status", status)
	if tx.Error != nil {
		response.Response(ctx, http.StatusInternalServerError, 422, nil, "server error")
		return
	}
	response.Success(ctx, nil)
}

func Delete(ctx *gin.Context) {
	db := common.GetDB()

	var good model.TbGood
	good.Id = ctx.PostForm("id")
	tx := db.Delete(&good)
	if tx.RowsAffected == 0 {
		response.Fail(ctx, nil)
		return
	}
	response.Success(ctx, nil)
}

func UpdateGood(ctx *gin.Context) {
	var good model.TbGood
	err := ctx.ShouldBind(&good)
	if err != nil {
		panic(err)
	}
	db := common.GetDB()
	db.Where("id=?", good.Id).Updates(model.TbGood{Name: good.Name, Category: good.Category, Brand: good.Brand, MarketPrice: good.MarketPrice, CelebrityPrice: good.CelebrityPrice, GoodUrl: good.GoodUrl, Intro: good.Intro})
	response.Success(ctx, nil)
}
