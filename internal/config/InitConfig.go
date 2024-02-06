package config

import (
	"fmt"

	"github.com/baiyaoyu/bpics-v2/internal/db"
	"github.com/baiyaoyu/bpics-v2/internal/logger"
	"github.com/spf13/viper"
)

var Logpath string
var DSN string
var DataPath string
var Port int
var Ip string
var TemplPath string
var MdPath string

func InitConfig(Path string) {
	viper.SetConfigName("configs")
	// 设置配置文件的类型
	viper.SetConfigType("yaml")
	// 添加配置文件的路径，指定 config 目录下寻找
	viper.AddConfigPath(Path)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	Logpath = viper.GetString("log.path")
	DSN = viper.GetString("middleware.mysql")
	DataPath = viper.GetString("data.path")
	Port = viper.GetInt("server.port")
	Ip = viper.GetString("server.ip")
	TemplPath = viper.GetString("templ")
	MdPath = viper.GetString("md.path")
}

func InitOther() {
	db.InitDB(DSN)
	logger.InitLog(Logpath)
}

func GetServerAddr() string {
	addr := fmt.Sprintf("%s:%d", Ip, Port)
	return addr
}
