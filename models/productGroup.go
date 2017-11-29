package models

import (
	"github.com/jinzhu/gorm"
)

type ProductGroup struct {
	gorm.Model	
	Name    string		`schema:"name"`
	ImageID string		`schema:"image_id"`
}
