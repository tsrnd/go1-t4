package utils

import (
	"encoding/base64"
	"fmt"
	"html/template"
	"image"
	"image/jpeg"
	_ "image/png"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gorilla/schema"
	"github.com/nfnt/resize"
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

func HandleImage(file io.Reader, header *multipart.FileHeader, id uint, baseUrl string) (string, string, error) {
	fileName := GetFileName(header.Filename)
	img, err := GetImageFrom(file)
	if err != nil {
		return "", "", err
	}
	img = resize.Resize(300, 300, img, resize.Lanczos3)
	out, err := SaveImage(baseUrl, fileName)
	if err != nil {
		return "", "", err
	}
	defer out.Close()
	err = ToJpeg(out, img)
	filePath := (strings.Join([]string{baseUrl, fileName + ".jpg"}, "/"))
	return fileName, filePath, err
}

func GetFileName(filename string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}

func GetImageFrom(file io.Reader) (image.Image, error) {
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return img, err
}

func SaveImage(baseUrl string, fileName string) (*os.File, error) {
	out, err := os.Create(strings.Join([]string{".", baseUrl, fileName + ".jpg"}, "/"))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return out, err
}

func ToJpeg(file *os.File, img image.Image) (err error) {
	err = jpeg.Encode(file, img, nil)
	if err != nil {
		log.Fatal(err)
	}
	return
}
