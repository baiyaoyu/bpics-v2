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

func (b *Blog) ListBlog() []Blog {
	list := make([]Blog, 0)
	db.DbPool.Select("id", "title", "path", "author", "type", "tag", "create_date", "modify_date", "deleted").Where("deleted = ?", false).Find(&list)
	return list
}

func (b *Blog) GetBlogById(id int) BlogVo {
	var blog Blog
	db.DbPool.Model(&Blog{}).Select("id", "title", "path", "author", "type", "tag", "create_date", "modify_date", "deleted").Find(&blog)
	vo := BlogVo{
		Id:         blog.Id,
		Path:       blog.Path,
		Title:      blog.Title,
		Tag:        blog.Tag,
		Author:     blog.Author,
		Type:       blog.Type,
		Date:       blog.CreateDate.Format("2006-01-03 15:00:01"),
		ModifyDate: blog.ModifyDate.Format("2006-01-03 15:00:01")}
	return vo
}
