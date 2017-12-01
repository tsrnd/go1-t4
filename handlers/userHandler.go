package handlers

import (
	"fmt"
	"net/http"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "need to implement ...")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "need to implement ...")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "need to implement ...")
}

func Register(w http.ResponseWriter, r *http.Request) {
	// utils.MapFormValues(,r)
}
