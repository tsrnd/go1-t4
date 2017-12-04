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
	Product		Product		//belong to Product
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
