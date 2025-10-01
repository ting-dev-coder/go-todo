package model

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Status      string    `json:"status"`
	IsDone      bool      `json:"isDone"`
	Description string    `json:"description"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

type AddTask struct {
	Name        string `json:"name"`
	Status      string `json:"status"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func GetAllTasks() ([]Task, error) {
	var todos []Task
	if err := Ctx.DB.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func CreateTask(task *Task) (*Task, error) {
	if err := Ctx.DB.Create(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func GetTaskById(id int) (*Task, error) {
	var todo *Task
	if err := Ctx.DB.First(&todo, id).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func UpdateTask(task *Task) error {
	res := Ctx.DB.Model(&Task{}).
		Where("id = ?", task.ID).
		Updates(map[string]interface{}{
			"name":        task.Name,
			"status":      task.Status,
			"is_done":     task.IsDone,
			"description": task.Description,
		})
	if res.Error != nil {
		return res.Error // 資料庫操作錯誤
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound // 自己補回去，讓上層能判斷是「找不到」
	}
	return nil
}

func DeleteTask(id int) error {
	res := Ctx.DB.Delete(&Task{}, id)

	if res.Error != nil {
		return res.Error // 資料庫操作錯誤
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound // 自己補回去，讓上層能判斷是「找不到」
	}
	return nil
}
