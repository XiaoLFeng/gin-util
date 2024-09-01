package bmiddle

import (
	"bytes"
	"encoding/json"
	"github.com/XiaoLFeng/go-general-utils/bcode"
	"github.com/XiaoLFeng/go-general-utils/bmodel"
	"github.com/XiaoLFeng/go-gin-util/bresult"
	"github.com/XiaoLFeng/go-gin-util/bstruct"
	"github.com/gin-gonic/gin"
)

// NoRouteMiddleware
//
// # 页面不存在中间件
//
// 用于处理页面不存在的中间件，返回 404 页面不存在的错误信息；
// 该方法的使用模式如下：
//
//	r.NoRoute(bmiddle.NoRouteMiddleware())
//
// 该方法会在页面路由不存在的时候，返回 404 页面不存在的错误信息；
func NoRouteMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		w := &bstruct.CustomWriter{Body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = w

		c.Writer.WriteHeader(int(bcode.NotFoundPage.Code / 100))
		errMsg := c.Request.RequestURI + " 页面不存在"
		marshal, err := json.Marshal(bmodel.BaseResponse{
			Output:       bcode.NotFoundPage.Output,
			Code:         bcode.NotFoundPage.Code,
			Message:      bcode.NotFoundPage.Message,
			ErrorMessage: &errMsg,
		})
		if err != nil {
			bresult.Error(c, bcode.BadRequestInvalidJson.Code, bcode.BadRequestInvalidJson.Output, bcode.BadRequestInvalidJson.Message, err.Error())
		}
		_, _ = c.Writer.Write(marshal)
		_, _ = w.WriteConfirm()

		c.Abort()
		return
	}
}
