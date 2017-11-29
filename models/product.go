package models

import (
	"github.com/goweb4/database"
	"log"
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Size    string
	Color   string
	Price   float64
	ImageID string
	InStock int
	GroupID int
}

func GetProducts() (products []Product, err error){
	products = []Product{}
	db, errConnection := database.DBConnection(); if errConnection != nil {
		log.Fatal(errConnection)
	}
	defer db.Close()

	err = db.Find(&products).Error

	return products, err
}

func GetProduct(id uint) (product Product, err error) {
	product = Product{}
	db, errConnection := database.DBConnection(); if errConnection != nil {
		log.Fatal(errConnection)
	}
	defer db.Close()

	err = db.Where("id = ?", id).Find(&product).Error

	return product, err
}

func UpdateProduct(id uint, oldProduct map[string]interface{}) (error) {
	product := Product{}
	db, errConnection := database.DBConnection(); if errConnection != nil {
		log.Fatal(errConnection)
	}
	defer db.Close()

	errFindProduct := db.Where("id = ?", id).Find(&product).Error; if errFindProduct != nil {
		return errFindProduct
	}
	errUpdate := db.Model(&product).Updates(oldProduct).Error;
	return errUpdate
}

func DeleteProduct(id uint) (error) {
	db, errConnection := database.DBConnection(); if errConnection != nil {
		log.Fatal(errConnection)
	}
	defer db.Close()

	err := db.Where("id = ?", id).Delete(&Product{}).Error

	return err
}

func CreateProduct(product Product) (error){
	db, errConnection := database.DBConnection(); if errConnection != nil {
		log.Fatal(errConnection)
	}
	defer db.Close()

	err := db.Create(&product).Error
	
	return err
}
