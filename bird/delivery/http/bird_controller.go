package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"
	"github.com/go-chi/chi"
	"github.com/goweb4/bird/usecase"
	"github.com/goweb4/services/cache"
)


type BirdController struct {
	Usecase usecase.BirdUsecase
	Cache   cache.Cache
}

func NewBirdController(r chi.Router, uc usecase.BirdUsecase, c cache.Cache) *BirdController {
	handler := &BirdController{
		Usecase: uc,
		Cache:   c,
	}
	r.Get("/bird", handler.Product)
	return handler
}

// Create func
func (ctrl *BirdController) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bird is being implement")
}