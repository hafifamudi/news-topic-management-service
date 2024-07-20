// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	model "news-topic-management-service/internal/core/news/model"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// MockNewsRepository is an autogenerated mock type for the NewsRepository type
type MockNewsRepository struct {
	mock.Mock
}

type MockNewsRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockNewsRepository) EXPECT() *MockNewsRepository_Expecter {
	return &MockNewsRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: news
func (_m *MockNewsRepository) Create(news *model.News) (*model.News, error) {
	ret := _m.Called(news)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *model.News
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.News) (*model.News, error)); ok {
		return rf(news)
	}
	if rf, ok := ret.Get(0).(func(*model.News) *model.News); ok {
		r0 = rf(news)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.News)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.News) error); ok {
		r1 = rf(news)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockNewsRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockNewsRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - news *model.News
func (_e *MockNewsRepository_Expecter) Create(news interface{}) *MockNewsRepository_Create_Call {
	return &MockNewsRepository_Create_Call{Call: _e.mock.On("Create", news)}
}

func (_c *MockNewsRepository_Create_Call) Run(run func(news *model.News)) *MockNewsRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.News))
	})
	return _c
}

func (_c *MockNewsRepository_Create_Call) Return(_a0 *model.News, _a1 error) *MockNewsRepository_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNewsRepository_Create_Call) RunAndReturn(run func(*model.News) (*model.News, error)) *MockNewsRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: newsID
func (_m *MockNewsRepository) Delete(newsID uuid.UUID) (*model.News, error) {
	ret := _m.Called(newsID)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 *model.News
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (*model.News, error)); ok {
		return rf(newsID)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) *model.News); ok {
		r0 = rf(newsID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.News)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(newsID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockNewsRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockNewsRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - newsID uuid.UUID
func (_e *MockNewsRepository_Expecter) Delete(newsID interface{}) *MockNewsRepository_Delete_Call {
	return &MockNewsRepository_Delete_Call{Call: _e.mock.On("Delete", newsID)}
}

func (_c *MockNewsRepository_Delete_Call) Run(run func(newsID uuid.UUID)) *MockNewsRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uuid.UUID))
	})
	return _c
}

func (_c *MockNewsRepository_Delete_Call) Return(_a0 *model.News, _a1 error) *MockNewsRepository_Delete_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNewsRepository_Delete_Call) RunAndReturn(run func(uuid.UUID) (*model.News, error)) *MockNewsRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Find provides a mock function with given fields: id
func (_m *MockNewsRepository) Find(id uuid.UUID) (*model.News, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Find")
	}

	var r0 *model.News
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (*model.News, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) *model.News); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.News)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockNewsRepository_Find_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Find'
type MockNewsRepository_Find_Call struct {
	*mock.Call
}

// Find is a helper method to define mock.On call
//   - id uuid.UUID
func (_e *MockNewsRepository_Expecter) Find(id interface{}) *MockNewsRepository_Find_Call {
	return &MockNewsRepository_Find_Call{Call: _e.mock.On("Find", id)}
}

func (_c *MockNewsRepository_Find_Call) Run(run func(id uuid.UUID)) *MockNewsRepository_Find_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uuid.UUID))
	})
	return _c
}

func (_c *MockNewsRepository_Find_Call) Return(_a0 *model.News, _a1 error) *MockNewsRepository_Find_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNewsRepository_Find_Call) RunAndReturn(run func(uuid.UUID) (*model.News, error)) *MockNewsRepository_Find_Call {
	_c.Call.Return(run)
	return _c
}

// GetAll provides a mock function with given fields: status, topicID
func (_m *MockNewsRepository) GetAll(status *string, topicID *uuid.UUID) ([]model.News, error) {
	ret := _m.Called(status, topicID)

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []model.News
	var r1 error
	if rf, ok := ret.Get(0).(func(*string, *uuid.UUID) ([]model.News, error)); ok {
		return rf(status, topicID)
	}
	if rf, ok := ret.Get(0).(func(*string, *uuid.UUID) []model.News); ok {
		r0 = rf(status, topicID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.News)
		}
	}

	if rf, ok := ret.Get(1).(func(*string, *uuid.UUID) error); ok {
		r1 = rf(status, topicID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockNewsRepository_GetAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAll'
type MockNewsRepository_GetAll_Call struct {
	*mock.Call
}

// GetAll is a helper method to define mock.On call
//   - status *string
//   - topicID *uuid.UUID
func (_e *MockNewsRepository_Expecter) GetAll(status interface{}, topicID interface{}) *MockNewsRepository_GetAll_Call {
	return &MockNewsRepository_GetAll_Call{Call: _e.mock.On("GetAll", status, topicID)}
}

func (_c *MockNewsRepository_GetAll_Call) Run(run func(status *string, topicID *uuid.UUID)) *MockNewsRepository_GetAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*string), args[1].(*uuid.UUID))
	})
	return _c
}

func (_c *MockNewsRepository_GetAll_Call) Return(_a0 []model.News, _a1 error) *MockNewsRepository_GetAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNewsRepository_GetAll_Call) RunAndReturn(run func(*string, *uuid.UUID) ([]model.News, error)) *MockNewsRepository_GetAll_Call {
	_c.Call.Return(run)
	return _c
}

// Preload provides a mock function with given fields: news
func (_m *MockNewsRepository) Preload(news *model.News) (*model.News, error) {
	ret := _m.Called(news)

	if len(ret) == 0 {
		panic("no return value specified for Preload")
	}

	var r0 *model.News
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.News) (*model.News, error)); ok {
		return rf(news)
	}
	if rf, ok := ret.Get(0).(func(*model.News) *model.News); ok {
		r0 = rf(news)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.News)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.News) error); ok {
		r1 = rf(news)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockNewsRepository_Preload_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Preload'
type MockNewsRepository_Preload_Call struct {
	*mock.Call
}

// Preload is a helper method to define mock.On call
//   - news *model.News
func (_e *MockNewsRepository_Expecter) Preload(news interface{}) *MockNewsRepository_Preload_Call {
	return &MockNewsRepository_Preload_Call{Call: _e.mock.On("Preload", news)}
}

func (_c *MockNewsRepository_Preload_Call) Run(run func(news *model.News)) *MockNewsRepository_Preload_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.News))
	})
	return _c
}

func (_c *MockNewsRepository_Preload_Call) Return(_a0 *model.News, _a1 error) *MockNewsRepository_Preload_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNewsRepository_Preload_Call) RunAndReturn(run func(*model.News) (*model.News, error)) *MockNewsRepository_Preload_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: id, news
func (_m *MockNewsRepository) Update(id uuid.UUID, news *model.News) (*model.News, error) {
	ret := _m.Called(id, news)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *model.News
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID, *model.News) (*model.News, error)); ok {
		return rf(id, news)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID, *model.News) *model.News); ok {
		r0 = rf(id, news)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.News)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID, *model.News) error); ok {
		r1 = rf(id, news)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockNewsRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockNewsRepository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - id uuid.UUID
//   - news *model.News
func (_e *MockNewsRepository_Expecter) Update(id interface{}, news interface{}) *MockNewsRepository_Update_Call {
	return &MockNewsRepository_Update_Call{Call: _e.mock.On("Update", id, news)}
}

func (_c *MockNewsRepository_Update_Call) Run(run func(id uuid.UUID, news *model.News)) *MockNewsRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uuid.UUID), args[1].(*model.News))
	})
	return _c
}

func (_c *MockNewsRepository_Update_Call) Return(_a0 *model.News, _a1 error) *MockNewsRepository_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNewsRepository_Update_Call) RunAndReturn(run func(uuid.UUID, *model.News) (*model.News, error)) *MockNewsRepository_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockNewsRepository creates a new instance of MockNewsRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockNewsRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockNewsRepository {
	mock := &MockNewsRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
