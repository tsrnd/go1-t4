package delivery

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/goweb4/class"
)

var (
	reader io.Reader //Ignore this for now
)

func TestCreateUser(t *testing.T) {
	classJS := []byte(`{"name":"class A", "student_number":30}`)
	request, err := http.NewRequest("POST", "/createClass", bytes.NewBuffer(classJS)) //Create request with JSON body
	if err != nil {
		t.Error("error while sending request", err) //Something is wrong while sending request
	}
	response := httptest.NewRecorder()

	handler := &ClassController{
		Usecase: gomock.Any(),
		Cache:   gomock.Any(),
	}
	handler.CreateClass(response, request)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

type ClassUsecaseMock struct {
}

func (c classUsecase) CreateClass(name string, number int) error {
	err := c.Repo.Create(name, number)
	return err
}

func (c classUsecase) GetClassById(id int) (cl class.Class, err error) {
	return cl, err
}

func (c classUsecase) UpdateClass() error {
	return nil
}

func (c classUsecase) DeleteClass(int) error {
	return nil
}
