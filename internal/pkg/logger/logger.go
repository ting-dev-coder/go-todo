package pkg

import (
	"io"
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
)

func NewBaseLogger(env string) zerolog.Logger {
	// 設定時間欄位格式 (ISO 8601)
	zerolog.TimeFieldFormat = time.RFC3339

	// 設定全域的 log level (最低是 DebugLevel)
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	// log 輸出位置：這裡只設定到 stdout
	logWriters := io.MultiWriter(os.Stdout)

	var logger zerolog.Logger

	// 保證只初始化一次 (雖然這裡每次呼叫函數都會重新跑，其實用處不大)
	once := sync.Once{}
	once.Do(func() {
		logger = zerolog.New(logWriters).With().
			Timestamp().     // 每筆 log 加上時間
			Caller().        // 每筆 log 加上呼叫來源（檔案 + 行號）
			Str("ENV", env). // 加上固定欄位 ENV，方便區分環境 (DEV/PROD)
			Logger()
	})

	return logger
}
