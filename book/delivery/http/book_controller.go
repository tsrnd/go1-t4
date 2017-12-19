package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	uc "github.com/goweb4/book/usecase"
	cache "github.com/goweb4/services/cache"
)

type BookController struct {
	Usecase uc.BookUsecase
	Cache   cache.Cache
}

func NewBookController(r chi.Router, uc uc.BookUsecase, c cache.Cache) *BookController {
	handler := &BookController{
		Usecase: uc,
		Cache:   c,
	}
	r.Get("/books", handler.Books)
	r.Post("/books", handler.Create)
	return handler
}
func (g *BookController) Books(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("token")
	userIDStr, err := g.Cache.Get(fmt.Sprintf("token_%s", token))
	if err != nil {
		http.Error(w, "Authentication failed", http.StatusBadRequest)
		return
	}
	_, err = strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Cannot processing user id", http.StatusBadRequest)
		return
	}
	nameStr := "vien"
	books, err := g.Usecase.GetByName(string(nameStr))
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func (gC *BookController) Create(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi("")
	if err != nil {
		http.Error(w, "Cannot processing user id", http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var cjr CreateBookRequest
	err = decoder.Decode(&cjr)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	_, err = gC.Usecase.Create(cjr.Name, cjr.Content, int64(userID))
	if err != nil {
		log.Fatalf("Creating a Book: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
