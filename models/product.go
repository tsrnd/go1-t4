package models

import (
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
