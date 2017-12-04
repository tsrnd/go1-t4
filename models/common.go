package models

import (
	"log"

	"github.com/goweb4/database"
)

func WithConnectionDB(fn func(db *database.DB)) {
	db, errConnection := database.DBConnection()
	if errConnection != nil {
		log.Fatal(errConnection)
	}
	defer db.Close()
	fn(db)
}
