package handler

import (
	"io"
	"net/http"
	"strconv"

	"github.com/baiyaoyu/bpics-v2/internal/biz/avator"
	"github.com/baiyaoyu/bpics-v2/internal/biz/blog"
	"github.com/baiyaoyu/bpics-v2/internal/config"
	"github.com/gin-gonic/gin"
)

type BlogHander struct {
	rootPath http.Dir
}

func NewBlogHandler() *BlogHander {
	path := http.Dir(config.MdPath)
	return &BlogHander{
		rootPath: path,
	}
}

// 模板渲染的
func (handler *BlogHander) BlogViewHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var avator avator.Avator
		var b blog.Blog
		avator = avator.FetchOneByDate()
		blogs := b.ListBlog()
		vo := &blog.BlogVo{
			Path: blogs[0].Path,
		}
		ctx.HTML(200, "blog.html", gin.H{
			"who":     avator.Who,
			"img":     avator.Path,
			"blogs":   blogs,
			"content": handler.readContent(vo),
		})
	}
}

// 纯json返回
func (handler *BlogHander) BlogJsonHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		idstr := ctx.Param("id")
		id, err := strconv.Atoi(idstr)
		if err != nil {
			ctx.JSON(500, nil)
		}
		var blog blog.Blog
		contentVo := blog.GetBlogById(id)
		handler.readContent(&contentVo)
		ctx.JSON(200, contentVo)
	}
}

// 读取对应blog的内容
func (handler *BlogHander) readContent(blog *blog.BlogVo) string {
	file, _ := handler.rootPath.Open(blog.Path)
	defer file.Close()
	bs, _ := io.ReadAll(file)
	blog.Content = string(bs)
	return string(bs)
}
