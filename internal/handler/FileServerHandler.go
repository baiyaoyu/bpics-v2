package handler

import (
	"net/http"
	"net/url"
	"time"

	"github.com/baiyaoyu/bpics-v2/internal/biz/avator"
	"github.com/baiyaoyu/bpics-v2/internal/config"
	"github.com/baiyaoyu/bpics-v2/internal/logger"
	"github.com/gin-gonic/gin"
)

type FsHandler struct {
	fsHandler http.Handler
	rootPath  http.Dir
}

type Item struct {
	Name       string `json:"name"`       // 目录或文件的名称
	IsDir      string `json:"type"`       // 是否为目录
	ModifyTime string `json:"modifyTime"` // 修改的时间
	Size       int64  `json:"size"`
	Href       string `json:"href"`
}

type FileReq struct {
	Path string `json:"path"`
	Op   string `json:"op"` // list download view upload
}

func NewFsHandler() FsHandler {
	path := http.Dir(config.DataPath)
	return FsHandler{
		rootPath:  path,
		fsHandler: http.Handler(http.StripPrefix("/", http.FileServer(path))),
	}
}

// 模板处理使用
func (handler *FsHandler) FileSystemHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		path := ctx.Request.RequestURI
		// 转义后可支持中文路径的文件
		filePath, _ := url.QueryUnescape(path)
		// fmt.Println(path)
		if handler.judgePath(filePath) {
			files := handler.listDir(filePath)
			var avatorPic avator.Avator
			avatorPic = avatorPic.FetchOneByDate()
			ctx.HTML(http.StatusOK, "index.html", gin.H{
				"who":      avatorPic.Who,
				"img":      avatorPic.Path,
				"pipeline": files,
				"path":     filePath,
			})
		} else {
			handler.fsHandler.ServeHTTP(ctx.Writer, ctx.Request)
		}
	}
}

// 前后端分离架构使用
func (handler *FsHandler) FileJsonHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var req FileReq
		ctx.ShouldBindJSON(&req)
		path := req.Path
		// op := req.Op
		// filePath, _ := url.QueryUnescape(req.Path)
		if handler.judgePath(path) {
			files := handler.listDir(req.Path)
			ctx.JSON(200, files)
		} else {
			handler.innerServe(ctx.Writer, ctx.Request, path)
		}
	}
}

func (handler *FsHandler) judgePath(path string) bool {
	file, _ := handler.rootPath.Open(path)
	info, _ := file.Stat()
	return info.IsDir()
}

func (handler *FsHandler) listDir(path string) []Item {
	file, _ := handler.rootPath.Open(path)
	fss, err := file.Readdir(-1)
	if err != nil {
		logger.Error(nil, err)
	}
	items := make([]Item, len(fss)+1)
	items[0] = Item{
		"返回上层",
		"上级目录",
		"",
		0,
		"../",
	}
	for i := 1; i < len(fss)+1; i++ {
		if fss[i-1].IsDir() {
			items[i].IsDir = "目录"
			items[i].Href = fss[i-1].Name() + "/"
		} else {
			items[i].IsDir = "文件"
			items[i].Href = fss[i-1].Name()
		}
		items[i].Name = fss[i-1].Name()
		items[i].ModifyTime = fss[i-1].ModTime().Format("2006-01-02 15:04:05")
		items[i].Size = fss[i-1].Size()
	}
	return items
}

func (handler *FsHandler) innerServe(w http.ResponseWriter, r *http.Request, name string) {
	file, _ := handler.rootPath.Open(name)
	http.ServeContent(w, r, name, time.Now(), file)
}
