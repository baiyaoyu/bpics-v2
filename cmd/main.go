package main

import (
	"github.com/baiyaoyu/bpics-v2/internal/config"
	"github.com/baiyaoyu/bpics-v2/internal/logger"
	"github.com/baiyaoyu/bpics-v2/internal/server"
)

func main() {
	config.InitConfig("./config")
	config.InitOther()
	server.InitEngine(config.TemplPath)
	Addr := config.GetServerAddr()
	logger.Debug(nil, "web即将启动")
	server.Default.Run(Addr)
}
