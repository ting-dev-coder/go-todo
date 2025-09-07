package model

import "time"

type Task struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Status      string    `json:"status"`
	IsDone      bool      `json:"isDone"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type AddTask struct {
	Name        string `json:"name"`
	Status      string `json:"status"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
