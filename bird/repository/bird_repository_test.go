package repository

import (
	"testing"
	"github.com/golang/mock/gomock"
	"github.com/goweb4/bird/repository/mocks"
	"github.com/stretchr/testify/assert"
	models "github.com/goweb4/bird"
)

func TestBirdCreateSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockBirdRepository(mockCtrl)
	mockRepo.EXPECT().Create("My bird", "blue", "blue bird").Return(int64(1), nil).Times(1)
	_, err := mockRepo.Create("My bird", "blue", "blue bird")
	assert.Equal(t, nil, err)
}

func TestBirdCreateErrors(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	var expectData *models.Bird

	mockRepo := mocks.NewMockBirdRepository(mockCtrl)
	mockRepo.EXPECT().GetByID(int64(1)).Return(expectData, nil).Times(1)

	rs, err := mockRepo.GetByID(int64(1))
	assert.Equal(t, nil, err)
	assert.NotEqual(t, rs, nil)
}		