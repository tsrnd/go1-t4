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

func GetProductGroup(id uint) (productGroup ProductGroup, err error) {
	db, errConnection := database.DBConnection(); if errConnection != nil {
		log.Fatal(errConnection)
		return productGroup, errConnection
	}
	defer db.Close()

	err = db.Where("id = ?", id).Find(&productGroup).Error
	return
}

func GetProductsByGroupID(id uint) (products []Product, err error) {
	products = []Product{}
	db, errConnection := database.DBConnection()
	if errConnection != nil {
		log.Fatal(errConnection)
	}
	defer db.Close()

	err = db.Where("group_id = ?", id).Find(&products).Error
	// db.Model(&user).Related(&emails)

	return products, err
}
