package common

import (
	"api/config"
	"api/drivers/redis"
	"api/utils"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strings"
	"time"
)

func GetCtxToken(ctx *gin.Context) string {
	token := GetCtxValueString(ctx, "token")
	if token != "" {
		return token
	}

	authToken := ctx.GetHeader("Authorization")
	token = ExtractBearerToken(authToken)

	if token != "" {
		if isValidToken(token) {
			ctx.Set("token", token)
			return token
		}
	}

	return ""
}

func GetOrGenerateCtxToken(ctx *gin.Context) (string, error) {
	token := GetCtxToken(ctx)
	if token != "" {
		return token, nil
	}
	regenerateToken, err := RegenerateToken(ctx)
	return regenerateToken, err
}

func RegenerateToken(ctx *gin.Context) (string, error) {
	token := utils.CalculateSHA1(uuid.New().String())
	ctx.Set("token", token)
	err := redis.Redis.Set(context.Background(),
		redis.GenKey(token, "create"),
		time.Now().Unix(),
		time.Duration(config.T.Http.SessionLiveSecond)*time.Second).Err()
	return token, err
}

func DeleteToken(ctx *gin.Context) {
	token := GetCtxValueString(ctx, "token")
	if token != "" {
		redis.Redis.Del(context.Background(), redis.GetKey(token))
	}
}

func ExtractBearerToken(authToken string) string {
	if strings.HasPrefix(authToken, "Bearer ") {
		return strings.TrimPrefix(authToken, "Bearer ")
	}
	return ""
}

func isValidToken(token string) bool {
	get := redis.Redis.Get(context.Background(), redis.GenKey(token, "create"))
	i, err := get.Int64()
	if err != nil || i <= 0 {
		return false
	}
	if i+int64(config.T.Http.SessionLiveMaxHour*3600) < time.Now().Unix() {
		redis.Redis.Del(context.Background(), redis.GetKey(token))
		return false
	}
	redis.Redis.Expire(context.Background(),
		redis.GetKey(token),
		time.Duration(config.T.Http.SessionLiveSecond)*time.Second)
	return true
}
