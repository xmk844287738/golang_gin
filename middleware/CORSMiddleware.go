package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 处理前台访问时发生的跨域问题
func CORSMiddleware() (gin.HandlerFunc)  {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(200)
		}else {
			c.Next()
		}
	}
}
