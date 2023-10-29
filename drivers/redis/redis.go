package redis

import (
	"api/config"
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

const (
	Prefix = "go:server:"
)

var Redis *redis.Client

func InitRedis() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if Redis != nil {
		if err := PingRedis(ctx, Redis); err == nil {
			return nil
		}
	}
	Redis = redis.NewClient(&redis.Options{
		Addr:     config.T.Redis.Addr,
		Password: config.T.Redis.Password,
		DB:       0,
		PoolSize: 100,
	})

	if err := PingRedis(ctx, Redis); err != nil {
		return err
	}
	return nil
}
