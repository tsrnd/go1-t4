package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/goweb4/gift"

	"github.com/golang/mock/gomock"

	repos "github.com/goweb4/gift/repository/mock"
)

func TestGetByFromUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	resultExpect := []*gift.Gift{}
	mockRepo := repos.NewMockGiftRepository(ctrl)
	mockRepo.EXPECT().GetByFromUserID(int64(1)).Return(resultExpect, nil)
	Usecase := NewGiftUsecase(mockRepo)
	_, err := Usecase.GetByFromUserID(1)
	assert.NoError(t, err)
}
