package redis

import (
	"github.com/boldd/internal/config"
	goredis "github.com/redis/go-redis/v9"
)

func NewRedisClient(cfg *config.RedisConfig) *goredis.Client {
	return goredis.NewClient(&goredis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
		DB:       cfg.DB,
		Protocol: cfg.Protocol,
	})
}
