package handlers

import (
  "fmt"
  "net/http"
  "github.com/goweb4/models"
)

/**
  * User crate new order 
  */
func AddOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Need to be implement");
}

/**
  * Show form order's edit 
  */
func EditOrder(w http.ResponseWriter, r *http.Request) {
  models.TestDB()

	fmt.Fprintln(w, "Need to be implements");
}

/**
  * User update order's infor
  */
func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Need to be implement");
}
