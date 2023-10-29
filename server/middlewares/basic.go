package middlewares

import (
	"api/dto/resp"
	"github.com/gin-gonic/gin"
	"log"
	"runtime/debug"
	"time"
)

func RegisterInitMiddleware(r *gin.Engine) {
	r.Use(liteLogger)
	r.Use(Recover)
}

func liteLogger(c *gin.Context) {
	start := time.Now()
	path := c.Request.URL.Path
	raw := c.Request.URL.RawQuery

	c.Next()

	if raw != "" {
		path = path + "?" + raw
	}

	log.Println(time.Now().Sub(start), c.Writer.Status(), c.ClientIP(), path)
}

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			if gin.Mode() == gin.DebugMode {
				debug.PrintStack()
			}
			c.JSON(500, resp.CodeMsg(500, "server error"))
			c.Abort()
		}
	}()
	c.Next()
}

// recover错误，转string
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
