package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	// "path"
	// "strconv"
	"github.com/go-chi/chi"
	"github.com/goweb4/bird/usecase"
	"github.com/goweb4/services/cache"
)


type BirdController struct {
	Usecase usecase.BirdUsecase
}

func NewBirdController(r chi.Router, uc usecase.BirdUsecase, c cache.Cache) *BirdController {
	handler := &BirdController{
		Usecase: uc,
	}
	r.Get("/bird", handler.Create)
	r.Post("/bird", handler.Create)
	return handler
}

// Create func
func (ctrl *BirdController) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintln(w, "Method not found", http.StatusNotFound)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var cjr CreateBirdRequest
	err := decoder.Decode(&cjr)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	bird, err := ctrl.Usecase.Create(cjr.Name, cjr.Color, cjr.Description)
	if err != nil {
		log.Fatalf("Creating a bird: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, http.StatusCreated, bird)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}