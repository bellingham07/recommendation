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
	"recommendation/redis"
	"recommendation/response"
	"time"
)

func CeleRegister(ctx *gin.Context) {
	// connect database
	db := database.GetDB()

	// get register params
	password := ctx.PostForm("password")
	email := ctx.PostForm("email")
	MailCode := ctx.PostForm("mail_code")
	fmt.Println("mail code", MailCode)

	// 邮箱是否有效
	if common.IsEmailExisted(db, email) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "邮箱已存在")
		return
	}

	// 验证码是否匹配
	if !common.IsEmailCodeValid(MailCode, email) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "验证码错误")
		return
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "server error")
		return
	}

	// use snowflake generate a new id
	node, err := common.NewWorker(1)
	if err != nil {
		panic(err)
	}

	newId := node.GetId()

	// create a new entity to save the info
	newCele := model.TbCelebrity{
		Id:       newId,
		Email:    email,
		Name:     "新用户",
		Password: string(hashedPassword),
	}

	// save to database
	db.Debug().Create(&newCele)

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
	db := database.GetDB()

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
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 500, "msg": "user is not existed"})
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
	db := database.GetDB()

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
	db := database.GetDB()

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
	db := database.GetDB()

	db.Select("id,name,phone_number,email,avatar,sex,age,intro,platform,platform_url,credit_point").Find(&user)
	if user == nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "not find")
		return
	}
	response.Success(ctx, gin.H{"data": user})
}

func UpdateAvatar(ctx *gin.Context) {
	db := database.GetDB()

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

func IsLiked(c *gin.Context) {
	// 获取被点赞对象id
	id := c.PostForm("likedId")
	// 获取当前登录用户uid
	uid := common.GetId(c)
	// 判断当前被点赞对象是否被当前用户点赞 select * form tb_like where like_id= iid and liked_id=id
	flag := isLiked(uid, id)
	if flag {
		c.JSON(200, gin.H{"data": true})
	} else {
		// 不存在返回false
		c.JSON(200, gin.H{"data": false})
	}
}

func Like(c *gin.Context) {
	db := database.GetDB()

	// 获取被点赞对象id
	id := c.PostForm("id")
	// 获取当前登录用户id
	uid := common.GetId(c)
	// 查询是否被点赞
	flag := isLiked(uid, id)
	if !flag {
		node, err := common.NewWorker(1)
		if err != nil {
			panic(err)
		}

		newId := node.GetId()
		// 没有，添加点赞信息到点赞表
		like := model.TbLike{
			Id:      newId,
			LikedId: id,
			LikeId:  uid,
		}
		db.Debug().Save(&like)
		// 点赞数加一
		db.Debug().Exec("UPDATE tb_eshop SET likes=likes+1 WHERE id=?", id)
		c.JSON(200, gin.H{"data": true})
	} else {
		// 有，取消点赞 删除点赞表点赞信息
		db.Debug().Where("like_id=?", uid).Where("liked_id=?", id).Delete(&model.TbLike{})
		// 点赞数减一
		db.Debug().Exec("UPDATE tb_eshop SET likes=likes-1 WHERE id=?", id)
		c.JSON(200, gin.H{"data": false})
	}
}

func isLiked(likeId string, likedId string) bool {
	db := database.GetDB()

	var like model.TbLike
	tx := db.Debug().Where("like_id=?", likeId).Where("liked_id=?", likedId).Find(&like)
	// 存在返回true
	if tx.RowsAffected != 0 {
		return true
	}
	return false
}

func SendMailCode(c *gin.Context) {
	email := c.PostForm("email")
	code := common.RandCode()
	// 验证码存入redis
	rds := redis.GetRedis()
	rds.Set(c, email, code, 5*time.Minute)
	// 发送验证码
	err := common.MailSendCode(email, code)
	if err != nil {
		return
	}
}

func User(c *gin.Context) {
	var req model.ReqListUsers
	err := c.ShouldBindQuery(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	fmt.Println(req)
}
