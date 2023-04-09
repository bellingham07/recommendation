package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"recommendation/common"
	"recommendation/model"
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

	response.Success(ctx, gin.H{"data": token}, "register successful")
}
