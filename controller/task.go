package controller

import (
	"errors"
	"gin-todo/dto"
	"gin-todo/model"
	"gin-todo/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetTasks(c *gin.Context) {
	tasks, err := service.GetTaskList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func GetTask(c *gin.Context) {
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

func CreateTask(c *gin.Context) {
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

func UpdateTask(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "task updated"})
	// var req dto.UpdateTaskReq
	// if err := c.ShouldBindBodyWithJSON(&req); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// m := model.Task{
	// 	ID:          req.ID,
	// 	Name:        req.Name,
	// 	Status:      req.Status,
	// 	Description: req.Description,
	// 	IsDone:      req.IsDone,
	// }

	// _, err := service.UpdateTask(&m)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{"message": "task updated"})
}

func DeleteTask(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "task deleted"})
}
