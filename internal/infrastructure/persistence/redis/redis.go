package redis

import (
	"fmt"

	"github.com/boldd/internal/config"
	goredis "github.com/redis/go-redis/v9"
)

func NewRedisClient(cfg *config.RedisConfig) *goredis.Client {
	return goredis.NewClient(&goredis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})
}
