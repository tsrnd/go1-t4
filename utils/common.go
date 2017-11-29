package utils

import (
	"fmt"
	"html/template"
	"net/http"
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
 