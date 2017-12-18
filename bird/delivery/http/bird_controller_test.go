package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	mockUc "github.com/goweb4/bird/usecase/mocks"
	"github.com/stretchr/testify/assert"
	models "github.com/goweb4/bird"	
	"bytes"
)

func TestCreateBird(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	var jsonStr = []byte(`{ "name": "ssss", "color": "sss", "description": "sss" }`)
	req, err := http.NewRequest("POST", "/bird", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	rp := httptest.NewRecorder()

	var dataExpect *models.Bird
	mockBirdUc := mockUc.NewMockBirdUsecase(mockCtrl)
	mockBirdUc.EXPECT().Create("ssss", "sss", "sss").Return(dataExpect, nil).Times(1)

	handler := &BirdController{
		Usecase: mockBirdUc,
	}

	handler.Create(rp, req)
	
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, http.StatusCreated, rp.Code)
}