package usecase_test

import (
	"context"
	"errors"
	"task-management-api/domain/entities"
	"task-management-api/domain/mocks"
	"task-management-api/domain/model"
	"task-management-api/usecase"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)


func TestGetTasks(t *testing.T) {
	mockTaskRepository := new(mocks.TaskRepository)
	ctx := context.TODO()
	userID := "testUserID"

	t.Run("success", func(t *testing.T) {

		mockTasks := []*entities.Task{
			{Title: "Task 1", Description: "Description 1"},
			{Title: "Task 2", Description: "Description 2"},
		}
		expectedTaskInfos := []*model.TaskInfo{
			{
				Title:       "Task 1",
				Description: "Description 1",
				DueDate:     (time.Now()).Add(3 * 24 * time.Hour).Format(time.RFC3339),
			},
			{
				Title:       "Task 2",
				Description: "Description 2",
				DueDate:     (time.Now()).Add(3 * 24 * time.Hour).Format(time.RFC3339),
			},
		}

		mockTaskRepository.On("GetTasks", ctx, userID).Return(mockTasks, nil).Once()

		tuc := usecase.NewTaskUsecase(mockTaskRepository)

		tasks, err := tuc.GetTasks(userID)

		assert.NoError(t, err)
		assert.Equal(t, len(expectedTaskInfos), len(tasks))
		for i := range tasks {
			assert.Equal(t, expectedTaskInfos[i].Title, tasks[i].Title)
			assert.Equal(t, expectedTaskInfos[i].Description, tasks[i].Description)
			assert.Equal(t, expectedTaskInfos[i].DueDate, tasks[i].DueDate)
		}

		mockTaskRepository.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		expectedErr := errors.New("repository error")

		mockTaskRepository.On("GetTasks", ctx, userID).Return(nil, expectedErr).Once()

		u := usecase.NewTaskUsecase(mockTaskRepository)

		tasks, err := u.GetTasks(userID)

		assert.Nil(t, tasks)
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)

		mockTaskRepository.AssertExpectations(t)
	})
}

func TestGetTaskByID(t *testing.T) {
	mockTaskRepository := new(mocks.TaskRepository)
	ctx := context.TODO()
	taskID := "testTaskID"
	userID := "testUserID"

	t.Run("success", func(t *testing.T) {
		mockTaskEntity := &entities.Task{
			Title:       "Sample Task",
			Description: "Sample Description",
		}
		expectedTaskInfo := &model.TaskInfo{
			Title:       "Sample Task",
			Description: "Sample Description",
			DueDate:     (time.Now()).Add(3 * 24 * time.Hour).Format(time.RFC3339),
		}

		mockTaskRepository.On("GetTaskByID", ctx, taskID, userID).Return(mockTaskEntity, nil).Once()

		tuc := usecase.NewTaskUsecase(mockTaskRepository)

		taskInfo, err := tuc.GetTaskByID(taskID, userID)

		assert.NoError(t, err)
		assert.Equal(t, expectedTaskInfo.Title, taskInfo.Title)
		assert.Equal(t, expectedTaskInfo.Description, taskInfo.Description)
		assert.Equal(t, expectedTaskInfo.DueDate, taskInfo.DueDate)

		mockTaskRepository.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		expectedErr := errors.New("repository error")

		mockTaskRepository.On("GetTaskByID", ctx, taskID, userID).Return(nil, expectedErr).Once()

		tuc := usecase.NewTaskUsecase(mockTaskRepository)

		taskInfo, err := tuc.GetTaskByID(taskID, userID)

		assert.Nil(t, taskInfo)
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)

		mockTaskRepository.AssertExpectations(t)
	})
}

func TestUpdateTask(t *testing.T) {
	mockTaskRepository := new(mocks.TaskRepository)
	ctx := context.TODO()
	taskID := "testTaskID"
	userID := "testUserID"
	updatedTask := entities.Task{
		Title:       "Updated Task",
		Description: "Updated Description",
	}

	t.Run("success", func(t *testing.T) {
		mockTaskRepository.On("UpdateTask", ctx, taskID, updatedTask, userID).Return(nil).Once()

		tuc := usecase.NewTaskUsecase(mockTaskRepository)

		err := tuc.UpdateTask(taskID, updatedTask, userID)

		assert.NoError(t, err)

		mockTaskRepository.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		expectedErr := errors.New("update error")

		mockTaskRepository.On("UpdateTask", ctx, taskID, updatedTask, userID).Return(expectedErr).Once()

		tuc := usecase.NewTaskUsecase(mockTaskRepository)

		err := tuc.UpdateTask(taskID, updatedTask, userID)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)

		mockTaskRepository.AssertExpectations(t)
	})
}

func TestDeleteTask(t *testing.T) {
	mockTaskRepository := new(mocks.TaskRepository)
	ctx := context.TODO()
	taskID := "testTaskID"
	userID := "testUserID"

	t.Run("success", func(t *testing.T) {
		mockTaskRepository.On("DeleteTask", ctx, taskID, userID).Return(nil).Once()

		tuc := usecase.NewTaskUsecase(mockTaskRepository)

		err := tuc.DeleteTask(taskID, userID)

		assert.NoError(t, err)

		mockTaskRepository.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		expectedErr := errors.New("delete error")

		mockTaskRepository.On("DeleteTask", ctx, taskID, userID).Return(expectedErr).Once()

		tuc := usecase.NewTaskUsecase(mockTaskRepository)

		err := tuc.DeleteTask(taskID, userID)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)

		mockTaskRepository.AssertExpectations(t)
	})
}

func TestCreateTask(t *testing.T) {
	mockTaskRepository := new(mocks.TaskRepository)
	ctx := context.TODO()
	newTask := entities.Task{
		Title:       "New Task",
		Description: "New Task Description",
	}

	t.Run("success", func(t *testing.T) {
		mockTaskRepository.On("CreateTask", ctx, newTask).Return(nil).Once()

		tuc := usecase.NewTaskUsecase(mockTaskRepository)

		err := tuc.CreateTask(newTask)

		assert.NoError(t, err)

		mockTaskRepository.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		expectedErr := errors.New("creation error")

		mockTaskRepository.On("CreateTask", ctx, newTask).Return(expectedErr).Once()

		tuc := usecase.NewTaskUsecase(mockTaskRepository)

		err := tuc.CreateTask(newTask)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)

		mockTaskRepository.AssertExpectations(t)
	})
}



