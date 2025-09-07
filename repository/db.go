package repository

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/joho/godotenv"
)

var DB *gorm.DB

func Init() {
	// 可選：讀取 .env
	_ = godotenv.Load()

	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		// 預設：XAMPP root 無密碼
		// 如果你的 root 有密碼，把 root:你的密碼@... 換掉
		dsn = "root:tp2793371@tcp(127.0.0.1:3306)/todolist_db?parseTime=true&loc=Local&charset=utf8mb4"
	}

	gcfg := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn), // 想看更詳細的 SQL 可改成 logger.Info
	}

	db, err := gorm.Open(mysql.Open(dsn), gcfg)
	if err != nil {
		log.Fatalf("open db failed: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("get sqlDB failed: %v", err)
	}

	// 連線池設定（可依需要調整）
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	// 自動建表
	// if err := db.AutoMigrate(&model.Task{}); err != nil {
	// 	log.Fatalf("auto-migrate failed: %v", err)
	// }

	DB = db
}
