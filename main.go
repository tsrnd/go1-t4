package main

import (
	"log"
	"net/http"

	"github.com/goweb4/config"
)

func main() {
	router := config.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
