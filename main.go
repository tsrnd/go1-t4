package main

import (
	"log"
	"net/http"
	"os"

	"github.com/goweb4/config"
	"github.com/goweb4/database"
)

func main() {
	var errConnection error
	database.DBCon, errConnection = database.DBConnection()
	if errConnection != nil {
		log.Fatal(errConnection)
		panic(errConnection)
	}
	defer database.DBCon.Close()
	router := config.NewRouter()
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
