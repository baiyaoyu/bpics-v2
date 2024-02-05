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
	SaveBlog(blog)
}

func TestFetchById(t *testing.T) {
	blog := GetBlogById(10001)
	fmt.Println(blog)
}

func TestListBlog(t *testing.T) {
	blogs := ListBlog()
	fmt.Println(blogs)
}
