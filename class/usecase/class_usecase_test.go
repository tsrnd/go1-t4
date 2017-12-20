package usecase

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/goweb4/class/repository/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreateClass(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockClassRepository(mockCtrl)
	mockRepo.EXPECT().Create("class A", 20).Return(nil).Times(1)
	r := NewClassUsecase(mockRepo).CreateClass("class A", 20)
	assert.Equal(t, nil, r)
}

func TestCreateClassReturnError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockClassRepository(mockCtrl)
	dummyError := errors.New("error for testing")
	mockRepo.EXPECT().Create("class A", 20).Return(dummyError).Times(1)

	r := NewClassUsecase(mockRepo).CreateClass("class A", 20)
	if r != dummyError {
		t.Fail()
	}
}
