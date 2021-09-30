package initialize

import (
	"github.com/go-redis/redis/v8"
	"mall.com/global"
)

func RedisClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	global.RDb = rdb
}