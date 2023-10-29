package common

import (
	"api/drivers/redis"
	"api/dto/resp"
	"api/i18n"
	"context"
	"github.com/gin-gonic/gin"
)

type ApiHandleFunc func(*ApiContext)

type ApiContext struct {
	*gin.Context
}

func (c *ApiContext) GetToken() (string, error) {
	return GetOrGenerateCtxToken(c.Context)
}

func (c *ApiContext) HasToken() bool {
	return len(GetCtxToken(c.Context)) > 0
}

func (c *ApiContext) GetCtxValueString(key string, def ...string) string {
	return GetCtxValueString(c.Context, key, def...)
}

func (c *ApiContext) GetAcceptLang() string {
	return GetAcceptLang(c.Context)
}

// JSONI msg国际化,返回resp.T
func (c *ApiContext) JSONI(code int, obj *resp.M) {
	msg := obj.GetMsg()
	if msg != "" {
		obj.Msg(i18n.Use(c.GetAcceptLang()).Get(msg))
	}
	c.Context.JSON(code, obj)
}

func (c *ApiContext) VerifyCode() (bool, error) {
	code := c.Param("code")
	if code == "" {
		return false, nil
	}
	token, err := c.GetToken()
	if err != nil {
		return false, err
	}
	result, err := redis.Redis.Get(context.Background(), redis.GenKey(token, "verify_code")).Result()
	if err != nil {
		return false, err
	}
	redis.Redis.Del(context.Background(), redis.GenKey(token, "verify_code"))
	if result == "" || result != code {
		return false, nil
	}
	return true, nil
}
