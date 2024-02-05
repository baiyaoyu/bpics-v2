package server

import (
	"github.com/baiyaoyu/bpics-v2/internal/handler"
	"github.com/gin-gonic/gin"
)

var Default *gin.Engine

func InitEngine(TemplPath string) {
	r := gin.Default()
	r.LoadHTMLGlob(TemplPath)
	r.GET("/*name", handler.FileSystemHandler())
	Default = r
}
