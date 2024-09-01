package berror

import "github.com/XiaoLFeng/go-general-utils/bcode"

type BusinessError struct {
	ErrorCode bcode.ErrorCode
	ErrorMsg  string
	Data      interface{}
}

type BusinessErrImpl interface {
	Error() string
	GetCode() uint
	GetMessage() string
	GetErrorMessage() string
}

// Error
//
// # 获取错误信息
func (e BusinessError) Error() string {
	return e.ErrorMsg
}

func (e BusinessError) GetCode() uint {
	return e.ErrorCode.Code
}

func New(berror bcode.ErrorCode, errMsg string) error {
	return BusinessError{
		ErrorCode: berror,
		ErrorMsg:  errMsg,
	}
}

func NewWithData(berror bcode.ErrorCode, errMsg string, data interface{}) error {
	return BusinessError{
		ErrorCode: berror,
		ErrorMsg:  errMsg,
		Data:      data,
	}
}
