package main

import (
	"fmt"
	"gin-todo/conf"
	"gin-todo/internal/model"
	zlog "gin-todo/internal/pkg/logger"
	"gin-todo/internal/router"
)

var (
	config = conf.Config
	logger = zlog.NewBaseLogger(config.ENV)
)

func main() {

	if err := run(); err != nil {
		logger.Error().Msg(err.Error())
	}

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"message": "pong"})
	// })

	// r.Run(":8080") // http://localhost:8080

}

func run() error {
	model.Ctx.InitWithDSN(config.Dsn)

	api, err := router.SetupRouter()
	if err != nil {
		return fmt.Errorf("error happened in setup router: %w", err)
	}

	api.Run(":8085")
	return nil
}
