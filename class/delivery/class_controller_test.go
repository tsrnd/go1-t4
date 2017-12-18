package delivery

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	usecaseMock "github.com/goweb4/class/usecase/mocks"
	cacheMock "github.com/goweb4/services/cache/mocks"
)

func TestCreateUser(t *testing.T) {
	classJS := []byte(`{"name":"class A", "student_number":20}`)
	request, err := http.NewRequest("POST", "/createClass", bytes.NewBuffer(classJS)) //Create request with JSON body
	if err != nil {
		t.Error("error while sending request", err) //Something is wrong while sending request
	}
	response := httptest.NewRecorder()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockClassUc := usecaseMock.NewMockClassUsecase(mockCtrl)
	mockClassUc.EXPECT().CreateClass("class A", 20).Return(nil).Times(1)

	mockCache := cacheMock.NewMockCache(mockCtrl)
	handler := &ClassController{
		Usecase: mockClassUc,
		Cache:   mockCache,
	}
	handler.CreateClass(response, request)
	checkResponseCode(t, http.StatusCreated, response.Code)
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
