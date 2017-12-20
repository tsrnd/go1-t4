package mock_usecase

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gift "github.com/goweb4/gift"
)

// MockGiftUsecase is a mock of GiftUsecase interface
type MockGiftUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockGiftUsecaseMockRecorder
}

// MockGiftUsecaseMockRecorder is the mock recorder for MockGiftUsecase
type MockGiftUsecaseMockRecorder struct {
	mock *MockGiftUsecase
}

// NewMockGiftUsecase creates a new mock instance
func NewMockGiftUsecase(ctrl *gomock.Controller) *MockGiftUsecase {
	mock := &MockGiftUsecase{ctrl: ctrl}
	mock.recorder = &MockGiftUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGiftUsecase) EXPECT() *MockGiftUsecaseMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockGiftUsecase) Create(fromUserID, toUserID, productID int64, message string) (*gift.Gift, error) {
	ret := m.ctrl.Call(m, "Create", fromUserID, toUserID, productID, message)
	ret0, _ := ret[0].(*gift.Gift)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockGiftUsecaseMockRecorder) Create(fromUserID, toUserID, productID, message interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockGiftUsecase)(nil).Create), fromUserID, toUserID, productID, message)
}

// Update mocks base method
func (m *MockGiftUsecase) Update(id, productID int64, message string) (*gift.Gift, error) {
	ret := m.ctrl.Call(m, "Update", id, productID, message)
	ret0, _ := ret[0].(*gift.Gift)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockGiftUsecaseMockRecorder) Update(id, productID, message interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockGiftUsecase)(nil).Update), id, productID, message)
}

// Delete mocks base method
func (m *MockGiftUsecase) Delete(id int64) error {
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockGiftUsecaseMockRecorder) Delete(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockGiftUsecase)(nil).Delete), id)
}

// GetByID mocks base method
func (m *MockGiftUsecase) GetByID(id int64) (*gift.Gift, error) {
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(*gift.Gift)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockGiftUsecaseMockRecorder) GetByID(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockGiftUsecase)(nil).GetByID), id)
}

// GetByFromUserID mocks base method
func (m *MockGiftUsecase) GetByFromUserID(id int64) ([]*gift.Gift, error) {
	ret := m.ctrl.Call(m, "GetByFromUserID", id)
	ret0, _ := ret[0].([]*gift.Gift)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByFromUserID indicates an expected call of GetByFromUserID
func (mr *MockGiftUsecaseMockRecorder) GetByFromUserID(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByFromUserID", reflect.TypeOf((*MockGiftUsecase)(nil).GetByFromUserID), id)
}

// Fetch mocks base method
func (m *MockGiftUsecase) Fetch(offset, limit int64) ([]*gift.Gift, error) {
	ret := m.ctrl.Call(m, "Fetch", offset, limit)
	ret0, _ := ret[0].([]*gift.Gift)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch
func (mr *MockGiftUsecaseMockRecorder) Fetch(offset, limit interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockGiftUsecase)(nil).Fetch), offset, limit)
}
