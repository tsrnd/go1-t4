package utils

import (
	"fmt"
	"html/template"
	"net/http"
	"github.com/gorilla/schema"
)

func GenerateTemplate(writer http.ResponseWriter, fn ...string) {
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

	if err := templates.ExecuteTemplate(writer, fn[0] + ".html", ""); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func MapFormValues(dst interface{}, r *http.Request) (error){
	errParse := r.ParseForm(); if errParse != nil {
		return errParse
	}
	decoder := schema.NewDecoder()
	err := decoder.Decode(dst, r.PostForm)
	return err
}
