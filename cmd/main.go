package main

import (
	"github.com/baiyaoyu/bpics-v2/internal/config"
	"github.com/baiyaoyu/bpics-v2/internal/db"
	"github.com/baiyaoyu/bpics-v2/internal/handler"
	"github.com/baiyaoyu/bpics-v2/internal/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.InitConfig("./config")
	db.InitDB()
	logger.InitLog()
	handler.InitHandler()
	r.LoadHTMLGlob(config.TemplPath)
	r.GET("/*name", handler.FileSystemHandler())
	r.Run(config.GetServerAddr())
	logger.Debug(nil, "web服务已启动")
}
