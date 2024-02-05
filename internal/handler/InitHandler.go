package handler

import (
	"net/http"
)

func InitHandler(DataPath string) {
	RootPath = http.Dir(DataPath)
	RootSystem = http.Dir(RootPath)
	FsHandler = http.Handler(http.StripPrefix("/", http.FileServer(RootSystem)))
}
