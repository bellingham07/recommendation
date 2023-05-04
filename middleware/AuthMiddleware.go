package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"recommendation/common"
	"recommendation/model"
)

// AuthMiddleware 用户信息验证
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//验证通过后获取claims中的userId
		userId := common.GetId(ctx)
		DB := common.GetDB()
		var user model.TbEshop
		DB.First(&user, userId)

		//用户信息失效
		if user.Id == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足3"})
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
		//验证通过后获取claims中的userId
		userId := common.GetId(ctx)
		DB := common.GetDB()
		var user model.TbCelebrity
		DB.First(&user, userId)

		//用户信息失效
		if user.Id == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足3"})
			ctx.Abort()
			return
		}

		//用户信息存在 将user信息写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}
