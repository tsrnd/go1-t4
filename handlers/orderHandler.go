package handlers

import (
  "fmt"
  "net/http"
  "github.com/goweb4/models"
  "github.com/gorilla/mux"
  "strconv"
)

/**
  * User create new order 
  */
func AddOrder(w http.ResponseWriter, r *http.Request) {
  fmt.Println("testt")
	fmt.Fprintln(w, "Need to be implement");
}

/**
  * Show form order's edit 
  */
func EditOrder(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  orderId, _ := strconv.Atoi(vars["id"]);
  order, err := models.GetOrder(orderId); if err != nil {
    fmt.Fprintln(w, err);
  }
  
	fmt.Fprintln(w, order);
}

/**
  * User update order's infor
  */
func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Need to be implement");
}
