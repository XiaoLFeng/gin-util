package bmiddle

import (
	"github.com/XiaoLFeng/gin-util/bcode"
	"github.com/XiaoLFeng/gin-util/berror"
	"github.com/XiaoLFeng/gin-util/bmodel"
	"github.com/gin-gonic/gin"
)

// NormalizedOutputHandler
//
// # 标准输出处理
//
// 标准输出处理，用于处理业务异常，统一输出
// 业务异常处理
//
// # 参数
//   - c 	*gin.Context	上下文
func NormalizedOutputHandler(c *gin.Context) {
	c.Next()

	// 业务异常处理
	if len(c.Errors) > 1 {
		getErrMessage := "超出输出异常处理范围，请检查输出是否出现冲突"
		c.JSON(500, bmodel.BaseResponse{
			Output:       bcode.SERVER_INTERNAL_ERROR.Output,
			Code:         bcode.SERVER_INTERNAL_ERROR.Code,
			Message:      bcode.SERVER_INTERNAL_ERROR.Message,
			ErrorMessage: &getErrMessage,
		})
		c.Abort()
		return
	}
	thrError := berror.ConversionError(c.Errors[0])
	if thrError != nil {
		getError := thrError.Error()
		// 处理捕获的响应内容
		getCode := thrError.GetErrorCode().Code / 100
		c.JSON(getCode, bmodel.BaseResponse{
			Output:       thrError.GetErrorCode().Output,
			Code:         thrError.GetErrorCode().Code,
			Message:      thrError.GetErrorCode().Message,
			ErrorMessage: &getError,
			Data:         thrError.GetData(),
		})
	}
	c.Abort()
}
