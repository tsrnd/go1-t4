package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	uc "github.com/goweb4/gift/usecase"
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
func (g *BookController) Books(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("token")
	userIDStr, err := g.Cache.Get(fmt.Sprintf("token_%s", token))
	if err != nil {
		http.Error(w, "Authentication failed", http.StatusBadRequest)
		return
	}
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Cannot processing user id", http.StatusBadRequest)
		return
	}
	books, err := g.Usecase.GetByFromUserID(int64(userID))
	if err != nil {
		http.Error(w, "Has error during get book", http.StatusForbidden)
		return
	}
	//response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// package http

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"path"
// 	"strconv"

// 	"github.com/go-chi/chi"
// 	"github.com/goweb4/book/usecase"
// 	"github.com/goweb4/services/cache"
// )

// // BookController type
// type BookController struct {
// 	Usecase usecase.BookUsecase
// 	Cache   cache.Cache
// }

// // NewBookController func
// func NewBookController(r chi.Router, uc usecase.BookUsecase, c cache.Cache) *BookController {
// 	handler := &BookController{
// 		Usecase: uc,
// 		Cache:   c,
// 	}
// 	r.Post("/books", handler.Book)
// 	r.Get("/books", handler.Book)
// 	r.Put("/books", handler.Book)
// 	r.Post("/books/create", handler.Create)
// 	return handler
// }

// // Create func
// func (ctrl *BookController) Create(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "POST" {
// 		http.Error(w, "Not found", http.StatusNotFound)
// 		return
// 	}
// 	token := r.Header.Get("token")
// 	userIDStr, err := ctrl.Cache.Get(fmt.Sprintf("token_%s", token))
// 	if err != nil {
// 		http.Error(w, "Invalid token", http.StatusForbidden)
// 		return
// 	}
// 	userID, err := strconv.Atoi(userIDStr)
// 	if err != nil {
// 		log.Fatalf("Convert user id to int: %s", err)
// 		http.Error(w, "", http.StatusInternalServerError)
// 		return
// 	}
// 	decoder := json.NewDecoder(r.Body)
// 	var cjr CreateBookRequest
// 	err = decoder.Decode(&cjr)
// 	if err != nil {
// 		http.Error(w, "Invalid request body", http.StatusBadRequest)
// 		return
// 	}
// 	_, err = ctrl.Usecase.Create(cjr.Name, cjr.Content, int64(userID))
// 	if err != nil {
// 		log.Fatalf("Creating a book: %s", err)
// 		http.Error(w, "", http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteHeader(http.StatusCreated)
// }

// // Book func
// func (ctrl *BookController) Book(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "POST" {
// 		http.Error(w, "Not Found", http.StatusNotFound)
// 		return
// 	}
// 	bookID, err := strconv.Atoi(path.Base(r.URL.Path))
// 	if err != nil {
// 		http.Error(w, "Not Found", http.StatusNotFound)
// 		return
// 	}
// 	book, err := ctrl.Usecase.GetByID(int64(bookID))
// 	if err != nil {
// 		http.Error(w, "Not Found", http.StatusNotFound)
// 		return
// 	}
// 	if r.Method == "GET" {
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(book)
// 	}
// 	token := r.Header.Get("token")
// 	userIDStr, err := ctrl.Cache.Get(fmt.Sprintf("token_%s", token))
// 	if err != nil {
// 		http.Error(w, "Invalid token", http.StatusForbidden)
// 		return
// 	}
// 	userID, err := strconv.Atoi(userIDStr)
// 	if err != nil {
// 		log.Fatalf("Convert user id to int: %s", err)
// 		http.Error(w, "", http.StatusInternalServerError)
// 		return
// 	}
// 	if int64(userID) != book.UserID {
// 		http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 		return
// 	}
// 	if r.Method == "PUT" {
// 		decoder := json.NewDecoder(r.Body)
// 		var ujr UpdateBookRequest
// 		err = decoder.Decode(&ujr)
// 		if err != nil {
// 			http.Error(w, "Invalid request body", http.StatusBadRequest)
// 			return
// 		}
// 		_, err = ctrl.Usecase.Update(book.ID, ujr.Name, ujr.Content)
// 		if err != nil {
// 			log.Fatalf("Updating a book: %s", err)
// 			http.Error(w, "", http.StatusInternalServerError)
// 			return
// 		}
// 	}
// 	if r.Method == "DELETE" {
// 		err = ctrl.Usecase.Delete(int64(book.ID))
// 		if err != nil {
// 			http.Error(w, "", http.StatusInternalServerError)
// 			return
// 		}
// 	}
// }

// // Feed func
// func (ctrl *BookController) Feed(w http.ResponseWriter, r *http.Request) {
// 	var err error
// 	if r.Method != "GET" {
// 		http.Error(w, "Not found", http.StatusNotFound)
// 		return
// 	}
// 	offset := 0
// 	offsetStr, ok := r.URL.Query()["offset"]
// 	if ok {
// 		offset, err = strconv.Atoi(offsetStr[0])
// 		if err != nil {
// 			offset = 0
// 		}
// 	}

// 	limit := 10
// 	limitStr, ok := r.URL.Query()["limit"]
// 	if ok {
// 		limit, err = strconv.Atoi(limitStr[0])
// 		if err != nil {
// 			limit = 1
// 		}
// 	}
// 	books, err := ctrl.Usecase.Fetch(int64(offset), int64(limit))
// 	w.Header().Set("Content-Type", "application/json")
// 	if err != nil {
// 		http.Error(w, "", http.StatusInternalServerError)
// 		return
// 	}
// 	json.NewEncoder(w).Encode(books)
// }
