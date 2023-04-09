package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"recommendation/common"
	"recommendation/model"
	"strings"
)

// 用户信息验证
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取authorization header
		tokenString := ctx.GetHeader("Authorization")

		//validate token format
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		tokenString = tokenString[7:] //Bearer 占七位

		token, claims, err := common.ParseToken(tokenString)
		//解析失败或者token无效
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		//验证通过后获取claims中的userId
		userId := claims.UserId
		DB := common.GetDB()
		var user model.TbEshop
		DB.First(&user, userId)

		//用户信息失效
		if user.Id == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		//用户信息存在 将user信息写入上下文
		ctx.Set("user", user)

		ctx.Next()
	}
}

func AuthMiddlewareForCele() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取authorization header
		tokenString := ctx.GetHeader("Authorization")

		//validate token format
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		tokenString = tokenString[7:] //Bearer 占七位

		token, claims, err := common.ParseToken(tokenString)
		//解析失败或者token无效
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		//验证通过后获取claims中的userId
		userId := claims.UserId
		DB := common.GetDB()
		var user model.TbCelebrity
		DB.First(&user, userId)

		//用户信息失效
		if user.Id == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		//用户信息存在 将user信息写入上下文
		ctx.Set("user", user)

		ctx.Next()
	}
}
