package test

import (
	"flag"
	"testing"

	"wobbs-server/config"
	"wobbs-server/model"
)

func TestM(t *testing.T) {
	var configPath string
	flag.StringVar(&configPath, "f", "./config/config.yaml", "配置文件路径")
	flag.Parse()
	// 初始化配置
	config.InitConfig(configPath)
	// 初始化日志
	config.InitLogger(config.Conf.LogConfig)
	// 初始化数据库
	config.InitDB(config.Conf.MySQLConfig)
	config.GetDB().Create(&model.User{UserID: 24234324, Username: "novo44", Password: "helloworld"})
	config.GetDB().Delete(&model.User{BaseModel: model.BaseModel{ID: 1}})
}
