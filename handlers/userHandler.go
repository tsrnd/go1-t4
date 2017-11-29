package handlers

import (
	"fmt"
	"net/http"

	"github.com/goweb4/models"
	"github.com/goweb4/utils"
)

func Index(w http.ResponseWriter, r *http.Request) {
	utils.GenerateTemplate(w, "index")
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
func Test(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// userId, _ := strconv.Atoi(vars["id"])
	// user, err := models.GetUser(userId)

	users, err := models.GetUsers()
	if err != nil {
		fmt.Fprintln(w, err)
	} else {
		fmt.Fprintln(w, users)
	}
}
