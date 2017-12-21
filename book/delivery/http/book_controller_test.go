package http

import (
	"net/http"
	"net/http/httptest"
	_ "strings"
	"testing"

	"github.com/golang/mock/gomock"

	mockUc "github.com/goweb4/book/usecase/mocks"
	"github.com/stretchr/testify/assert"
)

func TestBooks(t *testing.T) {
	request, err := http.NewRequest("GET", "/books/test", nil)
	assert.NoError(t, err)
	response := httptest.NewRecorder()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsecase := mockUc.NewMockBookUsecase(mockCtrl)
	mockUsecase.EXPECT().GetByName("test").Return(nil, nil)

	handler := &BookController{
		Usecase: mockUsecase,
	}

	handler.Books(response, request)
	assert.Equal(t, http.StatusOK, response.Code)
}
