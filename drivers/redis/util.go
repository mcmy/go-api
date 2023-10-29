package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func PingRedis(ctx context.Context, client *redis.Client) error {
	_, err := client.Ping(ctx).Result()
	return err
}

func GenKey(token string, key string) string {
	return Prefix + token + ":" + key
}

func GetKey(token string) string {
	return Prefix + token
}

func GenGlobalKey(key string) string {
	return Prefix + key
}

func ExistsKey(ctx context.Context, key ...string) (bool, error) {
	exists, err := Redis.Exists(ctx, key...).Result()
	if err != nil {
		return false, err
	}
	return exists > 0, nil
}
