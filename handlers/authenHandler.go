package handlers

import (
	"fmt"
	"net/http"

	"github.com/goweb4/utils"
)

type LoginInfo struct {
	Name     string `schema:"username"`
	Password string `schema:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	utils.GenerateTemplate(w, "login_register", "login", "register")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	info := new(LoginInfo)

	// r.PostForm is a map of our POST form values
	err2 := utils.MapFormValues(info, r)
	if err2 != nil {
		fmt.Println("cannot decode login info: ", err2)
	} else {
		fmt.Println(info.Name, " && ", info.Password)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {

}
