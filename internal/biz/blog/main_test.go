package blog

import (
	"os"
	"testing"

	"github.com/baiyaoyu/bpics-v2/internal/config"
	"github.com/baiyaoyu/bpics-v2/internal/db"
	_ "github.com/lib/pq"
)

func TestMain(m *testing.M) {
	config.InitConfig("../../../config")
	db.InitDB()
	os.Exit(m.Run())
}
