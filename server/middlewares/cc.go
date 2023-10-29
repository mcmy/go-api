package middlewares

import (
	"api/drivers/redis"
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

func CCInterceptor(c *gin.Context) {
	ip := c.ClientIP()
	ipKey := redis.GenGlobalKey("cc:" + ip)
	exists, err := redis.ExistsKey(context.Background(), ipKey)
	if err != nil {
		c.AbortWithStatus(403)
		return
	}
	if !exists {
		redis.Redis.Set(context.Background(), ipKey, 1, 10*time.Second)
		c.Next()
		return
	}
	result, err := redis.Redis.Incr(context.Background(), ipKey).Result()
	if err != nil || result > 10 {
		c.AbortWithStatus(403)
		return
	}
}
