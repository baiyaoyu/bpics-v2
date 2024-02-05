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
	ModifyDate string
}

func SaveBlog(blog Blog) {
	tx, _ := db.DbPool.Begin()
	defer tx.Commit()
	blog.CreateDate = time.Now()
	tx.Exec("insert into biz_blog(id,title,path,author,type,tag,create_date,modify_date,deleted) values(?,?,?,?,?,?,?,?,?)", blog.Id, blog.Title, blog.Path, blog.Author, blog.Type, blog.Tag, time.Now(), time.Now(), false)
}

func ListBlog() []Blog {
	list := make([]Blog, 0)
	rows, _ := db.DbPool.Query("select id,title,path,author,type,tag,create_date,modify_date from biz_blog where deleted = false")
	defer rows.Close()
	for rows.Next() {
		var item Blog
		rows.Scan(&item.Id, &item.Title, &item.Path, &item.Author, &item.Type, &item.Tag, &item.CreateDate, &item.ModifyDate)
		list = append(list, item)
	}
	return list
}

func GetBlogById(id int) BlogVo {
	rows, _ := db.DbPool.Query("select id,title,path,author,type,tag,create_date,modify_date from biz_blog where id = ? and deleted = false", id)
	defer rows.Close()
	rows.Next()
	var res BlogVo
	var createDate time.Time
	var modifyDate time.Time
	rows.Scan(&res.Id, &res.Title, &res.Path, &res.Author, &res.Type, &res.Tag, &createDate, &modifyDate)
	res.CreateDate = createDate.Format("2006-03-04")
	res.ModifyDate = modifyDate.Format("2006-03-04")
	return res
}
