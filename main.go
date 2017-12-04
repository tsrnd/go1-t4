package main

import (
	"log"
	"net/http"
	"os"

	"github.com/goweb4/config"
)

func main() {
	router := config.NewRouter()
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
