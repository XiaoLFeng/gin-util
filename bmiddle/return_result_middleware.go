package bmiddle

import (
	"bytes"
	"errors"
	"github.com/XiaoLFeng/go-general-utils/bcode"
	"github.com/XiaoLFeng/go-gin-util/berror"
	"github.com/XiaoLFeng/go-gin-util/bresult"
	"github.com/XiaoLFeng/go-gin-util/bstruct"
	"github.com/gin-gonic/gin"
)

// ReturnResultMiddleware
//
// # 返回结果中间件
//
// 用于处理返回结果的中间件，返回结果信息；
// 该方法的使用模式如下：
//
//	# Ok result(demo)
//	bresult.Ok(c, "操作完成")
//	# Fail result(demo)
//	_ = c.Error(berror.New(bcode.ServerInternalError, "服务器内部错误(测试错误)"))
//	return
//	# or
//	_ = c.Error(berror.BusinessError{
//		ErrorCode: bcode.ServerInternalError,
//		ErrorMsg:  "服务器内部错误(测试错误)",
//	})
//	return
//
// 该方法检测出现错误的时候是捕获 gin 中 Errors 列表信息进行判断，若 gin 中 Errors 列表并未包含错误信息，将会返回 bresult.Ok 的结果集；
// 若在方法中不包含 c.Error(err error) 也不包含 bresult.Ok，系统将会报出 bcode.NotImplemented 表示不包含任何输出结果集；
// 若在方法中包含 c.Error(err error) 也包含 bresult.Ok，两者存在的时候将会报出 bcode.Conflict 表示信息冲突(一次只可输出一个结果信息)；
// 若 c.Error(err error) 中的 error 非 berror.BusinessError ，将会默认显示 bcode.NotImplemented，errorMessage 会显示自定义的 err.Error()。
func ReturnResultMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_ = c.Error(berror.BusinessError{
			ErrorCode: bcode.ServerInternalError,
			ErrorMsg:  "服务器内部错误(测试错误)",
		})
		w := &bstruct.CustomWriter{Body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = w

		// 执行请求
		c.Next()

		// 若直接输入了结果直接返回结果信息
		if !c.Writer.Written() {
			// 发生错误进行拦截
			if len(c.Errors) > 0 {
				// 检查是否提前输出了数据
				if len(w.Body.Bytes()) > 0 {
					w.Body.Reset()
					bresult.Error(c, bcode.Conflict.Code, bcode.Conflict.Output, bcode.Conflict.Message, "不被允许的操作，出现错误需要返回，不许输出结果")
					_, _ = w.WriteConfirm()
					return
				}
				// 清空 body 数据
				w.Body.Reset()
				// 对 error 进行类型判断
				var err berror.BusinessError
				if errors.As(c.Errors.Last(), &err) {
					bresult.Error(c, err.ErrorCode.Code, err.ErrorCode.Output, err.ErrorCode.Message, err.ErrorMsg)
				} else {
					bresult.Error(c, bcode.NotImplemented.Code, bcode.NotImplemented.Output, bcode.NotImplemented.Message, c.Errors.Last().Error())
				}
				_, _ = w.WriteConfirm()
			} else {
				// 检查 body 是否有数据
				if len(w.Body.Bytes()) <= 0 {
					// 检查页面是否存在
					w.Body.Reset()
					bresult.Error(c, bcode.NotImplemented.Code, bcode.NotImplemented.Output, bcode.NotImplemented.Message, "没有返回数据")
					_, _ = w.WriteConfirm()
				}
			}
		}
	}
}
