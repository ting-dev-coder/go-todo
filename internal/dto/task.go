package dto

import "gin-todo/internal/model"

type UpdateTaskReq struct {
	Name        string `json:"name" binding:"required"`
	Status      string `json:"status" binding:"required"`
	Description string `json:"description" binding:"required"`
	IsDone      bool   `json:"isDone"`
}

type AddTaskReq struct {
	Name        string `json:"name" binding:"required"`
	Status      string `json:"status" binding:"required"`
	Description string `json:"description" binding:"required"`
	IsDone      bool   `json:"isDone"`
}

type TaskResp struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	IsDone    bool   `json:"isDone"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func FromModel(m model.Task) TaskResp {
	return TaskResp{
		ID:        uint(m.ID),
		Title:     m.Name,
		IsDone:    m.IsDone,
		CreatedAt: m.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: m.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
