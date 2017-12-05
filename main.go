package main

import (
	"log"
	"net/http"
	"github.com/goweb4/database"
	"github.com/goweb4/config"
)

func main() {
	var errConnection error
	database.DBCon, errConnection = database.DBConnection(); if errConnection != nil {
		log.Fatal(errConnection)
		panic(errConnection)
	}
	defer database.DBCon.Close()
	router := config.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
