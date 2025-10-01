package router

import (
	"gin-todo/internal/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() (*gin.Engine, error) {
	r := gin.Default()
	// if err := r.SetTrustedProxies(nil); err != nil {
	// 	return nil, err
	// }

	// r.Use(middleware.Auth())
	v1 := r.Group("/v1")

	{
		controller.InitTaskAPI(v1)
	}
	return r, nil
}
