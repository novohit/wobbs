package test

import (
	"flag"
	"testing"
	"wobbs-server/config"
	"wobbs-server/logic"
)

func TestVoting(t *testing.T) {
	var configPath string
	flag.StringVar(&configPath, "f", "../config/config.yaml", "配置文件路径")
	flag.Parse()

	// 初始化配置
	config.InitConfig(configPath)
	// 初始化日志
	config.InitLogger(config.Conf.LogConfig)
	// 初始化缓存
	config.InitRedis(config.Conf.RedisConfig)
	//logic.PostVoting(778797973, dto.VoteDTO{PostID: 4, Type: 1})
	//logic.PostVoting(4, dto.VoteDTO{PostID: 6, Type: 1})
	//logic.PostVoting(5, dto.VoteDTO{PostID: 5, Type: -1})
	//logic.PostVoting(10386626773520384, dto.VoteDTO{PostID: 9, Type: 0})

	logic.GetVotes([]string{"9", "7", "5", "6"})
}
