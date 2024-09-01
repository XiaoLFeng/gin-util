package test

import (
	"github.com/XiaoLFeng/go-gin-util/bconfig"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestGinLogger(t *testing.T) {
	r := gin.Default()
	bconfig.LogConfiguration("logger.log", true, logrus.TraceLevel)
	r.GET("/test", func(c *gin.Context) {
		logrus.Info("test")
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
