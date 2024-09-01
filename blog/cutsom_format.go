package blog

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
)

// 定义颜色常量
const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
)

// CustomFormatter
//
// # 自定义日志格式化
//
// 自定义日志格式化，实现 logrus.Formatter 接口；
// 包含颜色
type CustomFormatter struct{}

// CustomFormatterNoColor
//
// # 自定义日志格式化
//
// 自定义日志格式化，实现 logrus.Formatter 接口；
// 不包含颜色
type CustomFormatterNoColor struct{}

// Format
//
// # 格式化日志
//
// 格式化日志，实现 logrus.Formatter 接口；
//
//	格式如下：[LOG] 2024/09/01 - 15:00:00 | INFO | <message>
//
// 该格式包含了日志的时间、等级、消息内容，其中等级使用颜色区分
//
// # 参数
//   - entry: 日志实体
//
// # 返回
//   - []byte: 格式化后的日志信息
func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 获取当前时间，并格式化
	timestamp := entry.Time.Format("2006/01/02 - 15:04:05")

	// 获取日志等级并赋予颜色
	var levelColor string
	switch entry.Level {
	case logrus.DebugLevel:
		levelColor = Blue
	case logrus.InfoLevel:
		levelColor = Green
	case logrus.WarnLevel:
		levelColor = Yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = Red
	case logrus.TraceLevel:
		levelColor = Magenta
	default:
		levelColor = White
	}

	level := entry.Level.String()
	level = strings.ToUpper(level[:4]) // 取前4个字符

	// 格式化最终的日志消息
	logMessage := fmt.Sprintf("%s[LOG] %s | %s%s%s | %s\n",
		Reset, timestamp, levelColor, level, Reset, entry.Message)
	return []byte(logMessage), nil
}

// Format
//
// # 格式化日志
//
// 格式化日志，实现 logrus.Formatter 接口；
//
//	格式如下：[LOG] 2024/09/01 - 15:00:00 | INFO | <message>
//
// 该格式包含了日志的时间、等级、消息内容，其中等级不使用颜色区分
//
// # 参数
//   - entry: 日志实体
//
// # 返回
//   - []byte: 格式化后的日志信息
func (f *CustomFormatterNoColor) Format(entry *logrus.Entry) ([]byte, error) {
	// 获取当前时间，并格式化
	timestamp := entry.Time.Format("2006/01/02 - 15:04:05")

	// 获取日志等级并赋予颜色
	level := entry.Level.String()
	level = strings.ToUpper(level[:4]) // 取前4个字符

	// 格式化最终的日志消息
	logMessage := fmt.Sprintf("[LOG] %s | %s | %s\n", timestamp, level, entry.Message)
	return []byte(logMessage), nil
}
