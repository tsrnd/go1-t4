package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/goweb4/handlers"
)

func TestLogin(t *testing.T) {
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/login", nil)
	homeValue := handlers.HomePageVars{}
	homeValue.Name = "test"
	mockData := &HandlerTest{HomePageVars: homeValue, Message: "Huyen"}
	handler := &handlers.Handler{Variable: mockData}
	handler.Login(writer, request)
	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}

type HandlerTest struct {
	HomePageVars handlers.HomePageVars
	Message      string
}

func (h HandlerTest) NewHomePageVars(r *http.Request) handlers.HomePageVars {
	return h.HomePageVars
}

func (h HandlerTest) ShowMessage(w http.ResponseWriter, r *http.Request) string {
	return h.Message
}

func (h HandlerTest) GenerateTemplate(w http.ResponseWriter, home handlers.HomePageVars) {
}
