package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"

	"github.com/goweb4/book"
	mockRepos "github.com/goweb4/book/repository/mocks"
)

func TestGetByName(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := mockRepos.NewMockBookRepository(mockCtrl)
	resultExpect := []*book.Book{}
	mockRepo.EXPECT().GetByName("name").Return(resultExpect, nil)
	result, err := mockRepo.GetByName("name")
	assert.NoError(t, err)
	assert.Equal(t, result, resultExpect)

}
