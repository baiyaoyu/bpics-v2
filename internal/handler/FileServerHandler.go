package handler

import (
	"net/http"
	"net/url"
	"time"

	"github.com/baiyaoyu/bpics-v2/internal/logger"
	"github.com/gin-gonic/gin"
)

type Item struct {
	Name       string `json:"name"`       // 目录或文件的名称
	IsDir      string `json:"type"`       // 是否为目录
	ModifyTime string `json:"modifyTime"` // 修改的时间
	Size       int64  `json:"size"`
	Href       string `json:"href"`
}

var RootPath http.Dir
var RootSystem http.FileSystem
var FsHandler http.Handler

func FileSystemHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		path := ctx.Request.RequestURI
		// 转义后可支持中文路径的文件
		filePath, _ := url.QueryUnescape(path)
		// fmt.Println(path)
		if JudgePath(filePath) {
			files := ListDir(filePath)
			// avatorPic := ava.GetOnePicByDate()
			ctx.HTML(http.StatusOK, "filedir.html", gin.H{
				"who":      "潜伏的喵星人",
				"img":      "843053.jpg",
				"pipeline": files,
				"path":     filePath,
			})
		} else {
			FsHandler.ServeHTTP(ctx.Writer, ctx.Request)
		}
	}
}

func JudgePath(path string) bool {
	file, _ := RootPath.Open(path)
	info, _ := file.Stat()

	return info.IsDir()
}

func ListDir(path string) []Item {
	file, _ := RootPath.Open(path)
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

func InnerServe(w http.ResponseWriter, r *http.Request, name string) {
	file, _ := RootPath.Open(name)
	http.ServeContent(w, r, name, time.Now(), file)
}
