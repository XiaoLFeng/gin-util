package bcode

type ErrorCode struct {
	Output  string `json:"output"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// --------------------------------------------------------------------------------
// 错误码，用于返回错误信息；错误码包含输出，状态码，消息
// --------------------------------------------------------------------------------

var (
	SUCCESS                           = ErrorCode{Output: "Success", Code: 20000, Message: "成功"}
	BAD_REQUEST_INVALID_INPUT         = ErrorCode{Output: "Invalid", Code: 40001, Message: "无效的输入"}
	BAD_REQUEST_MISSING_PARAMETER     = ErrorCode{Output: "Invalid", Code: 40002, Message: "缺少必需的参数"}
	BAD_REQUEST_INVALID_PARAMETER     = ErrorCode{Output: "Invalid", Code: 40003, Message: "无效的参数值"}
	BAD_REQUEST_INVALID_JSON          = ErrorCode{Output: "Invalid", Code: 40004, Message: "无效的 JSON 格式"}
	UNAUTHORIZED                      = ErrorCode{Output: "Fail", Code: 40101, Message: "未授权"}
	UNAUTHORIZED_TOKEN_EXPIRED        = ErrorCode{Output: "Fail", Code: 40102, Message: "令牌已过期"}
	UNAUTHORIZED_INVALID_TOKEN        = ErrorCode{Output: "Fail", Code: 40103, Message: "无效的令牌"}
	UNAUTHORIZED_TOKEN_NOT_PROVIDED   = ErrorCode{Output: "Fail", Code: 40104, Message: "未提供令牌"}
	FORBIDDEN                         = ErrorCode{Output: "Deny", Code: 40301, Message: "禁止访问"}
	FORBIDDEN_ACCESS_DENIED           = ErrorCode{Output: "Deny", Code: 40302, Message: "访问被拒绝"}
	FORBIDDEN_INSUFFICIENT_PRIVILEGES = ErrorCode{Output: "Deny", Code: 40303, Message: "权限不足"}
	NOT_FOUND                         = ErrorCode{Output: "Error", Code: 40401, Message: "资源未找到"}
	NOT_FOUND_USER                    = ErrorCode{Output: "Error", Code: 40402, Message: "用户未找到"}
	NOT_FOUND_PAGE                    = ErrorCode{Output: "Error", Code: 40403, Message: "页面未找到"}
	METHOD_NOT_ALLOWED                = ErrorCode{Output: "Error", Code: 40501, Message: "方法不允许"}
	REQUEST_TIMEOUT                   = ErrorCode{Output: "Timeout", Code: 40801, Message: "请求超时"}
	CONFLICT                          = ErrorCode{Output: "Conflict", Code: 40901, Message: "冲突"}
	CONFLICT_RESOURCE_ALREADY_EXISTS  = ErrorCode{Output: "Conflict", Code: 40902, Message: "资源已存在"}
	GONE                              = ErrorCode{Output: "Gone", Code: 41001, Message: "资源已删除"}
	PAYLOAD_TOO_LARGE                 = ErrorCode{Output: "Error", Code: 41301, Message: "请求实体过大"}
	UNSUPPORTED_MEDIA_TYPE            = ErrorCode{Output: "Error", Code: 41501, Message: "不支持的媒体类型"}
	UNPROCESSABLE_ENTITY              = ErrorCode{Output: "Invalid", Code: 42201, Message: "不可处理的实体"}
	UNPROCESSABLE_ENTITY_VALIDATION   = ErrorCode{Output: "Invalid", Code: 42202, Message: "验证错误"}
	TOO_MANY_REQUESTS                 = ErrorCode{Output: "Limit", Code: 42901, Message: "请求过多"}
	SERVER_INTERNAL_ERROR             = ErrorCode{Output: "Error", Code: 50001, Message: "服务器内部错误"}
	SERVER_DATABASE_ERROR             = ErrorCode{Output: "Error", Code: 50002, Message: "数据库错误"}
	SERVER_CONFIGURATION_ERROR        = ErrorCode{Output: "Error", Code: 50003, Message: "配置错误"}
	BAD_GATEWAY                       = ErrorCode{Output: "Error", Code: 50201, Message: "错误的网关"}
	SERVICE_UNAVAILABLE               = ErrorCode{Output: "Error", Code: 50301, Message: "服务不可用"}
	SERVICE_UNDER_MAINTENANCE         = ErrorCode{Output: "Error", Code: 50302, Message: "服务维护中"}
	GATEWAY_TIMEOUT                   = ErrorCode{Output: "Timeout", Code: 50401, Message: "网关超时"}
)
