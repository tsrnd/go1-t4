package models

import (
	"log"
	"github.com/jinzhu/gorm"
	"github.com/goweb4/database"
)

type ProductGroup struct {
	gorm.Model	
	Name    string		`schema:"name"`
}

func GetProductGroups() (productGroups []ProductGroup, err error) {
	productGroups = []ProductGroup{}
	db, errConnection := database.DBConnection(); if errConnection != nil {
		log.Fatal(errConnection)
		return productGroups, errConnection
	}
	defer db.Close()

	err = db.Find(&productGroups).Error

	return productGroups, err
}
