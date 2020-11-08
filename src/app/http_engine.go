package app

import (
	"github.com/gin-gonic/gin"
	"github.com/mandatorySuicide/golang-code-quality/src/controller"
)

func SetupServer() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", controller.Ping)
	return r
}
