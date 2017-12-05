package main

import (
	"log"
	"net/http"
	"os"
	"github.com/goweb4/database"
	"github.com/goweb4/config"
)

func main() {
	database.DBCon, errConnection = database.DBConnection(); if errConnection != nil {
		log.Fatal(errConnection)
	}
	defer database.DBCon.close()
	router := config.NewRouter()
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
