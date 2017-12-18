package http

import (
	"encoding/json"
	"net/http"

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
	return handler
}
func (g *BookController) Books(w http.ResponseWriter, r *http.Request) {
	name := "vien"
	books, err := g.Usecase.GetByName(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	//response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
