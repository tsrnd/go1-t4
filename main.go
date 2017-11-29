package main

import (
	"log"
	"net/http"

	"github.com/goweb4/config"
)

func main() {
	router := config.NewRouter()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Fatal(http.ListenAndServe(":8080", router))
}
