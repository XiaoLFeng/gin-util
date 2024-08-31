package bresult

import (
	"github.com/XiaoLFeng/go-general-utils/bmodel"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Ok
//
// # 成功响应
//
// 用于返回成功响应，返回成功响应码 200，默认消息 "Ok"；
// 不包含返回数据，只返回信息通知操作成功。
//
// # 参数
//   - c: gin.Context		请求上下文
//   - message: string		成功消息
func Ok(c *gin.Context, message string) {
	c.JSON(200, bmodel.BaseResponse{
		Output:       "Ok",
		Code:         200,
		Message:      message,
		ErrorMessage: nil,
		Data:         nil,
	})
}

// OkWithData
//
// # 成功响应
//
// 用于返回成功响应，返回成功响应码 200，默认消息 "Ok"；
// 包含返回数据，返回信息通知操作成功；
// 数据类型为 interface{} 内容需要自定义操作。
//
// # 参数
//   - c: gin.Context		请求上下文
//   - message: string		成功消息
//   - data: interface{}	成功数据
func OkWithData(c *gin.Context, message string, data interface{}) {
	c.JSON(200, bmodel.BaseResponse{
		Output:       "Ok",
		Code:         200,
		Message:      message,
		ErrorMessage: nil,
		Data:         data,
	})
}

// Error
//
// # 错误响应
//
// 用于返回错误响应，返回错误响应码，错误消息，错误数据；
// 不包含返回数据，只返回信息通知操作失败；
// code 需要自行定义，code 如果填写 HTTP 状态码，将会按照填写内容返回 HTTP 状态码，如果填写非 HTTP 状态码，将会按照前三位当做 HTTP 状态 码（建议对 code 起名时候进行数据安排）；
// 若 code 填写的状态码不符合 HTTP 状态码规范，将会尝试提取前三位转为 HTTP 状态码，若状态码转换失败，将会返回 500 状态码；
// output 为通用的输出消息，message 为通用的成功消息，errorMessage 为具体的错误消息。
//
// # 参数
//   - c: gin.Context		请求上下文
//   - code: uint			错误响应码
//   - output: string		输出消息
//   - message: string		成功消息
//   - errorMessage: string	错误消息
func Error(c *gin.Context, code uint, output, message, errorMessage string) {
	var statusCode int
	if code > 1000 {
		// 大于 1000 的错误码，选择前三位返回
		getCode, err := strconv.Atoi(strconv.Itoa(int(code))[:3])
		if err != nil {
			Error(c, 500, "Error", "Error", "Error code is invalid")
		}
		statusCode = getCode
	} else {
		statusCode = int(code)
	}
	c.JSON(statusCode, bmodel.BaseResponse{
		Output:       output,
		Code:         code,
		Message:      message,
		ErrorMessage: &errorMessage,
		Data:         nil,
	})
	c.Abort()
}

// ErrorWithData
//
// # 错误响应
//
// 用于返回错误响应，返回错误响应码，错误消息，错误数据；
// 包含返回数据，返回信息通知操作失败；
// code 需要自行定义，code 如果填写 HTTP 状态码，将会按照填写内容返回 HTTP 状态码，如果填写非 HTTP 状态码，将会按照前三位当做 HTTP 状态 码（建议对 code 起名时候进行数据安排）；
// 若 code 填写的状态码不符合 HTTP 状态码规范，将会尝试提取前三位转为 HTTP 状态码，若状态码转换失败，将会返回 500 状态码；
// output 为通用的输出消息，message 为通用的成功消息，errorMessage 为具体的错误消息，data 为返回数据。
// 数据类型为 interface{} 内容需要自定义操作。
//
// # 参数
//   - c: gin.Context		请求上下文
//   - code: uint			错误响应码
//   - output: string		输出消息
//   - message: string		成功消息
//   - errorMessage: string	错误消息
//   - data: interface{}	错误数据
func ErrorWithData(c *gin.Context, code uint, output, message, errorMessage string, data interface{}) {
	var statusCode int
	if code > 1000 {
		// 大于 1000 的错误码，选择前三位返回
		getCode, err := strconv.Atoi(strconv.Itoa(int(code))[:3])
		if err != nil {
			Error(c, 500, "Error", "Error", "Error code is invalid")
		}
		statusCode = getCode
	} else {
		statusCode = int(code)
	}
	c.JSON(statusCode, bmodel.BaseResponse{
		Output:       output,
		Code:         code,
		Message:      message,
		ErrorMessage: &errorMessage,
		Data:         data,
	})
	c.Abort()
}
