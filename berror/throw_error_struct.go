package berror

import "github.com/XiaoLFeng/gin-util/bcode"

type ThrowError struct {
	ErrorCode bcode.ErrorCode
	Err       error
	Data      interface{}
}
