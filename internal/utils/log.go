package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/goo-site/log"
	"github.com/goo-site/log/writers"
)

func LogInit() {
	consoleWriter := &writers.ConsoleWriter{}

	fileWriter := &writers.FileWriter{}
	fileWriter.SetLogDir("/home/william/go/src/web-server/log")
	fileWriter.SetLogSize(1024 * 1024)

	log.SetLogLevel("Info")
	log.AddWriter(consoleWriter, fileWriter)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Info("| request | %15s | %s %s |", c.ClientIP(), c.Request.Method, c.Request.URL.Path)
		c.Next()
		log.Info("| response | %3d |", c.Writer.Status())
	}
}
