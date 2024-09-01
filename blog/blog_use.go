package blog

import "github.com/sirupsen/logrus"

// Trace
//
// # 跟踪日志
//
// 用于记录跟踪日志信息，该方法会记录日志信息到日志文件中；
// 该方法会记录日志信息到日志文件中；
// 等级为 Trace，用于记录跟踪日志信息。
//
// # 参数
//   - args: 日志信息
func Trace(args ...interface{}) {
	logrus.Trace(args)
}

// Tracef
//
// # 格式化跟踪日志
//
// 用于记录格式化跟踪日志信息，该方法会记录日志信息到日志文件中；
// 该方法会记录日志信息到日志文件中；
// 等级为 Trace，用于记录跟踪日志信息。
//
// # 参数
//   - format: 格式化字符串
//   - args: 格式化参数
func Tracef(format string, args ...interface{}) {
	logrus.Tracef(format, args)
}

// Debug
//
// # 调试日志
//
// 用于记录调试日志信息，该方法会记录日志信息到日志文件中；
// 该方法会记录日志信息到日志文件中；
// 等级为 Debug，用于记录调试日志信息。
//
// # 参数
//   - args: 日志信息
func Debug(args ...interface{}) {
	logrus.Debug(args)
}

// Debugf
//
// # 格式化调试日志
//
// 用于记录格式化调试日志信息，该方法会记录日志信息到日志文件中；
// 该方法会记录日志信息到日志文件中；
// 等级为 Debug，用于记录调试日志信息。
//
// # 参数
//   - format: 格式化字符串
//   - args: 格式化参数
func Debugf(format string, args ...interface{}) {
	logrus.Debugf(format, args)
}

// Info
//
// # 信息日志
//
// 用于记录信息日志信息，该方法会记录日志信息到日志文件中；
// 该方法会记录日志信息到日志文件中；
// 等级为 Info，用于记录信息日志信息。
//
// # 参数
//   - args: 日志信息
func Info(args ...interface{}) {
	logrus.Info(args)
}

// Infof
//
// # 格式化信息日志
//
// 用于记录格式化信息日志信息，该方法会记录日志信息到日志文件中；
// 该方法会记录日志信息到日志文件中；
// 等级为 Info，用于记录信息日志信息。
//
// # 参数
//   - format: 格式化字符串
//   - args: 格式化参数
func Infof(format string, args ...interface{}) {
	logrus.Infof(format, args)
}

// Warn
//
// # 警告日志
//
// 用于记录警告日志信息，该方法会记录日志信息到日志文件中；
// 该方法会记录日志信息到日志文件中；
// 等级为 Warn，用于记录警告日志信息。
//
// # 参数
//   - args: 日志信息
func Warn(args ...interface{}) {
	logrus.Warn(args)
}

// Warnf
//
// # 格式化警告日志
//
// 用于记录格式化警告日志信息，该方法会记录日志信息到日志文件中；
// 该方法会记录日志信息到日志文件中；
// 等级为 Warn，用于记录警告日志信息。
//
// # 参数
//   - format: 格式化字符串
//   - args: 格式化参数
func Warnf(format string, args ...interface{}) {
	logrus.Warnf(format, args)
}

// Error
//
// # 错误日志
//
// 用于记录错误日志信息，该方法会记录日志信息到日志文件中；
// 该方法会记录日志信息到日志文件中；
// 等级为 Error，用于记录错误日志信息。
//
// # 参数
//   - args: 日志信息
func Error(args ...interface{}) {
	logrus.Error(args)
}

// Errorf
//
// # 格式化错误日志
//
// 用于记录格式化错误日志信息，该方法会记录日志信息到日志文件中；
// 该方法会记录日志信息到日志文件中；
// 等级为 Error，用于记录错误日志信息。
//
// # 参数
//   - format: 格式化字符串
//   - args: 格式化参数
func Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args)
}

// Fatal
//
// # 致命日志
//
// 用于记录致命日志信息，该方法会记录日志信息到日志文件中；
// 该方法会记录日志信息到日志文件中；
// 等级为 Fatal，用于记录致命日志信息。
//
// # 参数
//   - args: 日志信息
func Fatal(args ...interface{}) {
	logrus.Fatal(args)
}

// Fatalf
//
// # 格式化致命日志
//
// 用于记录格式化致命日志信息，该方法会记录日志信息到日志文件中；
// 该方法会记录日志信息到日志文件中；
// 等级为 Fatal，用于记录致命日志信息。
//
// # 参数
//   - format: 格式化字符串
//   - args: 格式化参数
func Fatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args)
}

// Panic
//
// # 异常日志
//
// 用于记录异常日志信息，该方法会记录日志信息到日志文件中；
// 该方法会记录日志信息到日志文件中；
// 等级为 Panic，用于记录异常日志信息。
//
// # 参数
//   - args: 日志信息
func Panic(args ...interface{}) {
	logrus.Panic(args)
}

// Panicf
//
// # 格式化异常日志
//
// 用于记录格式化异常日志信息，该方法会记录日志信息到日志文件中；
// 该方法会记录日志信息到日志文件中；
// 等级为 Panic，用于记录异常日志信息。
//
// # 参数
//   - format: 格式化字符串
//   - args: 格式化参数
func Panicf(format string, args ...interface{}) {
	logrus.Panicf(format, args)
}
