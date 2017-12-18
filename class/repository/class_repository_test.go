package repository

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/goweb4/class/repository/mocks"
	"github.com/stretchr/testify/assert"
)

func TestClassCreate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockClassRepo := mocks.NewMockClassRepository(mockCtrl)
	mockClassRepo.EXPECT().Create("class A", 20).Return(nil).Times(1)
	r := mockClassRepo.Create("class A", 20)
	assert.Equal(t, nil, r)
}

func TestClassCreateReturnError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockClassRepo := mocks.NewMockClassRepository(mockCtrl)
	dummyError := errors.New("error for testing")
	mockClassRepo.EXPECT().Create("class A", 20).Return(dummyError).Times(1)
	error := mockClassRepo.Create("class A", 20)
	if error != dummyError {
		t.Fail()
	}
}
