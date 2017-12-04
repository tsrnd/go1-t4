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

func (image *Image) GetRelationship() map[string]interface{}{
	relationship := map[string]interface{} {
		"Product": &image.Product,
	}
	return relationship
}

const IMG_BASE_URL = "uploads/images"

func StoreImage(image *Image) (errCreateImage error) {
	WithConnectionDB(func(db *database.DB) {
		errCreateImage = db.Create(&image).Error
	})
	return errCreateImage
}

func GetImageByProductId(productId uint) (image Image, err error) {
	WithConnectionDB(func(db *database.DB) {
		err = db.Where("product_id = ?", productId).First(&image).Error
	})
	return
}
