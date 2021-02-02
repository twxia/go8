// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_book is a generated GoMock package.
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
func (m *MockRepository) All(ctx context.Context) ([]*models.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx)
	ret0, _ := ret[0].([]*models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All
func (mr *MockRepositoryMockRecorder) All(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockRepository)(nil).All), ctx)
}

// Find mocks base method
func (m *MockRepository) Find(ctx context.Context, bookID int64) (*models.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, bookID)
	ret0, _ := ret[0].(*models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockRepositoryMockRecorder) Find(ctx, bookID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockRepository)(nil).Find), ctx, bookID)
}

// Update mocks base method
func (m *MockRepository) Update(ctx context.Context, book *models.Book) (*models.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, book)
	ret0, _ := ret[0].(*models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
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
func (m *MockRepository) Search(ctx context.Context, req *book.Request) ([]*models.Book, error) {
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

// MockTest is a mock of Test interface
type MockTest struct {
	ctrl     *gomock.Controller
	recorder *MockTestMockRecorder
}

// MockTestMockRecorder is the mock recorder for MockTest
type MockTestMockRecorder struct {
	mock *MockTest
}

// NewMockTest creates a new mock instance
func NewMockTest(ctrl *gomock.Controller) *MockTest {
	mock := &MockTest{ctrl: ctrl}
	mock.recorder = &MockTestMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTest) EXPECT() *MockTestMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockTest) Create(ctx context.Context, book *models.Book) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, book)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockTestMockRecorder) Create(ctx, book interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTest)(nil).Create), ctx, book)
}

// All mocks base method
func (m *MockTest) All(ctx context.Context) ([]*models.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx)
	ret0, _ := ret[0].([]*models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All
func (mr *MockTestMockRecorder) All(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockTest)(nil).All), ctx)
}

// Find mocks base method
func (m *MockTest) Find(ctx context.Context, bookID int64) (*models.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, bookID)
	ret0, _ := ret[0].(*models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockTestMockRecorder) Find(ctx, bookID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockTest)(nil).Find), ctx, bookID)
}

// Update mocks base method
func (m *MockTest) Update(ctx context.Context, book *models.Book) (*models.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, book)
	ret0, _ := ret[0].(*models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockTestMockRecorder) Update(ctx, book interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTest)(nil).Update), ctx, book)
}

// Delete mocks base method
func (m *MockTest) Delete(ctx context.Context, bookID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, bookID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockTestMockRecorder) Delete(ctx, bookID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTest)(nil).Delete), ctx, bookID)
}

// Search mocks base method
func (m *MockTest) Search(ctx context.Context, req *book.Request) ([]*models.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", ctx, req)
	ret0, _ := ret[0].([]*models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockTestMockRecorder) Search(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockTest)(nil).Search), ctx, req)
}

// Close mocks base method
func (m *MockTest) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockTestMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockTest)(nil).Close))
}

// Drop mocks base method
func (m *MockTest) Drop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Drop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Drop indicates an expected call of Drop
func (mr *MockTestMockRecorder) Drop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Drop", reflect.TypeOf((*MockTest)(nil).Drop))
}

// Up mocks base method
func (m *MockTest) Up() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Up")
	ret0, _ := ret[0].(error)
	return ret0
}

// Up indicates an expected call of Up
func (mr *MockTestMockRecorder) Up() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Up", reflect.TypeOf((*MockTest)(nil).Up))
}
