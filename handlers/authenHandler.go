package handlers

import (
	"net/http"

	"github.com/goweb4/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	utils.GenerateTemplate(w, "register")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

}

func Logout(w http.ResponseWriter, r *http.Request) {

}
