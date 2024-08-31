package bmiddle

import (
	"bytes"
	"errors"
	"github.com/XiaoLFeng/go-general-utils/bcode"
	"github.com/XiaoLFeng/go-gin-util/berror"
	"github.com/XiaoLFeng/go-gin-util/bresult"
	"github.com/gin-gonic/gin"
)

type customWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w customWriter) Write(b []byte) (int, error) {
	// 将写入的数据缓存到 body 中
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func ReturnResultMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 执行请求
		c.Next()

		// 发生错误进行拦截
		if len(c.Errors) > 0 {
			// 清空 body 数据
			c.Writer = &customWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
			// 对 error 进行类型判断
			var err berror.BusinessData
			if errors.As(c.Errors.Last(), &err) {
				bresult.Error(c, err.ErrorCode.Code, err.ErrorCode.Output, err.ErrorCode.Message, err.ErrorMsg)
			} else {
				bresult.Error(c, bcode.NotImplemented.Code, bcode.NotImplemented.Output, bcode.NotImplemented.Message, c.Errors.Last().Error())
			}
		} else {
			// 检查 body 是否有数据
			if !c.Writer.Written() {
				bresult.Error(c, bcode.NotImplemented.Code, bcode.NotImplemented.Output, bcode.NotImplemented.Message, "No data returned")
			}
		}
	}
}
