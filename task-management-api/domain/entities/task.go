package entities

import (
	"context"
	"example/GO-PRACTICE-EXERCISE/GO-API-exercise/domain/model"
)

type Task struct {
    ID          string `json:"id" bson:"_id"`
    UserName        string `json:"username"`
    Password    string `json:"password"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Status      string `json:"status"`
}

type TaskRepository interface {
	GetTasks(ctx context.Context, param string) ([]Task, error)
	GetTaskByID(ctx context.Context, id string) (*model.TaskInfo, error)
	UpdateTask(ctx context.Context, id string, updatedTask Task) error
	DeleteTask(ctx context.Context, id string) error
	CreateTask(nctx context.Context, ewTask Task) error
}

type TaskUsecase interface {
	GetTasks() (*[]*model.TaskInfo, error)
	GetTaskByID(ctx context.Context, id string) (*model.TaskInfo, error)
	UpdateTask(id string) (*model.TaskUpdate, error)
	DeleteTask(id string) error
	CreateTask(newTask Task) error
}