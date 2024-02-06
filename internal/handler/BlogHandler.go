package handler

import (
	"github.com/baiyaoyu/bpics-v2/internal/biz/avator"
	"github.com/baiyaoyu/bpics-v2/internal/biz/blog"
	"github.com/gin-gonic/gin"
)

type BlogHander struct {
}

// 模板渲染的
func (blogHandler *BlogHander) BlogViewHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		avatorPic := avator.FetchOneByDate()
		blogs := blog.ListBlog()
		ctx.HTML(200, "blog.html", gin.H{
			"who":   avatorPic.Who,
			"img":   avatorPic.Path,
			"blogs": blogs,
		})
	}
}

// 纯json返回
func (blogHandler *BlogHander) BlogJsonHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(200, "")
	}
}
