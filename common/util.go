package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"recommendation/model"
	"strings"
)

// IsGoodExist 商品是否存在
func IsGoodExist(name string) bool {
	db := GetDB()
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
