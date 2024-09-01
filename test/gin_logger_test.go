package test

import (
	"github.com/XiaoLFeng/go-gin-util/bconfig"
	"github.com/XiaoLFeng/go-gin-util/blog"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"testing"
)

// TestGinLogger
//
// # 测试 GIN 日志
//
// 测试 GIN 日志功能, 通过访问 http://localhost:8080/test 进行测试
func TestGinLogger(t *testing.T) {
	r := gin.Default()
	bconfig.LogConfiguration("../.logs", "logger", true, logrus.TraceLevel)
	r.GET("/test", func(c *gin.Context) {
		blog.Trace("Trace")
		blog.Tracef("Tracef")
		blog.Debug("Debug")
		blog.Debugf("Debugf")
		blog.Info("Info")
		blog.Infof("Infof")
		blog.Warn("Warn")
		blog.Warnf("Warnf")
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})
	r.GET("/error", func(c *gin.Context) {
		blog.Error("Error")
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})
	r.GET("/errorf", func(c *gin.Context) {
		blog.Errorf("Errorf")
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})
	r.GET("/panic", func(c *gin.Context) {
		blog.Panic("Panic")
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})
	r.GET("/panicf", func(c *gin.Context) {
		blog.Panicf("Panicf")
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
