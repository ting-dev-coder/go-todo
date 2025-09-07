package service

import (
	"gin-todo/model"
	"gin-todo/repository"
)

func GetTaskList() ([]model.Task, error) {
	tasks, err := repository.GetAllTasks()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func AddTask(task *model.Task) (*model.Task, error) {

	task.IsDone = false

	task, err := repository.CreateTask(task)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func GetTask(id int) (*model.Task, error) {
	task, err := repository.GetTaskById(id)

	if err != nil {
		return nil, err
	}
	return task, nil
}
