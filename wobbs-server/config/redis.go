package config

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

var ctx = context.Background()

var RDB *redis.Client

func InitRedis(cfg *RedisConfig) {
	host := cfg.Host
	port := cfg.Port
	password := cfg.Password
	database := cfg.Database
	RDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: password, // no password set
		DB:       database, // use default DB
	})

	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	zap.L().Info("redis init success")
	return
	//
	//err := rdb.Set(ctx, "key", "value", 0).Err()
	//if err != nil {
	//	panic(err)
	//}
	//
	//val, err := rdb.Get(ctx, "key").Result()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("key", val)
	//
	//val2, err := rdb.Get(ctx, "key2").Result()
	//if err == redis.Nil {
	//	fmt.Println("key2 does not exist")
	//} else if err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println("key2", val2)
	//}
	// Output: key value
	// key2 does not exist
}
