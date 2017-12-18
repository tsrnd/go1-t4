package repository

import (
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/goweb4/book"
	mockRepos "github.com/goweb4/book/repository/mock"
)

func TestGetByName(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := mockRepos.NewMockBookRepository(mockCtrl)
	resultExpect := []*book.Book{}
	mockRepo.EXPECT().GetByName(name).Return(resultExpect, nil)
	result, err := mockRepo.GetByName("vien")
	assert.NoError(t, err)
	assert.Equal(t, result, resultExpect)
}
