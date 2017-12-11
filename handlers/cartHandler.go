package handlers

import (
	"net/http"

	"github.com/goweb4/utils"
)

/**
 * User create new order
 */
func ShowCart(w http.ResponseWriter, r *http.Request) {
	HomeVars := NewHomePageVars(r)
	utils.GenerateTemplate(w, HomeVars, "cart")
}
