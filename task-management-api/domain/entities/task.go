package entities

import (
	"context"
	"task-management-api/domain/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
    ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserID 		string `json:"user_id"`
    UserName    string `json:"username"`
    Password    string `json:"password"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Status      string `json:"status"`
}

type TaskRepository interface {
	GetTasks(ctx context.Context, userID string) ([]*model.TaskInfo, error)
	GetTaskByID(ctx context.Context, id string, userID string) (*Task, error)
	UpdateTask(ctx context.Context, id string, updatedTask Task, userID string) error
	DeleteTask(ctx context.Context, id string, userID string) error
	CreateTask(ctx context.Context, newTask Task) error
}

type TaskUsecase interface {
	GetTasks() (*[]*model.TaskInfo, error)
	GetTaskByID(ctx context.Context, id string, userID string) (*model.TaskInfo, error)
	UpdateTask(id string, userID string) (*model.TaskUpdate, error)
	DeleteTask(id string, userID string) error
	CreateTask(newTask Task) error
}