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

type GiftController struct {
	Usecase uc.GiftUsecase
	Cache   cache.Cache
}

func NewGiftController(r chi.Router, uc uc.GiftUsecase, c cache.Cache) *GiftController {
	handler := &GiftController{
		Usecase: uc,
		Cache:   c,
	}
	r.Get("/gifts", handler.Gifts)
	return handler
}

func (g *GiftController) Gifts(w http.ResponseWriter, r *http.Request) {
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
	gifts, err := g.Usecase.GetByFromUserID(int64(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	//response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gifts)
}
