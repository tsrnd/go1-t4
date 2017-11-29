package models

import "github.com/jinzhu/gorm"

type Payment struct {
	gorm.Model
	ID     int
	Method string
}
