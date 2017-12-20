package mock_repository

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gift "github.com/goweb4/gift"
)

// MockGiftRepository is a mock of GiftRepository interface
type MockGiftRepository struct {
	ctrl     *gomock.Controller
	recorder *MockGiftRepositoryMockRecorder
}

// MockGiftRepositoryMockRecorder is the mock recorder for MockGiftRepository
type MockGiftRepositoryMockRecorder struct {
	mock *MockGiftRepository
}

// NewMockGiftRepository creates a new mock instance
func NewMockGiftRepository(ctrl *gomock.Controller) *MockGiftRepository {
	mock := &MockGiftRepository{ctrl: ctrl}
	mock.recorder = &MockGiftRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGiftRepository) EXPECT() *MockGiftRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockGiftRepository) Create(fromUserID, toUserID, productID int64, message string) (int64, error) {
	ret := m.ctrl.Call(m, "Create", fromUserID, toUserID, productID, message)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockGiftRepositoryMockRecorder) Create(fromUserID, toUserID, productID, message interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockGiftRepository)(nil).Create), fromUserID, toUserID, productID, message)
}

// Update mocks base method
func (m *MockGiftRepository) Update(id, productID int64, message string) error {
	ret := m.ctrl.Call(m, "Update", id, productID, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockGiftRepositoryMockRecorder) Update(id, productID, message interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockGiftRepository)(nil).Update), id, productID, message)
}

// Delete mocks base method
func (m *MockGiftRepository) Delete(id int64) error {
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockGiftRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockGiftRepository)(nil).Delete), id)
}

// GetByID mocks base method
func (m *MockGiftRepository) GetByID(id int64) (*gift.Gift, error) {
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(*gift.Gift)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockGiftRepositoryMockRecorder) GetByID(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockGiftRepository)(nil).GetByID), id)
}

// GetByFromUserID mocks base method
func (m *MockGiftRepository) GetByFromUserID(id int64) ([]*gift.Gift, error) {
	ret := m.ctrl.Call(m, "GetByFromUserID", id)
	ret0, _ := ret[0].([]*gift.Gift)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByFromUserID indicates an expected call of GetByFromUserID
func (mr *MockGiftRepositoryMockRecorder) GetByFromUserID(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByFromUserID", reflect.TypeOf((*MockGiftRepository)(nil).GetByFromUserID), id)
}

// Fetch mocks base method
func (m *MockGiftRepository) Fetch(offset, limit int64) ([]*gift.Gift, error) {
	ret := m.ctrl.Call(m, "Fetch", offset, limit)
	ret0, _ := ret[0].([]*gift.Gift)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch
func (mr *MockGiftRepositoryMockRecorder) Fetch(offset, limit interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockGiftRepository)(nil).Fetch), offset, limit)
}
