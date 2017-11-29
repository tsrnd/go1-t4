package utils

import (
	"fmt"
	"html/template"
	"net/http"
	"github.com/gorilla/schema"
)

func GenerateTemplate(writer http.ResponseWriter, file string) {

	templates, err := template.ParseFiles(
		"templates/frontend/header.html",
		"templates/frontend/footer.html",
		fmt.Sprintf("templates/frontend/%s.html", file),		
	)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := templates.ExecuteTemplate(writer, file + ".html", ""); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func MapFormValues(dst interface{}, r *http.Request) (error){
	decoder := schema.NewDecoder()
	errMapForm := decoder.Decode(dst, r.PostForm);
	return errMapForm
}
 