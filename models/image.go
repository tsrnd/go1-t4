package models

import "github.com/goweb4/database"
import "time"

type Image struct {
	Model
	Name      string   `schema:"name"`
	URL       string   `schema:"url"`
	ProductId uint     `schema:"product_id"`
	Product   *Product //belong to Product
}

const IMG_BASE_URL = "uploads/images"

func StoreImage(image *Image) (errCreateImage error) {
	errCreateImage = database.DBCon.Db.
		QueryRow("INSERT INTO images (name, created_at) VALUES($1,$2) returning id;", image.Name, time.Now()).Scan(&image.ID)
	return
}
