package usecase

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/goweb4/class/usecase/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreateClass(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockClassUc := mocks.NewMockClassUsecase(mockCtrl)
	mockClassUc.EXPECT().CreateClass("class A", 20).Return(nil).Times(1)
	r := mockClassUc.CreateClass("class A", 20)
	assert.Equal(t, nil, r)
}

func TestCreateClassReturnError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockClassUc := mocks.NewMockClassUsecase(mockCtrl)

	dummyError := errors.New("error for testing")
	mockClassUc.EXPECT().CreateClass("class A", 0).Return(dummyError).Times(1)

	r := mockClassUc.CreateClass("class A", 0)
	if r != dummyError {
		t.Fail()
	}
}
