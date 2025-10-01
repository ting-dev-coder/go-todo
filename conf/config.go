package conf

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

type config struct {
	ENV      string `default:"DEV"`
	Port     string `default:"8085"`
	Dsn      string `mapstructure:"dsn"`
	SqlDebug bool   `mapstructure:"sql_debug"`
}

var Config = &config{}

func init() {
	once := sync.Once{}
	once.Do(func() {
		// 載入環境變量文件
		godotenv.Load(".env.local", ".env")

		// 初始化config預設值
		Config.ENV = getEnv("ENV", "DEV")
		Config.Port = getEnv("PORT", "8085")
		Config.Dsn = getEnv("DSN", "")
		Config.SqlDebug = getEnvBool("SQL_DEBUG", false)
		fmt.Printf(Config.Dsn)
	})
}

// getEnv 獲取環境變量，如果不存在則返回預設值
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvBool 獲取布林型環境變量
func getEnvBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if result, err := strconv.ParseBool(value); err == nil {
			return result
		}
	}
	return defaultValue
}
