package middleware

import (
	"github.com/goweb4/utils"
	"github.com/goweb4/handlers"
	"net/http"

)

func FilterAdmin(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := handlers.GetAuthName(r)
		if username == "" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		if utils.IsAdminRole(username) {
			h(w, r)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
