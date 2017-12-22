package http

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	_ "strings"
	"testing"

	_ "github.com/go-chi/chi"

	"github.com/goweb4/gift"

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

func TestCreate(t *testing.T) {
	giftJSon := []byte(`{"to_user_id":1, "product_id":12, "message":"test create"}`)
	request, err := http.NewRequest("POST", "/gifts", bytes.NewBuffer(giftJSon))
	if err != nil {
		t.Error("error while sending request", err) //Something is wrong while sending request
	}
	response := httptest.NewRecorder()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUsecase := mockUc.NewMockGiftUsecase(mockCtrl)
	mockUsecase.EXPECT().Create(int64(1), int64(1), int64(12), "test create").Return(nil, nil).Times(1)

	mockCac := mockCache.NewMockCache(mockCtrl)
	mockCac.EXPECT().Get("token_").Return("1", nil).Times(1)

	handler := &GiftController{
		Usecase: mockUsecase,
		Cache:   mockCac,
	}

	handler.Create(response, request)
	assert.Equal(t, http.StatusCreated, response.Code)
}

func TestGetByID(t *testing.T) {
	request, err := http.NewRequest("GET", "/gifts/2", nil)
	assert.NoError(t, err)
	response := httptest.NewRecorder()

	mockCtrl := gomock.NewController(t)

	mockUsecase := mockUc.NewMockGiftUsecase(mockCtrl)
	mockUsecase.EXPECT().GetByID(int64(2)).Return(&gift.Gift{}, nil).Times(1)

	mockCac := mockCache.NewMockCache(mockCtrl)
	mockCac.EXPECT().Get("token_").Return("1", nil).Times(1)

	handler := &GiftController{
		Usecase: mockUsecase,
		Cache:   mockCac,
	}

	handler.Show(response, request)
	assert.Equal(t, http.StatusOK, response.Code)
}
