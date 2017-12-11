package models

type Image struct {
	Model
	Name      string  `schema:"name"`
	URL       string  `schema:"url"`
	ProductId uint    `schema:"product_id"`
	Product   *Product //belong to Product
}

const IMG_BASE_URL = "uploads/images"

func StoreImage(image *Image) (errCreateImage error) {
	// errCreateImage = database.DBCon.Create(&image).Error
	// return errCreateImage
	return nil
}
