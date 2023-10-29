package test

import (
	"api/config"
	"api/dto/resp"
	"api/i18n"
	"fmt"
	"log"
	"testing"
	"time"
)

func init() {
	err := config.InitConfig("../conf/config.toml")
	if err != nil {
		log.Fatalln("config error")
	}
	err = i18n.InitI18n("../conf/lang.toml")
	if err != nil {
		log.Fatalln("i18n error")
	}
}

func TestMsg(t *testing.T) {
	msg := resp.Code(200).Msg("ok")
	log.Println(msg)
}

func TestSpeed(t *testing.T) {
	var a string
	_, _ = fmt.Scanln(&a)
	log.Println(a)
	start := time.Now().UnixNano()
	for i := 0; i < 100000; i++ {
		if len(a) > 0 {

		}
	}
	log.Println(time.Now().UnixNano() - start)
	start = time.Now().UnixNano()
	for i := 0; i < 100000; i++ {
		if a == "" {

		}
	}
	log.Println(time.Now().UnixNano() - start)
}
