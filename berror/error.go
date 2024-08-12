package berror

import (
	"github.com/XiaoLFeng/gin-util/bcode"
)

type BusinessError interface {
	Error() string
	GetErrorCode() bcode.ErrorCode
	GetData() interface{}
}

// GetError
//
// # 获取错误消息
//
// 获取错误消息, 返回错误消息
//
// # 返回
//   - string	错误消息
func (e *ThrowError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return ""
}

// GetErrorCode
//
// # 获取错误响应码
//
// 获取错误响应码, 返回错误响应码
//
// # 返回
//   - ErrorCode	错误响应码
func (e *ThrowError) GetErrorCode() bcode.ErrorCode {
	return e.ErrorCode
}

// GetData
//
// # 获取错误数据
//
// 获取错误数据, 返回错误数据
//
// # 返回
//   - interface{}	错误数据
func (e *ThrowError) GetData() interface{} {
	return e.Data
}
