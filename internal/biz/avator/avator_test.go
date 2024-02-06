package avator

import (
	"fmt"
	"testing"
	"time"
)

func TestAvator(t *testing.T) {
	avator := Avator{
		Id:          2,
		Path:        "http://www.hiddencat.fun:8000/fs/file/avator/tiger.jpg",
		Who:         "跳跳虎",
		Create_time: time.Now(),
	}
	avator.SaveToDb(avator)
}

func TestFetchOneAvator(t *testing.T) {
	var avator Avator
	avator = avator.FetchOneByDate()
	fmt.Println(avator)
}

func TestGetOneById(t *testing.T) {
	var avator Avator
	avator = avator.GetOneById(1)
	fmt.Println(avator)
}

func TestListAvator(t *testing.T) {
	var avator Avator
	avators := avator.ListAll()
	fmt.Println(avators)
}
