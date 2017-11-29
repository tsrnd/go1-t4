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
  * User create new order 
  */
func ListProduct(w http.ResponseWriter, r *http.Request) {
	GenerateHTML(w, "index")
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	GenerateHTML(w, "add_product")
}


func GenerateHTML(writer http.ResponseWriter, file string) {
  
    templates, err := template.ParseFiles(
      "templates/admin/master.html",
      fmt.Sprintf("templates/admin/%s.html", file),		
    )
  
    if err != nil {
      http.Error(writer, err.Error(), http.StatusInternalServerError)
      return
    }
  
    if err := templates.ExecuteTemplate(writer, "master", ""); err != nil {
      http.Error(writer, err.Error(), http.StatusInternalServerError)
    }
  }
