package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"recommendation/common"
	"recommendation/model"
	"recommendation/response"
)

func Register(ctx *gin.Context) {
	db := common.GetDB()
	//获取参数
	var requestUser = model.TbEshop{}
	ctx.Bind(requestUser)
	fmt.Println("__________________________________________________")
	fmt.Println(requestUser)
	account := requestUser.Username
	telephone := requestUser.Tel
	name := requestUser.Name
	password := requestUser.Password
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

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.TbEshop
	db.Where("tel=?", telephone).First(&user)
	if user.Id != 0 {
		fmt.Println("are kidding me?")
		return true
	}
	return false
}
