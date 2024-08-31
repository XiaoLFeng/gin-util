package berror

import "github.com/XiaoLFeng/go-general-utils/bcode"

type BusinessData struct {
	ErrorCode bcode.ErrorCode
	Msg       string
	ErrorMsg  string
	Data      interface{}
}

type BusinessError interface {
	Error() string
	GetCode() uint
	GetMessage() string
	GetErrorMessage() string
}

// Error
//
// # 获取错误信息
func (e BusinessData) Error() string {
	return e.ErrorMsg
}

func (e BusinessData) GetCode() uint {
	return e.ErrorCode.Code
}

func New(berror bcode.ErrorCode, msg, errMsg string) error {
	return BusinessData{
		ErrorCode: berror,
		Msg:       msg,
		ErrorMsg:  errMsg,
	}
}

func NewWithData(berror bcode.ErrorCode, msg, errMsg string, data interface{}) error {
	return BusinessData{
		ErrorCode: berror,
		Msg:       msg,
		ErrorMsg:  errMsg,
		Data:      data,
	}
}
