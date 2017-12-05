package handlers

import (
  "fmt"
  "net/http"
  "github.com/goweb4/models"
  "github.com/gorilla/mux"
  "github.com/goweb4/utils"
  
  "strconv"
)

/**
  * User create new order 
  */
func AddOrder(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, r.FormValue("product_id"));
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

/**
  * User update order's infor
  */
  func Checkout(w http.ResponseWriter, r *http.Request) {
    if GetAuthName(r) == "" {
      http.Redirect(w, r, "/login", 302)      
    }
    utils.GenerateTemplate(w, NewHomePageVars(r), "checkout")
    // fmt.Fprintln(w, "Need to be implement");
  }
