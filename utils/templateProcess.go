package utils

import (
	"fmt"
	"html/template"
	"net/http"
)

func GenerateTemplate(writer http.ResponseWriter, i interface{}, fn ...string) {
	var files []string

	fn = append(fn, "header")
	fn = append(fn, "footer")

	for _, file := range fn {
		files = append(files, fmt.Sprintf("templates/frontend/%s.html", file))
	}
	templates, err := template.ParseFiles(files...)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := templates.ExecuteTemplate(writer, fn[0]+".html", i); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func GenerateTemplateAdmin(writer http.ResponseWriter, i interface{}, fn ...string) {
	var files []string

	fn = append(fn, "header")
	fn = append(fn, "footer")
	fn = append(fn, "aside")

	for _, file := range fn {
		files = append(files, fmt.Sprintf("templates/admin/%s.html", file))
	}
	templates, err := template.ParseFiles(files...)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := templates.ExecuteTemplate(writer, fn[0]+".html", i); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
