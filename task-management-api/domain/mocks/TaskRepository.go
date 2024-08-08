// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"
	entities "task-management-api/domain/entities"

	mock "github.com/stretchr/testify/mock"

	model "task-management-api/domain/model"
)

// TaskRepository is an autogenerated mock type for the TaskRepository type
type TaskRepository struct {
	mock.Mock
}

// CreateTask provides a mock function with given fields: ctx, newTask
func (_m *TaskRepository) CreateTask(ctx context.Context, newTask entities.Task) error {
	ret := _m.Called(ctx, newTask)

	if len(ret) == 0 {
		panic("no return value specified for CreateTask")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entities.Task) error); ok {
		r0 = rf(ctx, newTask)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTask provides a mock function with given fields: ctx, id, userID
func (_m *TaskRepository) DeleteTask(ctx context.Context, id string, userID string) error {
	ret := _m.Called(ctx, id, userID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTask")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, id, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTaskByID provides a mock function with given fields: ctx, id, userID
func (_m *TaskRepository) GetTaskByID(ctx context.Context, id string, userID string) (*entities.Task, error) {
	ret := _m.Called(ctx, id, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetTaskByID")
	}

	var r0 *entities.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*entities.Task, error)); ok {
		return rf(ctx, id, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *entities.Task); ok {
		r0 = rf(ctx, id, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, id, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTasks provides a mock function with given fields: ctx, userID
func (_m *TaskRepository) GetTasks(ctx context.Context, userID string) ([]*model.TaskInfo, error) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetTasks")
	}

	var r0 []*model.TaskInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*model.TaskInfo, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*model.TaskInfo); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.TaskInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTask provides a mock function with given fields: ctx, id, updatedTask, userID
func (_m *TaskRepository) UpdateTask(ctx context.Context, id string, updatedTask entities.Task, userID string) error {
	ret := _m.Called(ctx, id, updatedTask, userID)

	if len(ret) == 0 {
		panic("no return value specified for UpdateTask")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, entities.Task, string) error); ok {
		r0 = rf(ctx, id, updatedTask, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewTaskRepository creates a new instance of TaskRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTaskRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *TaskRepository {
	mock := &TaskRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}