package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"

	mockRepos "github.com/goweb4/gift/repository/mock"
)

func TestGetByFromUserID(t *testing.T) {
	// gifts := []gift.Gift{}
	mockCtrl := gomock.NewController(t)
	mockRepo := mockRepos.NewMockGiftRepository(mockCtrl)
	mockRepo.EXPECT().GetByFromUserID(1).Return(nil, nil)
	_, err := mockRepo.GetByFromUserID(1)
	assert.Equal(t, nil, err)
}
