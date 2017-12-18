package http

import (
	"net/http"
	"net/http/httptest"
	_ "strings"
	"testing"

	"github.com/golang/mock/gomock"

	mockUc "github.com/goweb4/gift/usecase/mock"
	mockCache "github.com/goweb4/services/cache/mock"
	"github.com/stretchr/testify/assert"
)

func TestGifts(t *testing.T) {
	request, err := http.NewRequest("GET", "/gifts", nil)
	assert.NoError(t, err)
	response := httptest.NewRecorder()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUsecase := mockUc.NewMockGiftUsecase(mockCtrl)
	mockUsecase.EXPECT().GetByFromUserID(int64(1)).Return(nil, nil)

	mockCac := mockCache.NewMockCache(mockCtrl)
	mockCac.EXPECT().Get("token_").Return("1", nil)

	handler := &GiftController{
		Usecase: mockUsecase,
		Cache:   mockCac,
	}

	handler.Gifts(response, request)
	assert.Equal(t, http.StatusOK, response.Code)
}
