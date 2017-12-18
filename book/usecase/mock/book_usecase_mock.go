// Code generated by MockGen. DO NOT EDIT.
// Source: book_usecase.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	"github.com/goweb4/book"
)

// MockBookUsecase is a mock of BookUsecase interface
type MockBookUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockBookUsecaseMockRecorder
}

// MockBookUsecaseMockRecorder is the mock recorder for MockBookUsecase
type MockBookUsecaseMockRecorder struct {
	mock *MockBookUsecase
}

// NewMockBookUsecase creates a new mock instance
func NewMockBookUsecase(ctrl *gomock.Controller) *MockBookUsecase {
	mock := &MockBookUsecase{ctrl: ctrl}
	mock.recorder = &MockBookUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBookUsecase) EXPECT() *MockBookUsecaseMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockBookUsecase) Create(title, description string, userID int64) (*book.Book, error) {
	ret := m.ctrl.Call(m, "Create", title, description, userID)
	ret0, _ := ret[0].(*book.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockBookUsecaseMockRecorder) Create(title, description, userID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockBookUsecase)(nil).Create), title, description, userID)
}

// GetByID mocks base method
func (m *MockBookUsecase) GetByID(id int64) (*book.Book, error) {
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(*book.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockBookUsecaseMockRecorder) GetByID(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockBookUsecase)(nil).GetByID), id)
}

// GetByTitle mocks base method
func (m *MockBookUsecase) GetByTitle(title string) (*book.Book, error) {
	ret := m.ctrl.Call(m, "GetByTitle", title)
	ret0, _ := ret[0].(*book.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByTitle indicates an expected call of GetByTitle
func (mr *MockBookUsecaseMockRecorder) GetByTitle(title interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByTitle", reflect.TypeOf((*MockBookUsecase)(nil).GetByTitle), title)
}

// Update mocks base method
func (m *MockBookUsecase) Update(bookID int64, title, description string) (*book.Book, error) {
	ret := m.ctrl.Call(m, "Update", bookID, title, description)
	ret0, _ := ret[0].(*book.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockBookUsecaseMockRecorder) Update(bookID, title, description interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockBookUsecase)(nil).Update), bookID, title, description)
}

// Delete mocks base method
func (m *MockBookUsecase) Delete(id int64) error {
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockBookUsecaseMockRecorder) Delete(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBookUsecase)(nil).Delete), id)
}

// Fetch mocks base method
func (m *MockBookUsecase) Fetch(offset, limit int64) ([]*book.Book, error) {
	ret := m.ctrl.Call(m, "Fetch", offset, limit)
	ret0, _ := ret[0].([]*book.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch
func (mr *MockBookUsecaseMockRecorder) Fetch(offset, limit interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockBookUsecase)(nil).Fetch), offset, limit)
}