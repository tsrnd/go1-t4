package models

import (
	"log"

	"github.com/goweb4/database"
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Size    		string  		`schema:"size"`
	Color   		string  		`schema:"color"`
	Price   		float64 		`schema:"price"`
	Name    		string  		`schema:"name"`
	InStock 		uint    		`schema:"in_stock"`
	GroupID 		uint    		`schema:"group_id"`
	ProductGroup	ProductGroup	//belong To Product Group
	OrderProducts	[]OrderProduct	//has many order products
	Images			[]Image			//has many image
}

func GetProducts() (products []Product, err error) {
	products = []Product{}
	db, errConnection := database.DBConnection()
	if errConnection != nil {
		log.Fatal(errConnection)
	}
	defer db.Close()

	err = db.Find(&products).Error

	return products, err
}

func GetProduct(id uint) (product Product, err error) {
	product = Product{}
	db, errConnection := database.DBConnection()
	if errConnection != nil {
		log.Fatal(errConnection)
	}
	defer db.Close()

	err = db.Where("id = ?", id).Find(&product).Error

	return product, err
}

func UpdateProduct(product *Product) error {
	db, errConnection := database.DBConnection()
	if errConnection != nil {
		log.Fatal(errConnection)
	}
	defer db.Close()
	errUpdate := db.Save(&product).Error
	return errUpdate
}

func DeleteProduct(id uint) error {
	product := Product{}
	db, errConnection := database.DBConnection()
	if errConnection != nil {
		log.Fatal(errConnection)
	}
	defer db.Close()

	product, errGet := GetProduct(id)
	if errGet != nil {
		return errGet
	}
	err := db.Delete(&product).Error
	return err
}

func CreateProduct(product *Product) (uint, error) {
	db, errConnection := database.DBConnection()
	if errConnection != nil {
		log.Fatal(errConnection)
	}
	defer db.Close()
	newProduct := db.Create(&product)
	err := newProduct.Error
	if err == nil {
		return product.ID, err
	}
	return 0, err
}
