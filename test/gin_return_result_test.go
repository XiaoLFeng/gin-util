package test

import (
	"github.com/XiaoLFeng/go-general-utils/bcode"
	"github.com/XiaoLFeng/go-gin-util/berror"
	"github.com/XiaoLFeng/go-gin-util/bmiddle"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestGinReturnResult(t *testing.T) {
	r := gin.Default()
	r.Use(bmiddle.ReturnResultMiddleware())
	r.GET("/demo", func(c *gin.Context) {
		c.Error(berror.BusinessData{
			ErrorCode: bcode.Forbidden,
			ErrorMsg:  "测试",
		})
	})
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
