package utils

import (
	"encoding/base64"
	"fmt"
	_ "image/png"
	"net/http"
	"strings"
	"time"
	"strconv"
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

type Paginator struct {
	Start int
	End int
	CurrentPage int
	PerPage int
	NextPage int
	PrevPage int
	PagePath []int
}
func Paginate(totalRecord int, perPage int, currentPage int) Paginator {
	var paginator Paginator

	if currentPage == 0 {
		currentPage++
	}

	TotalPage := totalRecord/perPage + 1
	paginator.Start = (currentPage - 1)*perPage
	paginator.End = paginator.Start + perPage
	if paginator.End > totalRecord {
		paginator.End = paginator.Start + totalRecord - paginator.Start
	}
	for i := 1; i <= TotalPage; i++ {
		paginator.PagePath = append(paginator.PagePath, i)
	}

	if currentPage >= TotalPage {
		paginator.NextPage = TotalPage
	} else {
		paginator.NextPage =  currentPage + 1
	}

	if currentPage <= 0 {
		paginator.PrevPage = 1
	} else { 
		paginator.PrevPage = currentPage - 1
	}
	paginator.PerPage = perPage
	paginator.CurrentPage = currentPage

	return paginator
}

func ConvertStrToUint(value string) (uint, error) {
	var result uint64
	result, err := strconv.ParseUint(value, 10, 32)
	return uint(result), err
}
