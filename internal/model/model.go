package model

import (
	"gin-todo/conf"
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var config = conf.Config

type DbContext struct {
	dsn string
	DB  *gorm.DB
}

var Ctx = &DbContext{}

func (ctx *DbContext) Connect() {

	if ctx.dsn == "" {
		log.Panic().Msg("dsn is empty")
		return
	}

	logLevel := logger.Silent
	if config.SqlDebug {
		logLevel = logger.Info
	}

	db, err := gorm.Open(mysql.Open(ctx.dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		log.Panic().Msg("failed to connect database: " + err.Error())
		return
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Panic().Msg("failed to get sqlDB: " + err.Error())
		return
	}

	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxIdleTime(time.Second * 5)
	sqlDB.SetConnMaxLifetime(time.Second * 10)

	ctx.DB = db

}

func (ctx *DbContext) InitWithDSN(dsn string) {
	ctx.dsn = dsn
	ctx.Connect()
}
