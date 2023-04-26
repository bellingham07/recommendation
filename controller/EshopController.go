package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"recommendation/common"
	"recommendation/dto"
	"recommendation/model"
	"recommendation/ossUtils"
	"recommendation/response"
	"strings"
)

func EshopRegister(ctx *gin.Context) {
	db := common.GetDB()
	//获取参数
	account := ctx.PostForm("username")
	telephone := ctx.PostForm("phonenumber")
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")
	//数据验证
	if common.IsTelephoneExist(db, telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户已经存在")
		return
	}

	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}

	node, err := common.NewWorker(1)
	if err != nil {
		panic(err)
	}

	newId := node.GetId()

	newUser := model.TbEshop{
		Id:       newId,
		Username: account,
		Name:     name,
		Tel:      telephone,
		Password: string(hasedPassword),
	}

	db.Create(&newUser)

	//发放token
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		return
	}

	response.Success(ctx, gin.H{"token": token}, "注册成功")
}

func EshopLogin(ctx *gin.Context) {
	db := common.GetDB()
	//get parameter
	var params model.TbEshop

	err1 := ctx.ShouldBind(&params)
	if err1 != nil {
		panic(err1)
	}
	//Determine if the user existed
	var eshop model.TbEshop
	db.Where("username=?", params.Username).First(&eshop)
	if eshop.Id == "" {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "user is not exist"})
		return
	}
	//Determine if the password is correct
	if err := bcrypt.CompareHashAndPassword([]byte(eshop.Password), []byte(params.Password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "password is not correct"})
		return
	}

	//distribute token
	token, err := common.ReleaseToken(eshop)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "server error"})
	}

	//return result
	response.Success(ctx, gin.H{"token": token}, "login successful")
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": dto.ToUserDto(user.(model.TbEshop))}})
}

func GetAllUser(ctx *gin.Context) {
	db := common.GetDB()
	var eshop []model.TbEshop
	db.Select("id,name,tel,email,avatar,intro,platform,platform_url,credit_point,age").Find(&eshop)

	response.Success(ctx, gin.H{"data": eshop}, "success")
}

func GetEshopInfo(ctx *gin.Context) {
	//获取authorization header
	tokenString := ctx.GetHeader("Authorization")
	//validate token format
	if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
		ctx.Abort()
		return
	}
	tokenString = tokenString[6:] // bearer 占六位
	token, claims, err := common.ParseToken(tokenString)
	//解析失败或者token无效
	if err != nil || !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
		ctx.Abort()
		return
	}

	db := common.GetDB()

	var eshop model.TbEshop
	eshop.Id = claims.UserId
	db.Select("id", "username", "name", "tel", "email", "seller", "avatar", "intro", "platform_url", "platform", "age", "credit_point").Find(&eshop)
	if eshop.Tel == "" {
		response.Fail(ctx, nil, "fail")
		return
	}
	response.Success(ctx, gin.H{"data": eshop}, "success")
}

func UpdateEshop(ctx *gin.Context) {
	var eshop model.TbEshop
	db := common.GetDB()
	err := ctx.ShouldBind(&eshop)
	if err != nil {
		panic(err)
	}
	tx := db.Model(&eshop).Where("tel=?", eshop.Tel).Updates(model.TbEshop{Name: eshop.Name, Tel: eshop.Tel, Age: eshop.Age, Email: eshop.Email, Platform: eshop.Platform, PlatformUrl: eshop.PlatformUrl, Intro: eshop.Intro})
	if tx.RowsAffected == 0 {
		panic(tx.Error)
		return
	}
}

func EUpdateAvatar(ctx *gin.Context) {
	db := common.GetDB()

	file, _ := ctx.FormFile("file")
	tel := ctx.PostForm("tel")
	username := ctx.PostForm("username")

	url := ossUtils.OssUtils(file, username)
	tx := db.Table("tb_eshop").Where("tel=?", tel).Update("avatar", url)
	if tx.Error != nil {
		fmt.Println("update fail")
		return
	}
	response.Success(ctx, gin.H{"url": url}, "success")
}
