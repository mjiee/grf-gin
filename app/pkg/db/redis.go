package db

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/mjiee/scaffold-gin/app/pkg/conf"
)

func NewRedis(cfg *conf.Config) (*redis.Client, func()) {
	redis := redis.NewClient(&redis.Options{
		Addr:         cfg.Redis.Addr,
		Username:     cfg.Redis.UserName,
		Password:     cfg.Redis.Password,
		DB:           cfg.Redis.DB,
		PoolSize:     cfg.Redis.PoolSize,
		MinIdleConns: cfg.Redis.MinIdleConns,
	})

	cleanup := func() {
		if err := redis.Close(); err != nil {
			fmt.Println("close redis connection error")
		}
	}

	return redis, cleanup
}
