package bmiddle

import "github.com/gin-gonic/gin"

// CrossDomainClearingMiddleware
//
// # 跨域清理中间件
//
// 用于清理跨域请求的中间件，清理跨域请求的请求头信息，允许所有的跨域请求；
// 该方法应该在测试的时候执行，不建议使用在生产模式下
//
// # 返回
//   - gin.HandlerFunc		返回 gin 的中间件处理方法
func CrossDomainClearingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}
