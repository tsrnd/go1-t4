package http

import (
	"net/http"
	"net/http/httptest"
	_ "strings"
	"testing"

	"github.com/golang/mock/gomock"

	mockUc "github.com/goweb4/book/usecase/mock"
	mockCache "github.com/goweb4/services/cache/mock"
	"github.com/stretchr/testify/assert"
)

func TestBooks(t *testing.T) {
	request, err := http.NewRequest("GET", "/books", nil)
	assert.NoError(t, err)
	response := httptest.NewRecorder()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUsecase := mockUc.NewMockBookUsecase(mockCtrl)
	mockUsecase.EXPECT().GetByName(name).Return(nil, nil)

	mockCac := mockCache.NewMockCache(mockCtrl)
	mockCac.EXPECT().Get("token_").Return("1", nil)

	handler := &BookController{
		Usecase: mockUsecase,
		Cache:   mockCac,
	}

	handler.Books(response, request)
	assert.Equal(t, http.StatusOK, response.Code)
}
