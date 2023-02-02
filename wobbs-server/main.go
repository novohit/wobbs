package main

import (
	"flag"
	"fmt"

	"wobbs-server/config"
	"wobbs-server/pkg/snowflake"
	"wobbs-server/router"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8081
// @BasePath /api
func main() {
	var configPath string
	flag.StringVar(&configPath, "f", "./config/config.yaml", "配置文件路径")
	flag.Parse()

	// 初始化配置
	config.InitConfig(configPath)
	// 初始化日志
	config.InitLogger(config.Conf.LogConfig)
	// 初始化雪花算法结点
	snowflake.Init(config.Conf.SnowflakeConfig.StartTime, config.Conf.SnowflakeConfig.MachineID)
	// 初始化数据库
	config.InitDB(config.Conf.MySQLConfig)
	// 初始化缓存
	config.InitRedis(config.Conf.RedisConfig)
	// 初始化路由
	r := router.InitRouter()
	r.Run(fmt.Sprintf("%s:%d", config.Conf.Host, config.Conf.Port))
}
