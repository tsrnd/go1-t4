package models

import (
	"github.com/jinzhu/gorm"
)

type Image struct {
	gorm.Model
	Name 		string		`schema:"name"`
	URL  		string		`schema:"url"`
	ProductId	uint		`schema:"product_id"`
}
