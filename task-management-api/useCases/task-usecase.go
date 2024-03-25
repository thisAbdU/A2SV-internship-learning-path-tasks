package task

import (
	task "example/GO-PRACTICE-EXERCISE/GO-API-exercise/entities"
)

type TaskUsecase struct {
    TaskRepository taskrepositroy.TaskRepository
}

func (uc *TaskUsecase) GetTasks() ([]task.Task, error) {
    return uc.TaskRepository.GetTasks()
}

func (uc *TaskUsecase) GetTaskByID(id string) (*task.Task, error) {
    return uc.TaskRepository.GetTaskByID(id)
}

func (uc *TaskUsecase) UpdateTask(id string, updatedTask task.Task) error {
    return uc.TaskRepository.UpdateTask(id, updatedTask)
}

func (uc *TaskUsecase) DeleteTask(id string) error {
    return uc.TaskRepository.DeleteTask(id)
}

func (uc *TaskUsecase) CreateTask(newTask task.Task) error {
    return uc.TaskRepository.CreateTask(newTask)
}
