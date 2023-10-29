package db

import (
	"api/config"
	"api/drivers/db/query"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

var Gorm *gorm.DB

func InitMysql() error {
	dsn := config.T.DB.DataSource
	var err error

	gormLogger := logger.Discard
	if gin.Mode() == gin.DebugMode {
		gormLogger = logger.New(
			Writer{},
			logger.Config{
				SlowThreshold: time.Second,
				LogLevel:      logger.Info,
				Colorful:      false,
			},
		)
	}
	Gorm, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			CreateBatchSize: 1000,
			PrepareStmt:     true,
			Logger:          gormLogger,
			QueryFields:     true,
			//SkipDefaultTransaction: true,
			//NamingStrategy: schema.NamingStrategy{
			//	TablePrefix: conf.Conf.Database.TablePrefix,
			//},
		},
	)
	db, err := Gorm.DB()
	if err != nil {
		return err
	}
	db.SetMaxIdleConns(config.T.DB.MaxIdle)
	db.SetMaxOpenConns(config.T.DB.MaxOpen)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err != nil {
		return err
	}

	query.SetDefault(Gorm)

	return nil
}

type Writer struct {
}

func (w Writer) Printf(format string, args ...interface{}) {
	log.Printf(format, args...)
}
