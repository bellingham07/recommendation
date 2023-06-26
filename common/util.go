package common

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"recommendation/database"
	"recommendation/model"
	"strings"
	"time"
)

// IsGoodExist 商品是否存在
func IsGoodExist(name string) bool {
	db := database.GetDB()

	var good model.TbGood
	db.Where("name=?", name).First(&good)
	if good.Id != "" {
		return true
	}
	return false
}

// GetId 从token中获取用户id
func GetId(c *gin.Context) string {
	// get authorization header
	tokenString := c.GetHeader("Authorization")
	// validate token format
	if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
		c.Abort()
	}
	// bearer's length is 6
	tokenString = tokenString[6:]

	token, claims, err := ParseToken(tokenString)
	// if parsing failure or token is invalid
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 402, "msg": "权限不足"})
		c.Abort()
	}
	return claims.UserId
}

// RandCode 随机成验证码码
func RandCode() string {
	s := "1234567890"
	code := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 6; i++ {
		code += string(s[rand.Intn(len(s))])
	}
	return code
}
