package app

import (
	"github.com/gin-gonic/gin"
	"github.com/mandatorySuicide/golang-code-quality/src/controller"
)

func SetupServer() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("src/templates/*")
	r.GET("/ping", controller.Ping)
	r.GET("/", controller.Index)
	return r
}
