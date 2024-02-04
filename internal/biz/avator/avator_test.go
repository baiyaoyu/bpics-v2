package avator

import (
	"fmt"
	"testing"
	"time"
)

func TestAvator(t *testing.T) {
	avator := Avator{
		Id:          1,
		Path:        "http://www.hiddencat.fun:8000/fs/file/avator/rick.jpeg",
		Who:         "rick",
		Create_time: time.Now(),
	}
	SaveToDb(avator)
}

func TestFetchOneAvator(t *testing.T) {
	avator := FetchOneByDate()
	fmt.Println(avator)
}

func TestGetOneById(t *testing.T) {
	avator := GetOneById(1)
	fmt.Println(avator)
}

func TestListAvator(t *testing.T) {
	avators := ListAll()
	fmt.Println(avators)
}
