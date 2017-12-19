package usecase

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/goweb4/book"
	mockUc "github.com/goweb4/book/usecase/mock"
	"github.com/stretchr/testify/assert"
)

func TestGetByFromUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	resultExpect := []*book.Book{}
	mockUsecase := mockUc.NewMockBookUsecase(ctrl)
	mockUsecase.EXPECT().GetByName("name").Return(resultExpect, nil)
	_, err := mockUsecase.GetByName("name")
	assert.NoError(t, err)
}
