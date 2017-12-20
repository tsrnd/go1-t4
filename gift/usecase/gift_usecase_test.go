package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/goweb4/gift"

	"github.com/golang/mock/gomock"
	mockUc "github.com/goweb4/gift/usecase/mock"
)

func TestGetByFromUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	resultExpect := []*gift.Gift{}
	mockUsecase := mockUc.NewMockGiftUsecase(ctrl)
	mockUsecase.EXPECT().GetByFromUserID(int64(1)).Return(resultExpect, nil)
	_, err := mockUsecase.GetByFromUserID(1)
	assert.NoError(t, err)
}
