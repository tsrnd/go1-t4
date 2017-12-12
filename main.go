package main

import (
	"log"
	"net/http"

	"github.com/goweb4/config"
	"github.com/goweb4/database"
)

func main() {
	var errConnection error
	database.DBCon.Db, errConnection = database.DBConnection()
	if errConnection != nil {
		log.Fatal(errConnection)
		panic(errConnection)
	}
	defer database.DBCon.Db.Close()
	router := config.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
