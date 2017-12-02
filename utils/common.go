package utils

import (
	"encoding/base64"
	"fmt"
	_ "image/png"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/schema"
)

func MapFormValues(dst interface{}, r *http.Request) error {
	contentType := r.Header.Get("Content-type")
	var errParse error
	if strings.Contains(contentType, "multipart/form-data") {
		errParse = r.ParseMultipartForm(30 << 32)
	} else {
		errParse = r.ParseForm()
	}
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
