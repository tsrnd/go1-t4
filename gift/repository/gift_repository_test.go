package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"

	"github.com/goweb4/gift"
	mockRepos "github.com/goweb4/gift/repository/mock"
)

func TestGetByFromUserID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := mockRepos.NewMockGiftRepository(mockCtrl)
	resultExpect := []*gift.Gift{}
	mockRepo.EXPECT().GetByFromUserID(int64(1)).Return(resultExpect, nil)
	_, err := mockRepo.GetByFromUserID(1)
	assert.NoError(t, err)
}
