package handlers

import (
	"fmt"
	"net/http"

	"github.com/goweb4/models"
	"github.com/goweb4/utils"
)

func UserProfile(w http.ResponseWriter, r *http.Request) {
	user, err := models.GetUserByUserName(GetAuthName(r))
	HomeVars := NewHomePageVars(r)
	HomeVars.User = user
	if err != nil {
		fmt.Fprintln(w, err)
	}
	utils.GenerateTemplate(w, HomeVars, "user_profile")
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "need to implement Add...")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "need to implement Update...")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "need to implement Delete...")
}
