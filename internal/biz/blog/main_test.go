package blog

import (
	"os"
	"testing"

	"github.com/baiyaoyu/bpics-v2/internal/config"
	_ "github.com/lib/pq"
)

func TestMain(m *testing.M) {
	config.InitConfig("../../../config")
	config.InitOther()
	os.Exit(m.Run())
}
