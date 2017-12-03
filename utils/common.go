package utils

import (
	"fmt"
	"html/template"
	"net/http"
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