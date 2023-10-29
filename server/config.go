package server

import (
	"api/config"
	"github.com/gin-gonic/gin"
	"log"
)

func InitGinAPIServer(engine *gin.Engine) {
	err := engine.SetTrustedProxies(config.T.Http.TrustProxyIpCidr)
	if err != nil {
		log.Fatalln("trusted proxies set error", err)
	}
}
