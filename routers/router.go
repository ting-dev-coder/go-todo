package routers

import (
	"gin-todo/controller"
	"gin-todo/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine) {
	//user = r.Group("/users",{	})
	v1 := r.Group("/api/v1")

	task := v1.Group("/tasks", middleware.Auth())
	{
		task.GET("/", controller.GetTasks)
		task.GET("/:id", controller.GetTask)
		task.POST("/", controller.CreateTask)
		task.PATCH("/:id", controller.UpdateTask)
		task.DELETE("/:id", controller.DeleteTask)
	}

}
