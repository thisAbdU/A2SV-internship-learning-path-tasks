package usecase

import (
	"context"
	"task-management-api/domain/entities"
	"task-management-api/domain/model"

	"time"
)

type TaskUsecase struct {
	TaskRepository entities.TaskRepository
	contextTimeout time.Duration
}

func NewTaskUsecase(taskRepository entities.TaskRepository) entities.TaskUsecase {
	return &TaskUsecase{
		TaskRepository: taskRepository,
		contextTimeout: 3 * time.Second,
	}
}

func (uc *TaskUsecase) GetTasks(userID string) ([]*model.TaskInfo, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    tasks, err := uc.TaskRepository.GetTasks(ctx, userID)
    if err != nil {
        return nil, err
    }

    var taskInfos []*model.TaskInfo
    for _, task := range tasks {
        taskInfos = append(taskInfos, &model.TaskInfo{
            Title:       task.Title,
            Description: task.Description,
            DueDate:     (time.Now()).Add(3 * 24 * time.Hour).Format(time.RFC3339),
        })
    }

    return taskInfos, nil
}


func (uc *TaskUsecase) GetTaskByID(id string, userID string) (*model.TaskInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	taskEntity, err := uc.TaskRepository.GetTaskByID(ctx, id, userID)
	if err != nil {
		return nil, err
	}

	taskInfo := &model.TaskInfo{
		Title:       taskEntity.Title,
		Description: taskEntity.Description,
	}

	return taskInfo, nil
}

func (uc *TaskUsecase) UpdateTask(id string, updatedTask entities.Task, userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := uc.TaskRepository.UpdateTask(ctx, id, updatedTask, userID)
	if err != nil {
		return err
	}
	return nil
}

func (uc *TaskUsecase) DeleteTask(id string, userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := uc.TaskRepository.DeleteTask(ctx, id, userID)
	if err != nil {
		return err
	}
	return nil
}

func (uc *TaskUsecase) CreateTask(newTask entities.Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := uc.TaskRepository.CreateTask(ctx, newTask)
	if err != nil {
		return err
	}
	return nil
}
