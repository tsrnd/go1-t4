package handlers

import (
	"fmt"
	"net/http"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "need to implement Add...")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "need to implement Update...")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "need to implement Delete...")
}
