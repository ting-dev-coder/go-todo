package service

import (
	"gin-todo/internal/model"
)

func GetTaskList() ([]model.Task, error) {
	tasks, err := model.GetAllTasks()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func AddTask(task *model.Task) (*model.Task, error) {

	task.IsDone = false

	task, err := model.CreateTask(task)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func GetTask(id int) (*model.Task, error) {
	task, err := model.GetTaskById(id)

	if err != nil {
		return nil, err
	}
	return task, nil
}

func UpdateTask(task *model.Task) error {

	err := model.UpdateTask(task)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTask(id int) error {

	err := model.DeleteTask(id)
	if err != nil {
		return err
	}
	return nil
}
