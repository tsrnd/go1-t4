package handlers

import (
	"log"
  "fmt"
  "net/http"
  "github.com/goweb4/models"
  "github.com/gorilla/mux"
  "strconv"
  "github.com/goweb4/utils"
)

/**
  * Index Product
  */
func IndexProduct(w http.ResponseWriter) {
	fmt.Fprintln(w, "Need to be implement")
}

/**
  * Show product
  */
func ShowProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
  id, _ := strconv.ParseUint(vars["id"], 10, 32);
  product, err := models.GetProduct(uint(id)); if err != nil {
    fmt.Fprintln(w, err);
  }
  fmt.Fprintln(w, product)
}

/**
  * Show form create new product
  */
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Need to be implement");
}

/**
  * Admin create new product
  */
func StoreProduct(w http.ResponseWriter, r *http.Request) {
  product := new(models.Product)
  err := r.ParseForm(); if err != nil {
    log.Fatal(err)
  }
  
  errMap := utils.MapFormValues(product, r); if errMap != nil {
    log.Fatal(errMap)
  }
  
  fmt.Println("Product: ", product)
}

/**
  * Show form product's edit 
  */
func EditProduct(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Need to be implement")
}

/**
  * User update product's infor
  */
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Need to be implement");
}

/**
  * Delete product
  */
func DestroyProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Need to be implement");
}