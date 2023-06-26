package middle

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CORSMiddleware 使用gin解决跨域问题
func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")  //允许访问的域名 *表示所有的都可以访问
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")   //缓存时间
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "*") //可以通过访问的方法 get post ...  *表示允许所有方法
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*") //允许请求带的header头信息
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(200)
		} else {
			ctx.Next()
		}
	}
}
