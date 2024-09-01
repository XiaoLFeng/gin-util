package bconfig

import (
	"github.com/XiaoLFeng/go-gin-util/blog"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strconv"
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
func LogConfiguration(path string, fileName string, enableColor bool, level logrus.Level) {
	logrus.SetFormatter(new(blog.CustomFormatter))
	logrus.SetLevel(level)

	// GIN 系统设置日志输出
	err := os.Mkdir(path, 0750)
	if err != nil {
		if os.IsExist(err) {
			blog.Trace("日志文件夹已存在")
		} else {
			blog.Fatalf("无法创建日志文件夹: %v", err)
		}
	}
	if !enableColor {
		gin.DisableConsoleColor()
	}

	var getPath string
	if path[len(path)-1:] == "/" || fileName[0:] == "/" {
		getPath = path + fileName + "-0.log"
	} else {
		getPath = path + "/" + fileName + "-0.log"
	}
	// 判断文件是否存在
	for i := 0; ; i++ {
		if _, err := os.Stat(getPath); os.IsNotExist(err) {
			iPath := i
			if i != 0 {
				iPath = i - 1
			}
			blog.Info("本次采用日志文件 " + fileName + "-" + strconv.Itoa(iPath) + ".log")
			blog.Tracef("日志文件路径: %s", getPath)
			_, _ = os.Create(getPath)
			break
		} else {
			if path[len(path)-1:] == "/" || fileName[0:] == "/" {
				getPath = path + fileName + "-" + strconv.Itoa(i) + ".log"
			} else {
				getPath = path + "/" + fileName + "-" + strconv.Itoa(i) + ".log"
			}
		}
	}
	// 打开一个日志文件
	file, err := os.OpenFile(getPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		blog.Fatalf("无法打开日志文件: %v", err)
	}
	logrus.SetOutput(io.MultiWriter(file, os.Stdout))
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
}
