package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"recommendation/common"
	"recommendation/model"
	"recommendation/response"
	"strings"
)

func GetAllGoods(ctx *gin.Context) {
	db := common.GetDB()

	//获取authorization header
	tokenString := ctx.GetHeader("Authorization")
	//validate token format
	if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足1.0"})
		ctx.Abort()
		return
	}
	tokenString = tokenString[6:] //Bearer 占六位

	token, claims, err := common.ParseToken(tokenString)
	//解析失败或者token无效
	if err != nil || !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足2.0"})
		ctx.Abort()
		return
	}

	id := claims.UserId

	var goods []model.TbGood
	db.Where("eshop=?", id).Find(&goods)
	response.Success(ctx, gin.H{"data": goods}, "success")
}

func SaveGood(ctx *gin.Context) {
	db := common.GetDB()

	//获取authorization header
	tokenString := ctx.GetHeader("Authorization")
	//validate token format
	if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足1.0"})
		ctx.Abort()
		return
	}
	tokenString = tokenString[6:] //Bearer 占六位

	token, claims, err := common.ParseToken(tokenString)
	//解析失败或者token无效
	if err != nil || !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足2.0"})
		ctx.Abort()
		return
	}

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

	good.Id = int(node.GetId())
	good.Eshop = int(claims.UserId)
	good.Status = 1

	if common.IsGoodExist(good.Name) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "good is existed")
		return
	}

	db.Save(&good)
	response.Success(ctx, nil, "添加成功")
}
