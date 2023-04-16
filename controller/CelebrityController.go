package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"recommendation/common"
	"recommendation/dto"
	"recommendation/model"
	"recommendation/response"
	"strings"
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
		Id:          int(newId),
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
	response.Success(ctx, gin.H{"token": token}, "register successful")
}

func CeleLogin(ctx *gin.Context) {
	// connect database
	db := common.GetDB()

	// get parameter
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	// determine if the user is existed
	var cele model.TbCelebrity
	db.Where("username=?", username).First(&cele)
	if cele.Id == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 500, "mag": "user is not existed"})
		return
	}
	//determine if the password is correct
	if err := bcrypt.CompareHashAndPassword([]byte(cele.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 500, "msg": "password is not correct"})
		return
	}

	// distribute token
	token, err := common.ReleaseTokenForCele(cele)
	if err != nil {
		panic(err)
	}

	response.Success(ctx, gin.H{"token": token}, "login successful")
}

func GetUserInfo(ctx *gin.Context) {
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

	// connect database
	db := common.GetDB()
	var cele model.TbCelebrity
	db.Where("id=?", claims.UserId).First(&cele)

	newCele := model.TbCelebrity{
		Username:    cele.Username,
		PhoneNumber: cele.PhoneNumber,
		Email:       cele.Email,
		Name:        cele.Name,
		RealName:    cele.RealName,
		Sex:         cele.Sex,
		Age:         cele.Age,
		Intro:       cele.Intro,
		CreditPoint: cele.CreditPoint,
	}
	response.Success(ctx, gin.H{"data": newCele}, "successful")
}

func InfoForCele(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": dto.ToCeleDto(user.(model.TbCelebrity))}})
}

func UpdateInfo(ctx *gin.Context) {
	db := common.GetDB()
	var cele model.TbCelebrity
	err := ctx.ShouldBind(&cele)
	fmt.Println(err)
	fmt.Println(cele.Age)
	res := db.Model(&cele).Where("phone_number=?", cele.PhoneNumber).Updates(model.TbCelebrity{Name: cele.Name, PhoneNumber: cele.PhoneNumber, Sex: cele.Sex, Age: cele.Age, Intro: cele.Intro})
	fmt.Println(res)
}
