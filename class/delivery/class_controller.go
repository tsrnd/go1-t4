package delivery

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/goweb4/class/usecase"
	"github.com/goweb4/services/cache"
)

type ClassController struct {
	Usecase usecase.ClassUsecase
	Cache   cache.Cache
}

func NewClassController(r chi.Router, uc usecase.ClassUsecase, cache cache.Cache) *ClassController {
	handler := &ClassController{
		Usecase: uc,
		Cache:   cache,
	}
	r.Get("/getClasses", handler.GetAllClass)
	r.Post("/createClass", handler.CreateClass)
	r.Delete("/deleteClass{id}", handler.DeleteClass)
	return handler
}

func (c *ClassController) GetAllClass(w http.ResponseWriter, r *http.Request) {

}

func (c *ClassController) CreateClass(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var classReq ClassCreationRequest
	err := decoder.Decode(&classReq)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	err = c.Usecase.CreateClass(classReq.Name, classReq.StudentNumber)
	if err != nil {
		fmt.Println("Error occur when creating class: ", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c *ClassController) DeleteClass(w http.ResponseWriter, r *http.Request) {

}
