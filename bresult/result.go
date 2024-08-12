package utils

import (
	"errors"
	"github.com/XiaoLFeng/gin-util/bcode"
	"github.com/XiaoLFeng/gin-util/berror"
	"github.com/gin-gonic/gin"
)

// Success
//
// # 成功响应
//
// 成功响应，返回成功响应码，成功消息，成功数据
//
// # 参数
//   - c: gin.Context
//   - message: string		成功消息
//   - data: interface{}	成功数据
func Success(c *gin.Context, message string, data interface{}) {
	_ = c.Error(&berror.ThrowError{
		ErrorCode: bcode.SUCCESS,
		Err:       errors.New(message),
		Data:      data,
	})
}

// SuccessNoData
//
// # 成功响应
//
// 成功响应，返回成功响应码，成功消息，无数据
//
// # 参数
//   - c: gin.Context
//   - message: string		成功消息
func SuccessNoData(c *gin.Context, message string) {
	_ = c.Error(&berror.ThrowError{
		ErrorCode: bcode.SUCCESS,
		Err:       errors.New(message),
		Data:      nil,
	})
}

// NewError
//
// # 错误响应
//
// 错误响应，返回错误响应码，错误消息，错误数据
//
// # 参数
//   - c: gin.Context
//   - errorCode: ErrorCode	错误响应码
//   - err: error				错误消息
func NewError(c *gin.Context, errorCode bcode.ErrorCode, err error, data interface{}) {
	_ = c.Error(&berror.ThrowError{
		ErrorCode: errorCode,
		Err:       err,
		Data:      data,
	})
}

// NewErrorHasMessage
//
// # 错误响应
//
// 错误响应，返回错误响应码，错误消息，错误数据
//
// # 参数
//   - c: gin.Context
//   - errorCode: ErrorCode	错误响应码
//   - message: string			错误消息
//   - data: interface{}		错误数据
func NewErrorHasMessage(c *gin.Context, errorCode bcode.ErrorCode, message string, data interface{}) {
	_ = c.Error(&berror.ThrowError{
		ErrorCode: errorCode,
		Err:       errors.New(message),
		Data:      data,
	})
}
