package models

import (
	"log"

	"github.com/goweb4/database"
)

type Model interface {
	GetRelationship() map[string]interface{}
}

func WithConnectionDB(fn func(db *database.DB)) {
	db, errConnection := database.DBConnection()
	if errConnection != nil {
		log.Fatal(errConnection)
	}
	defer db.Close()
	fn(db)
}

func GetRelatedData(model Model, relationshipName string) {
	data := model.GetRelationship()
	WithConnectionDB(func(db *database.DB) {
		db.Model(model).Association(relationshipName).Find(data[relationshipName])
	})
}
