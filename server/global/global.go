package global

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"mall.com/config"
)

var (
	Config config.Config
	Db     *gorm.DB
	Rdb    *redis.Client
)
