package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/goweb4/models"
	"github.com/goweb4/utils"
)

type UserHandler struct{}

func UserProfile(w http.ResponseWriter, r *http.Request) {
	// user, err := models.GetUserByUserName(GetAuthName(r))
	// HomeVars := NewHomePageVars(r)
	// HomeVars.User = user
	// if err != nil {
	// 	fmt.Fprintln(w, err)
	// } else {
	// 	utils.GenerateTemplate(w, HomeVars, "user_profile")
	// }
}

// func AddUser(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "need to implement Add...")
// }

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 32)
	user, err := models.GetUserById(uint(id))
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	errMap := utils.MapFormValues(&user, r)
	if errMap != nil {
		fmt.Fprintln(w, errMap)
		return
	}
	errUpdate := models.UpdateUser(&user)
	if errUpdate != nil {
		fmt.Fprintln(w, errUpdate)
		return
	}
	fmt.Fprintln(w, "Update this user success")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 32)
	err := models.DeleteUser(uint(id))
	if err != nil {
		fmt.Fprintln(w, err)
		return
	} else {
		fmt.Fprintln(w, "Delete this user success")
	}
}
