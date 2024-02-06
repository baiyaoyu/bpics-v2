package blog

import (
	"fmt"
	"testing"
	"time"
)

func TestSaveBlog(t *testing.T) {
	blog := Blog{
		Id:         10002,
		Title:      "关于我",
		Path:       "10001.md",
		Author:     "HiddenCat",
		Type:       "about",
		Tag:        "about",
		CreateDate: time.Now(),
		ModifyDate: time.Now(),
		Deleted:    false,
	}
	blog.SaveBlog(blog)
}

func TestFetchById(t *testing.T) {
	var blog BlogVo
	var b Blog
	blog = b.GetBlogById(10001)
	fmt.Println(blog)
}

func TestListBlog(t *testing.T) {
	var b Blog
	blogs := b.ListBlog()
	fmt.Println(blogs)
}
