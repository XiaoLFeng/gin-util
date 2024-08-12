package bmodel

// BaseResponse
//
// # 基础响应
//
// 用于返回基础响应，该响应包含输出，状态码，消息，错误消息，数据
type BaseResponse struct {
	Output       string      `json:"output" binding:"required"`
	Code         int         `json:"code" binding:"required"`
	Message      string      `json:"message" binding:"required"`
	ErrorMessage *string     `json:"error_message"`
	Data         interface{} `json:"data"`
}
