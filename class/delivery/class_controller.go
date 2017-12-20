package delivery

import (
	"encoding/json"
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
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	err = c.Usecase.CreateClass(classReq.Name, classReq.StudentNumber)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

func (c *ClassController) DeleteClass(w http.ResponseWriter, r *http.Request) {

}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (c *ClassController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	decode := json.NewDecoder(r.Body)
	var loginRequest ClassLoginRequest
	err := decode.Decode(&loginRequest)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

}
