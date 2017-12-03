package models

import (
	"github.com/jinzhu/gorm"
	"log"
	"github.com/goweb4/database"
)

type Image struct {
	gorm.Model
	Name 		string		`schema:"name"`
	URL  		string		`schema:"url"`
	ProductId	uint		`schema:"product_id"`
}

const IMG_BASE_URL = "uploads/images"

func StoreImage(image *Image) (error){
	//connect to db
	db, errConnection := database.DBConnection()
	if errConnection != nil {
		log.Fatal(errConnection)
		return errConnection
	}
	defer db.Close()
	errCreateImage := db.Create(&image).Error

	return errCreateImage
}

func GetImageByProductId(productId uint) (image Image, err error) {
	db, errConnection := database.DBConnection()
	if errConnection != nil {
		log.Fatal(errConnection)
		return image, errConnection
	}
	defer db.Close()
	err = db.Where("product_id = ?", productId).First(&image).Error
	return
}
