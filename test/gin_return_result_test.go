package test

import (
	"errors"
	"github.com/XiaoLFeng/go-general-utils/bcode"
	"github.com/XiaoLFeng/go-gin-util/berror"
	"github.com/XiaoLFeng/go-gin-util/bmiddle"
	"github.com/XiaoLFeng/go-gin-util/bresult"
	"github.com/gin-gonic/gin"
	"testing"
)

// TestGinReturnResult
//
// # 测试 Gin 框架的标准化返回
//
// 对 Gin 框架测试标准返回，检查正常的返回，错误的返回，以及被拦截的返回信息是否符合预期设计规范
func TestGinReturnResult(t *testing.T) {
	r := gin.Default()
	r.Use(bmiddle.ReturnResultMiddleware())
	r.GET("/success", func(c *gin.Context) {
		bresult.Ok(c, "操作完成")
	})
	r.GET("/fail", func(c *gin.Context) {
		_ = c.Error(berror.BusinessError{
			ErrorCode: bcode.Forbidden,
			ErrorMsg:  "被拦截测试",
		})
	})
	r.GET("/success-but-error", func(c *gin.Context) {
		_ = c.Error(berror.BusinessError{
			ErrorCode: bcode.Forbidden,
			ErrorMsg:  "返回成功后，中间出现拦截操作",
		})
		bresult.Ok(c, "操作完成")
	})
	r.GET("/empty", func(c *gin.Context) {
		return
	})
	r.GET("/custom-error", func(c *gin.Context) {
		_ = c.Error(errors.New("自定义 Error 抛出"))
	})
	r.NoRoute(bmiddle.NoRouteMiddleware())
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
