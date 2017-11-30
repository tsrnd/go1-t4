package handlers

import (
	"fmt"
	"net/http"

	"github.com/goweb4/models"
	"github.com/goweb4/utils"
	"github.com/stretchr/objx"
)

func Login(w http.ResponseWriter, r *http.Request) {
	utils.GenerateTemplate(w, "login_register", "login", "register")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var info models.User
	err := utils.MapFormValues(&info, r)
	if err != nil {
		fmt.Println("cannot decode login info: ", err)
	} else {
		if checkCredential(info) {
			setSession(info.UserName, w)
			http.Redirect(w, r, "/index", http.StatusSeeOther)
		}
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	clearSession(w)
	http.Redirect(w, r, "/login", 302)
}

func checkCredential(info models.User) bool {
	user := models.GetUser(info)
	if user == (models.User{}) {
		return false
	}
	return true
}

func setSession(userName string, response http.ResponseWriter) {
	authCookieValue := objx.New(map[string]interface{}{
		"name": userName,
	}).MustBase64()
	http.SetCookie(response, &http.Cookie{
		Name:  "auth",
		Value: authCookieValue,
		Path:  "/index.html"})

}

func clearSession(response http.ResponseWriter) {
	http.SetCookie(response, &http.Cookie{
		Name:   "auth",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
}
