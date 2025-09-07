package repository

import (
	"gin-todo/model"
)

func GetAllTasks() ([]model.Task, error) {
	var todos []model.Task
	if err := DB.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func CreateTask(task *model.Task) (*model.Task, error) {
	if err := DB.Create(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func GetTaskById(id int) (*model.Task, error) {
	var todo *model.Task
	if err := DB.First(&todo, id).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func UpdateTask(task *model.Task) error {
	return DB.Save(task).Error
}

func DeleteTask(id int) error {
	return DB.Delete(&model.Task{}, id).Error
}
