// Code generated by MockGen. DO NOT EDIT.
// Source: ../../repository.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	book "github.com/gmhafiz/go8/internal/domain/book"
	models "github.com/gmhafiz/go8/internal/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRepository is a mock of Repository interface
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockRepository) Create(ctx context.Context, book *models.Book) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, book)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockRepositoryMockRecorder) Create(ctx, book interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), ctx, book)
}

// All mocks base method
func (m *MockRepository) All(ctx context.Context, f *book.Filter) ([]*models.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, f)
	ret0, _ := ret[0].([]*models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All
func (mr *MockRepositoryMockRecorder) All(ctx, f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockRepository)(nil).All), ctx, f)
}

// Read mocks base method
func (m *MockRepository) Read(ctx context.Context, bookID int64) (*models.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, bookID)
	ret0, _ := ret[0].(*models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockRepositoryMockRecorder) Read(ctx, bookID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockRepository)(nil).Read), ctx, bookID)
}

// Update mocks base method
func (m *MockRepository) Update(ctx context.Context, book *models.Book) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, book)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockRepositoryMockRecorder) Update(ctx, book interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepository)(nil).Update), ctx, book)
}

// Delete mocks base method
func (m *MockRepository) Delete(ctx context.Context, bookID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, bookID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockRepositoryMockRecorder) Delete(ctx, bookID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), ctx, bookID)
}

// Search mocks base method
func (m *MockRepository) Search(ctx context.Context, req *book.Filter) ([]*models.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", ctx, req)
	ret0, _ := ret[0].([]*models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockRepositoryMockRecorder) Search(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockRepository)(nil).Search), ctx, req)
}
