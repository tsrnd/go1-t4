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

	user := models.User{UID: "uid", UserName: "username", Gender: "gender"}
	// user.UserName = "username1"
	// user.Gender = "gender1"
	// models.User.Save(&user)
	// user := models.User(&user).Update{(UserName: "username1")
	// err := db.Where("id = ?", id).Delete(&User{}).Error
	// id := models.User.id
	// user.UserName = "username1"
	// user.Gender = "gender1"
	// models.User.Save(&user)
	user, err := models.User.UpdateUser({UID: "uid", UserName: "username1", Gender: "gender1"})
	if err != nil {
		fmt.Fprintln(w, err)
	} else {
		fmt.Fprintln(w, user)
	}
}
