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

func SaveToDb(avator Avator) {
	tx, _ := db.DbPool.Begin()
	defer tx.Commit()
	avator.Create_time = time.Now()
	tx.Exec("insert into biz_title_pic(id,path,who,create_time) values(?,?,?,?)", avator.Id, avator.Path, avator.Who, avator.Create_time)
}

func FetchOneByDate() Avator {
	days := time.Now().Day()
	var count int
	tx, _ := db.DbPool.Begin()
	defer tx.Commit()
	rows, _ := tx.Query("select count(*) from biz_title_pic")
	defer rows.Close()
	rows.Next()
	rows.Scan(&count)
	rows.Close()
	id := days % count
	return GetOneById(id)
}

func GetOneById(id int) Avator {
	var res Avator
	rows, _ := db.DbPool.Query("select * from biz_title_pic where id = ?", id)
	defer rows.Close()
	rows.Next()
	rows.Scan(&res.Id, &res.Path, &res.Who, &res.Create_time)
	return res
}

func ListAll() []Avator {
	avators := make([]Avator, 0)
	rows, _ := db.DbPool.Query("select * from biz_title_pic")
	defer rows.Close()
	for rows.Next() {
		var item Avator
		rows.Scan(&item.Id, &item.Path, &item.Who, &item.Create_time)
		avators = append(avators, item)
	}
	return avators
}
