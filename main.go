package main

import (
	"api/config"
	"api/drivers/db"
	"api/drivers/redis"
	"api/i18n"
	"api/server"
	"api/server/middlewares"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	fmt.Println("    _   ________   ________  ______   _____ __________ _    ____________ \n   / | / / ____/  /  _/ __ \\/_  __/  / ___// ____/ __ \\ |  / / ____/ __ \\\n  /  |/ / /_      / // / / / / /     \\__ \\/ __/ / /_/ / | / / __/ / /_/ /\n / /|  / __/    _/ // /_/ / / /     ___/ / /___/ _, _/| |/ / /___/ _, _/ \n/_/ |_/_/      /___/\\____/ /_/     /____/_____/_/ |_| |___/_____/_/ |_|")
	if err := i18n.InitI18n(); err != nil {
		log.Fatalln("i18n init error:", err)
	}
	log.Println("default lang set:", i18n.DefaultLang)
	if err := config.InitConfig(); err != nil {
		log.Fatalln("config init error:", err)
	}
	if err := db.InitMysql(); err != nil {
		log.Fatalln("mysql connect error:", err)
	}
	if err := redis.InitRedis(); err != nil {
		log.Fatalln("redis connect error:", err)
	}
}

func main() {
	engine := gin.New()
	middlewares.RegisterInitMiddleware(engine)
	server.InitGinAPIServer(engine)
	server.RegisterRouter(engine)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.T.Http.Port),
		Handler: engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Println("server run on port", srv.Addr)
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
