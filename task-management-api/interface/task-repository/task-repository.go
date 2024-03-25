package task

import task "example/GO-PRACTICE-EXERCISE/GO-API-exercise/entities"

type TaskRepository interface {
    GetTasks() ([]task.Task, error)
    GetTaskByID(id string) (*task.Task, error)
    UpdateTask(id string, updatedTask task.Task) error
    DeleteTask(id string) error
    CreateTask(newTask task.Task) error
}

