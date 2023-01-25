package main

import (
	"flag"
	"fmt"
	"wobbs/config"
	"wobbs/router"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "f", "./config/config.yaml", "配置文件路径")
	flag.Parse()

	// 初始化配置
	config.InitConfig(configPath)
	// 初始化日志
	config.InitLogger(config.Conf.LogConfig)
	// 初始化数据库
	config.InitDB(config.Conf.MySQLConfig)
	// 初始化路由
	r := router.InitRouter()
	r.Run(fmt.Sprintf("%s:%d", config.Conf.Host, config.Conf.Port))
}
