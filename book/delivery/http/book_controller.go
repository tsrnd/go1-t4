package http

import (
	"encoding/json"
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
	r.Get("/books", handler.Create)
	return handler
}
func (g *BookController) Create(w http.ResponseWriter, r *http.Request) {
	// token := r.Header.Get("token")
	// userIDStr, err := g.Cache.Get(fmt.Sprintf("token_%s", token))
	// if err != nil {
	// 	http.Error(w, "Authentication failed", http.StatusBadRequest)
	// 	return
	// }
	bookID, err := strconv.Atoi("")
	// if err != nil {
	// 	http.Error(w, "Cannot processing user id", http.StatusBadRequest)
	// 	return
	// }
	books, err := g.Usecase.GetByID(int64(bookID))
	if err != nil {
		http.Error(w, "Has error during get book", http.StatusForbidden)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
