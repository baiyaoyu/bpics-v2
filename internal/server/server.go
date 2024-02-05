package server

import (
	"github.com/baiyaoyu/bpics-v2/internal/handler"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var Default *gin.Engine

func InitEngine(TemplPath string) {
	gin.DefaultWriter = logrus.StandardLogger().Out
	r := gin.Default()
	r.LoadHTMLGlob(TemplPath)
	fsHandler := handler.NewFsHandler()
	r.GET("/*name", fsHandler.FileSystemHandler())
	Default = r
}
