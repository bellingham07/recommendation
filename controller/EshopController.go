package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"recommendation/common"
	"recommendation/database"
	"recommendation/dto"
	"recommendation/model"
	"recommendation/ossUtils"
	"recommendation/response"
)

func EshopRegister(ctx *gin.Context) {
	db := database.GetDB()

	//获取参数
	var eshop model.TbEshop
	err1 := ctx.ShouldBind(&eshop)
	if err1 != nil {
		panic(err1)
	}
	//数据验证
	if common.IsEmailExisted(db, eshop.Email) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "邮箱已存在")
		return
	}

	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(eshop.Password), bcrypt.DefaultCost)
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
		Username: eshop.Username,
		Name:     eshop.Name,
		Tel:      eshop.Tel,
		Password: string(hasedPassword),
	}

	db.Debug().Create(&newUser)

	//发放token
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		return
	}

	response.Success(ctx, gin.H{"token": token})
}

func EshopLogin(ctx *gin.Context) {
	db := database.GetDB()

	//get parameter
	var params model.TbEshop
	err1 := ctx.ShouldBind(&params)
	if err1 != nil {
		panic(err1)
	}
	//Determine if the user existed
	var eshop model.TbEshop
	db.Debug().Where("username=?", params.Username).First(&eshop)
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
	response.Success(ctx, gin.H{"token": token})
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": dto.ToUserDto(user.(model.TbEshop))}})
}

func GetAllUser(ctx *gin.Context) {
	db := database.GetDB()

	var eshop []model.TbEshop
	db.Select("id,name,tel,email,avatar,intro,platform,platform_url,credit_point,age,likes").Find(&eshop)

	response.Success(ctx, gin.H{"data": eshop})
}

func GetEshopInfo(ctx *gin.Context) {
	db := database.GetDB()

	var eshop model.TbEshop
	eshop.Id = common.GetId(ctx)
	db.Select("id", "username", "name", "tel", "email", "seller", "avatar", "intro", "platform_url", "platform", "age", "credit_point").Find(&eshop)
	if eshop.Tel == "" {
		response.Fail(ctx, nil)
		return
	}
	response.Success(ctx, gin.H{"data": eshop})
}

func UpdateEshop(ctx *gin.Context) {
	var eshop model.TbEshop
	db := database.GetDB()

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
	db := database.GetDB()

	file, _ := ctx.FormFile("file")
	tel := ctx.PostForm("tel")
	username := ctx.PostForm("username")

	fmt.Println("user:", username)
	url := ossUtils.OssUtils(file, username)
	fmt.Println("url:", url)

	tx := db.Debug().Table("tb_eshop").Where("tel=?", tel).Update("avatar", url)
	if tx.Error != nil {
		fmt.Println("update fail")
		return
	}
	response.Success(ctx, gin.H{"url": url})
}
