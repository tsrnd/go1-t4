package utils

import (
	"fmt"
	"html/template"
	"net/http"
)

func GenerateTemplate(writer http.ResponseWriter, fn ...string) {
	var files []string
	for _, file := range fn {
		files = append(files, fmt.Sprintf("templates/frontend/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "main", "")
}
 