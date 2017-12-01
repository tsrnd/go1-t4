package utils

import (
	"encoding/base64"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/schema"
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

func MapFormValues(dst interface{}, r *http.Request) error {
	errParse := r.ParseForm()
	if errParse != nil {
		return errParse
	}
	decoder := schema.NewDecoder()
	err := decoder.Decode(dst, r.PostForm)
	return err
}

func ShowMessage(w http.ResponseWriter, r *http.Request, name string) (message string) {
	c, err := r.Cookie(name)
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Println(w, "No message found")
		}
	} else {
		rc := http.Cookie{
			Name:    name,
			MaxAge:  -1,
			Expires: time.Unix(1, 0),
		}
		http.SetCookie(w, &rc)
		val, _ := base64.URLEncoding.DecodeString(c.Value)
		message = string(val)
	}
	return message
}

func SetMessage(w http.ResponseWriter, message string, name string) {
	msg := []byte(message)
	c := http.Cookie{
		Name:  name,
		Value: base64.URLEncoding.EncodeToString(msg)}
	http.SetCookie(w, &c)
}
