package controller

import (
	"errors"
	"gin-todo/internal/dto"
	"gin-todo/internal/model"
	"gin-todo/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitTaskAPI(rg *gin.RouterGroup) {
	rgAuth := rg.Group("/tasks")
	{
		rgAuth.GET("/", getTasks)
		rgAuth.GET("/:id", getTask)
		rgAuth.POST("/", createTask)
		rgAuth.PUT("/:id", updateTask)
		rgAuth.DELETE("/:id", deleteTask)
	}
}

func createTask(c *gin.Context) {
	req := dto.AddTaskReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 入口轉
	m := model.Task{
		Name:        req.Name,
		Status:      req.Status,
		Description: req.Description,
	}
	created, err := service.AddTask(&m) // Service 回 Model
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.FromModel(*created))

}

func getTasks(c *gin.Context) {
	tasks, err := service.GetTaskList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func getTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	task, err := service.GetTask(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "error": "task not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func updateTask(c *gin.Context) {
	// 取 URL 裡的 id
	var uri struct {
		ID int `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var req dto.UpdateTaskReq
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	m := model.Task{
		ID:          int(uri.ID),
		Name:        req.Name,
		Status:      req.Status,
		Description: req.Description,
		IsDone:      req.IsDone,
	}

	err := service.UpdateTask(&m)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "error": "task not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "success"})
}

func deleteTask(c *gin.Context) {
	var uri struct {
		ID int `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := service.DeleteTask(uri.ID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "error": "task not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
