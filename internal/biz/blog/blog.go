package blog

import (
	"time"

	"github.com/baiyaoyu/bpics-v2/internal/db"
)

type Blog struct {
	Id         int
	Title      string
	Path       string
	Author     string
	Type       string
	Tag        string
	CreateDate time.Time
	ModifyDate time.Time
	Deleted    bool
}

type BlogVo struct {
	Id         int
	Title      string
	Path       string
	Author     string
	Type       string
	Tag        string
	Content    string
	CreateDate string
	Date       string
	ModifyDate string
}

const TbName string = "biz_blog"

// TbName Avator's table name
func (blog *Blog) TableName() string {
	return TbName
}

func (b *Blog) SaveBlog(blog Blog) {
	db.DbPool.Model(&Blog{}).Select("id", "title", "path", "author", "type", "tag", "create_date", "modify_date", "deleted").Create(blog)
}

func (b *Blog) ListBlog() []BlogVo {
	list := make([]Blog, 0)
	vos := make([]BlogVo, 0)
	db.DbPool.Select("id", "title", "path", "author", "type", "tag", "create_date", "modify_date", "deleted").Where("deleted = ?", false).Find(&list)
	for _, b := range list {
		vos = append(vos, b.convertToVo())
	}
	return vos
}

func (b *Blog) GetBlogById(id int) BlogVo {
	var blog Blog
	db.DbPool.Model(&Blog{}).Where("id = ?", id).Select("id", "title", "path", "author", "type", "tag", "create_date", "modify_date", "deleted").Find(&blog)
	return blog.convertToVo()
}

func (blog *Blog) convertToVo() BlogVo {
	vo := BlogVo{
		Id:         blog.Id,
		Path:       blog.Path,
		Title:      blog.Title,
		Tag:        blog.Tag,
		Author:     blog.Author,
		Type:       blog.Type,
		CreateDate: blog.CreateDate.Format("2006.01.02 15:04:05"),
		Date:       blog.CreateDate.Format("2006.01.02"),
		ModifyDate: blog.ModifyDate.Format("2006.01.02 15:04:05")}
	return vo
}
