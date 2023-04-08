package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"recommendation/common"
	"recommendation/dto"
	"recommendation/model"
	"recommendation/response"
)

func EshopRegister(ctx *gin.Context) {
	db := common.GetDB()
	//获取参数
	account := ctx.PostForm("username")
	telephone := ctx.PostForm("phonenumber")
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")
	//数据验证
	if isTelephoneExist(db, telephone) {
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
		Id:       int(newId),
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
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	//Determine if the user exists
	var eshop model.TbEshop
	db.Where("username=?", username).First(&eshop)
	if eshop.Id == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "user is not exist"})
		return
	}
	//Determine if the password is correct
	if err := bcrypt.CompareHashAndPassword([]byte(eshop.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "password is error"})
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

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.TbEshop
	db.Where("tel=?", telephone).First(&user)
	if user.Id != 0 {
		fmt.Println("are kidding me?")
		return true
	}
	return false
}
