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
func Trace(types string, args ...interface{}) {
	var getArg string
	for _, arg := range args {
		getArg += arg.(string)
	}
	logrus.Trace(types+" | ", getArg)
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
func Tracef(types string, format string, args ...interface{}) {
	logrus.Tracef(types+" | "+format, args...)
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
func Debug(types string, args ...interface{}) {
	var getArg string
	for _, arg := range args {
		getArg += arg.(string)
	}
	logrus.Debug(types+" | ", getArg)
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
func Debugf(types string, format string, args ...interface{}) {
	logrus.Debugf(types+" | "+format, args...)
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
func Info(types string, args ...interface{}) {
	var getArg string
	for _, arg := range args {
		getArg += arg.(string)
	}
	logrus.Info(types+" | ", getArg)
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
func Infof(types string, format string, args ...interface{}) {
	logrus.Infof(types+" | "+format, args...)
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
func Warn(types string, args ...interface{}) {
	var getArg string
	for _, arg := range args {
		getArg += arg.(string)
	}
	logrus.Warn(types+" | ", getArg)
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
func Warnf(types string, format string, args ...interface{}) {
	logrus.Warnf(types+" | "+format, args...)
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
func Error(types string, args ...interface{}) {
	var getArg string
	for _, arg := range args {
		getArg += arg.(string)
	}
	logrus.Error(types+" | ", getArg)
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
func Errorf(types string, format string, args ...interface{}) {
	logrus.Errorf(types+" | "+format, args...)
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
func Fatal(types string, args ...interface{}) {
	var getArg string
	for _, arg := range args {
		getArg += arg.(string)
	}
	logrus.Fatal(types+" | ", getArg)
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
func Fatalf(types string, format string, args ...interface{}) {
	logrus.Fatalf(types+" | "+format, args...)
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
func Panic(types string, args ...interface{}) {
	var getArg string
	for _, arg := range args {
		getArg += arg.(string)
	}
	logrus.Panic(types+" | ", getArg)
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
func Panicf(types string, format string, args ...interface{}) {
	logrus.Panicf(types+" | "+format, args...)
}

// Print
//
// # 打印日志
//
// 用于打印日志信息，该方法会记录日志信息到日志文件中；
// 该方法会记录日志信息到日志文件中；
// 等级为 Info，用于记录日志信息。
//
// # 参数
//   - args: 日志信息
func Print(types string, args ...interface{}) {
	var getArg string
	for _, arg := range args {
		getArg += arg.(string)
	}
	logrus.Print(types+" | ", getArg)
}

// Printf
//
// # 格式化打印日志
//
// 用于打印格式化日志信息，该方法会记录日志信息到日志文件中；
// 该方法会记录日志信息到日志文件中；
// 等级为 Info，用于记录日志信息。
//
// # 参数
//   - format: 格式化字符串
//   - args: 格式化参数
func Printf(types string, format string, args ...interface{}) {
	logrus.Printf(types+" | "+format, args...)
}
