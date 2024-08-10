package controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"task-management-api/controller"
	"task-management-api/domain/entities"
	"task-management-api/domain/mocks"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateTask(t *testing.T) {
	mockUsecase := new(mocks.TaskUsecase)
	mockEnvironment := new(mocks.Environment)

	router := gin.Default()

	tc := controller.NewTaskController(mockEnvironment, mockUsecase)

	router.POST("/tasks", tc.CreateTask)

	t.Run("success", func(t *testing.T) {
		mockUsecase.On("CreateTask", mock.AnythingOfType("entities.Task")).Return(nil)

		newTask := entities.Task{
			Title:       "Test Task",
			Description: "This is a test task",
		}

		taskJSON, _ := json.Marshal(newTask)

		req, _ := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(taskJSON))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req
		c.Set("user_id", "test_user_id")

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		assert.JSONEq(t, `{"message": "Task created successfully"}`, w.Body.String())

		mockUsecase.AssertCalled(t, "CreateTask", mock.MatchedBy(func(task entities.Task) bool {
			return task.Title == newTask.Title && task.Description == newTask.Description && task.UserID == "test_user_id"
		}))
	})

	t.Run("error", func(t *testing.T) {
		t.Run("unauthorized", func(t *testing.T) {
			newTask := entities.Task{
				Title:       "Test Task",
				Description: "This is a test task",
			}

			taskJSON, _ := json.Marshal(newTask)

			req, _ := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(taskJSON))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusUnauthorized, w.Code)
			assert.JSONEq(t, `{"message": "Please sign up to create a task"}`, w.Body.String())
		})

		t.Run("bad request", func(t *testing.T) {
			invalidTaskJSON := `{"title": ""}`

			req, _ := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBufferString(invalidTaskJSON))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = req
			c.Set("user_id", "test_user_id")

			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusBadRequest, w.Code)
			assert.JSONEq(t, `{"message": "Bad Request"}`, w.Body.String())
		})

		t.Run("internal server error", func(t *testing.T) {
			mockUsecase.On("CreateTask", mock.AnythingOfType("entities.Task")).Return(errors.New("some error"))

			newTask := entities.Task{
				Title:       "Test Task",
				Description: "This is a test task",
			}

			taskJSON, _ := json.Marshal(newTask)

			req, _ := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(taskJSON))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = req
			c.Set("user_id", "test_user_id")

			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusInternalServerError, w.Code)
			assert.JSONEq(t, `{"message": "Internal Server Error"}`, w.Body.String())
		})
	})
}

func TestGetTasks(t *testing.T) {
	mockUsecase := new(mocks.TaskUsecase)
	mockEnvironment := new(mocks.Environment)

	router := gin.Default()

	tc := controller.NewTaskController(mockEnvironment, mockUsecase)

	router.GET("/tasks", tc.GetTasks)

	t.Run("unauthorized", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.JSONEq(t, `{"error": "Please log in to acess tasks"}`, w.Body.String())
	})

	t.Run("internal server error - userID conversion", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req
		c.Set("user_id", 123)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.JSONEq(t, `{"message": "Internal server error"}`, w.Body.String())
	})

	t.Run("internal server error - GetTasks failure", func(t *testing.T) {
		mockUsecase.On("GetTasks", "test_user_id").Return(nil, errors.New("some error"))
		req, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req
		c.Set("user_id", "test_user_id")
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.JSONEq(t, `{"message": "error retrieving tasks"}`, w.Body.String())
	})

	t.Run("success", func(t *testing.T) {
		mockTasks := []entities.Task{
			{ID: primitive.NewObjectID(), Title: "Task 1", Description: "Description 1"},
			{ID: primitive.NewObjectID(), Title: "Task 2", Description: "Description 2"},
		}
		mockUsecase.On("GetTasks", "test_user_id").Return(mockTasks, nil)
		req, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req
		c.Set("user_id", "test_user_id")
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		expectedResponse, _ := json.Marshal(gin.H{"tasks": mockTasks})
		assert.JSONEq(t, string(expectedResponse), w.Body.String())
	})
}

func TestGetTaskByID(t *testing.T) {
	mockUsecase := new(mocks.TaskUsecase)
	mockEnvironment := new(mocks.Environment)

	router := gin.Default()

	tc := controller.NewTaskController(mockEnvironment, mockUsecase)

	router.GET("/tasks/:id", tc.GetTaskByID)

	t.Run("unauthorized", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/tasks/1", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.JSONEq(t, `{"error": "Please log in to acess tasks"}`, w.Body.String())
	})

	t.Run("task not found", func(t *testing.T) {
		mockUsecase.On("GetTaskByID", "1", "test_user_id").Return(nil, errors.New("mongo: no documents in result"))
		req, _ := http.NewRequest(http.MethodGet, "/tasks/1", nil)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req
		c.Set("user_id", "test_user_id")
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.JSONEq(t, `{"message": "Task not found"}`, w.Body.String())
	})

	t.Run("internal server error", func(t *testing.T) {
		mockUsecase.On("GetTaskByID", "1", "test_user_id").Return(nil, errors.New("some error"))
		req, _ := http.NewRequest(http.MethodGet, "/tasks/1", nil)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req
		c.Set("user_id", "test_user_id")
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.JSONEq(t, `{"message": "error retrieving task"}`, w.Body.String())
	})

	t.Run("success", func(t *testing.T) {
		mockTask := &entities.Task{ID: primitive.NewObjectID(), Title: "Task 1", Description: "Description 1"}
		mockUsecase.On("GetTaskByID", mockTask.ID, "test_user_id").Return(mockTask, nil)
		req, _ := http.NewRequest(http.MethodGet, "/tasks/1", nil)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req
		c.Set("user_id", "test_user_id")
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		expectedResponse, _ := json.Marshal(gin.H{"task": mockTask})
		assert.JSONEq(t, string(expectedResponse), w.Body.String())
	})
}

func TestUpdateTask(t *testing.T) {
	mockUsecase := new(mocks.TaskUsecase)
	mockEnvironment := new(mocks.Environment)

	router := gin.Default()

	tc := controller.NewTaskController(mockEnvironment, mockUsecase)

	router.PUT("/tasks/:id", tc.UpdateTask)

	t.Run("unauthorized", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPut, "/tasks/1", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.JSONEq(t, `{"error": "Please log in to acess tasks"}`, w.Body.String())
	})

	t.Run("task not found", func(t *testing.T) {
		mockUsecase.On("UpdateTask", mock.AnythingOfType("entities.Task")).Return(errors.New("mongo: no documents in result"))
		req, _ := http.NewRequest(http.MethodPut, "/tasks/1", nil)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req
		c.Set("user_id", "test_user_id")
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.JSONEq(t, `{"message": "Task not found"}`, w.Body.String())
	})

	t.Run("internal server error", func(t *testing.T) {
		mockUsecase.On("UpdateTask", mock.AnythingOfType("entities.Task")).Return(errors.New("some error"))
		req, _ := http.NewRequest(http.MethodPut, "/tasks/1", nil)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req
		c.Set("user_id", "test_user_id")
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.JSONEq(t, `{"message": "error updating task"}`, w.Body.String())
	})

	t.Run("success", func(t *testing.T) {
		mockTask := &entities.Task{ID: primitive.NewObjectID(), Title: "Task 1", Description: "Description 1"}
		mockUsecase.On("UpdateTask", mock.AnythingOfType("entities.Task")).Return(nil)
		taskJSON, _ := json.Marshal(mockTask)

		taskID := mockTask.ID.Hex()
		req, _ := http.NewRequest(http.MethodPut, "/tasks/" + taskID, bytes.NewBuffer(taskJSON))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req
		c.Set("user_id", "test_user_id")
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, `{"message": "Task updated successfully"}`, w.Body.String())
	})
}

func TestDeleteTask(t *testing.T) {
	mockUsecase := new(mocks.TaskUsecase)
	mockEnvironment := new(mocks.Environment)

	router := gin.Default()

	tc := controller.NewTaskController(mockEnvironment, mockUsecase)

	router.DELETE("/tasks/:id", tc.DeleteTask)

	t.Run("unauthorized", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodDelete, "/tasks/1", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.JSONEq(t, `{"error": "Please log in to acess tasks"}`, w.Body.String())
	})

	t.Run("task not found", func(t *testing.T) {
		mockUsecase.On("DeleteTask", "1", "test_user_id").Return(errors.New("mongo: no documents in result"))
		req, _ := http.NewRequest(http.MethodDelete, "/tasks/1", nil)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req
		c.Set("user_id", "test_user_id")
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.JSONEq(t, `{"message": "Task not found"}`, w.Body.String())
	})

	t.Run("internal server error", func(t *testing.T) {
		mockUsecase.On("DeleteTask", "1", "test_user_id").Return(errors.New("some error"))
		req, _ := http.NewRequest(http.MethodDelete, "/tasks/1", nil)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req
		c.Set("user_id", "test_user_id")
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.JSONEq(t, `{"message": "error deleting task"}`, w.Body.String())
	})

	t.Run("success", func(t *testing.T) {
		mockUsecase.On("DeleteTask", "1", "test_user_id").Return(nil)
		req, _ := http.NewRequest(http.MethodDelete, "/tasks/1", nil)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req
		c.Set("user_id", "test_user_id") 
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, `{"message": "Task deleted successfully"}`, w.Body.String())
	})
}