package bmiddle

import "github.com/gin-gonic/gin"

// AllowAllCorsHandler
//
// # 允许所有跨域
//
// 允许所有跨域请求
//
// # 参数
//   - c 	*gin.Context	上下文
func AllowAllCorsHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Next()
}

// AllowCurrentCorsHandler
//
// # 允许当前跨域
//
// 允许当前跨域请求
//
// # 参数
//   - c 	*gin.Context	上下文
func AllowCurrentCorsHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", c.GetHeader("Origin"))
	c.Next()
}
