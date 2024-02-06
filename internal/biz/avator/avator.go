package avator

import (
	"time"

	"github.com/baiyaoyu/bpics-v2/internal/db"
)

type Avator struct {
	Id          int
	Path        string
	Who         string
	Create_time time.Time
}

const TbName string = "biz_title_pic"

// TbName Avator's table name
func (avator *Avator) TableName() string {
	return TbName
}

func (a *Avator) SaveToDb(avator Avator) {
	db.DbPool.Model(&Avator{}).Select("id", "path", "who", "create_time").Create(avator)
}

func (a *Avator) FetchOneByDate() Avator {
	days := time.Now().Day()
	var count int64
	db.DbPool.Model(&Avator{}).Count(&count)
	id := days % int(count)
	return a.GetOneById(id)
}

func (a *Avator) GetOneById(id int) Avator {
	var res Avator
	db.DbPool.Select("id", "path", "who", "create_time").Where("id=?", id).Find(&res)
	return res
}

func (a *Avator) ListAll() []Avator {
	avators := make([]Avator, 0)
	db.DbPool.Select("id", "path", "who", "create_time").Find(&avators)
	return avators
}
