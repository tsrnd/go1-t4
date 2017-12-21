package http

import (
	"encoding/json"
	"log"
	"net/http"
	"path"
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
	r.Get("/books/{name}", handler.Books)
	r.Post("/books", handler.Create)
	return handler
}
func (g *BookController) Books(w http.ResponseWriter, r *http.Request) {
	name := path.Base(r.URL.Path)
	books, err := g.Usecase.GetByName(string(name))
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
