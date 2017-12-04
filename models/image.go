package models

import (
	"github.com/goweb4/database"
	"github.com/jinzhu/gorm"
)

type Image struct {
	gorm.Model
	Name 		string	`schema:"name"`
	URL  		string	`schema:"url"`
	ProductId	uint	`schema:"product_id"`
	Product		Product	//belong to Product
}

const IMG_BASE_URL = "uploads/images"

func StoreImage(image *Image) (errCreateImage error) {
	WithConnectionDB(func(db *database.DB) {
		errCreateImage = db.Create(&image).Error
	})
	return errCreateImage
}
