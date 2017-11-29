package handlers

import (
  "fmt"
  "net/http"
  "html/template"
  // "github.com/goweb4/models"
  // "github.com/gorilla/mux"
  // "strconv"
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
	fmt.Fprintln(w, "Need to be implement");
}

/**
  * Show form create new product
  */
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	templates, err := template.ParseFiles(
    "templates/admin/master.html",
    fmt.Sprintf("templates/admin/add_product.html"),		
  )

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  if err := templates.ExecuteTemplate(w, "master", ""); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

/**
  * Admin create new product
  */
func StoreProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Need to be implement");
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
