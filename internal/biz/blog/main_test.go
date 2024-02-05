package blog

import (
	"os"
	"testing"

	"github.com/baiyaoyu/bpics-v2/internal/db"
	_ "github.com/lib/pq"
)

func TestMain(m *testing.M) {
	db.InitDB("bpics:bpics@tcp(192.168.1.88:3306)/bpics-v2?charset=utf8mb4&parseTime=True&loc=Local")
	os.Exit(m.Run())
}
