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
)

func CeleRegister(ctx *gin.Context) {
	// connect database
	db := common.GetDB()
	// get register parameter
	account := ctx.PostForm("username")
	password := ctx.PostForm("username")
	name := ctx.PostForm("name")
	tel := ctx.PostForm("phonenumber")

	//data validation
	if common.IsTelephoneExist(db, tel) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "tel is existed")
		return
	}

	// encrypt password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "server error")
		return
	}

	//use snowflake generate a new id
	node, err := common.NewWorker(1)
	if err != nil {
		panic(err)
	}

	newId := node.GetId()

	// create a new entity to save the info
	newCele := model.TbCelebrity{
		Id:          newId,
		Username:    account,
		Name:        name,
		PhoneNumber: tel,
		Password:    string(hashedPassword),
	}

	// save to database
	db.Create(&newCele)

	// distribute token
	token, err := common.ReleaseTokenForCele(newCele)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "system error"})
		return
	}

	// return token
	response.Success(ctx, gin.H{"token": token})
}

func CeleLogin(ctx *gin.Context) {
	// connect database
	db := common.GetDB()

	// get parameter
	var params model.TbCelebrity
	err1 := ctx.ShouldBind(&params)
	if err1 != nil {
		panic(err1)
	}

	// determine if the user is existed
	var cele model.TbCelebrity

	db.Debug().Where("username=?", params.Username).First(&cele)
	if cele.Id == "" {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 500, "mag": "user is not existed"})
		return
	}
	//determine if the password is correct
	if err := bcrypt.CompareHashAndPassword([]byte(cele.Password), []byte(params.Password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 500, "msg": "password is not correct"})
		return
	}

	// distribute token
	token, err := common.ReleaseTokenForCele(cele)
	if err != nil {
		panic(err)
	}

	response.Success(ctx, gin.H{"token": token})
}

func GetUserInfo(ctx *gin.Context) {
	// connect database
	db := common.GetDB()
	var cele model.TbCelebrity
	db.Where("id=?", common.GetId(ctx)).First(&cele)

	newCele := model.TbCelebrity{
		Username:    cele.Username,
		PhoneNumber: cele.PhoneNumber,
		PlatformUrl: cele.PlatformUrl,
		Platform:    cele.Platform,
		Email:       cele.Email,
		Name:        cele.Name,
		RealName:    cele.RealName,
		Sex:         cele.Sex,
		Age:         cele.Age,
		Intro:       cele.Intro,
		CreditPoint: cele.CreditPoint,
		Avatar:      cele.Avatar,
	}
	response.Success(ctx, gin.H{"data": newCele})
}

func InfoForCele(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": dto.ToCeleDto(user.(model.TbCelebrity))}})
}

func UpdateInfo(ctx *gin.Context) {
	db := common.GetDB()
	var cele model.TbCelebrity
	err := ctx.ShouldBind(&cele)
	if err != nil {
		panic(err)
	}
	res := db.Model(&cele).Where("phone_number=?", cele.PhoneNumber).Updates(model.TbCelebrity{Name: cele.Name, PhoneNumber: cele.PhoneNumber, Sex: cele.Sex, Age: cele.Age, Intro: cele.Intro, Platform: cele.Platform, PlatformUrl: cele.PlatformUrl})
	fmt.Println(res)
}

func GetAll(ctx *gin.Context) {
	var user []model.TbCelebrity
	db := common.GetDB()
	db.Select("id,name,phone_number,email,avatar,sex,age,intro,platform,platform_url,credit_point").Find(&user)
	if user == nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "not find")
		return
	}
	response.Success(ctx, gin.H{"data": user})
}

func UpdateAvatar(ctx *gin.Context) {
	db := common.GetDB()

	file, _ := ctx.FormFile("file")
	tel := ctx.PostForm("tel")
	username := ctx.PostForm("username")

	url := ossUtils.OssUtils(file, username)
	tx := db.Table("tb_celebrity").Where("phone_number=?", tel).Update("avatar", url)
	if tx.Error != nil {
		fmt.Println("update fail")
		return
	}
	response.Success(ctx, gin.H{"url": url})
}
