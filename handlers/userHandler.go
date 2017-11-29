package handlers

import (
	"fmt"
	"net/http"
	"github.com/goweb4/utils"	
)

func Index(w http.ResponseWriter, r *http.Request) {
	utils.GenerateTemplate(w, "index")
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "need to implement ...")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "need to implement ...")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "need to implement ...")
}
 