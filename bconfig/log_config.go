package bconfig

import (
	"github.com/XiaoLFeng/go-gin-util/blog"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

// LogConfiguration
//
// # 日志配置
//
// 对日志信息进行配置操作，包括日志文件路径、是否启用颜色
//
// # 参数
//   - path: 日志文件路径
//   - enableColor: 是否启用颜色
func LogConfiguration(path string, enableColor bool, level logrus.Level) {
	// GIN 系统设置日志输出
	log, _ := os.Create(path)
	if !enableColor {
		gin.DisableConsoleColor()
	}
	gin.DefaultWriter = io.MultiWriter(log, os.Stdout)

	// LOG 创建一个新的 logger 实例
	logrus.SetFormatter(new(blog.CustomFormatter))
	logrus.SetLevel(level)
	// 打开一个日志文件
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("无法打开日志文件: %v", err)
	}
	logrus.SetOutput(io.MultiWriter(file, os.Stdout))
}
