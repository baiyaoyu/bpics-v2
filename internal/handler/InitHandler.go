package handler

import (
	"net/http"

	"github.com/baiyaoyu/bpics-v2/internal/config"
)

func InitHandler() {
	RootPath = http.Dir(config.DataPath)
	RootSystem = http.Dir(RootPath)
	FsHandler = http.Handler(http.StripPrefix("/", http.FileServer(RootSystem)))
}
