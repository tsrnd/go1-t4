package models

import (
	"github.com/jinzhu/gorm"
)

type OrderProduct struct {
	gorm.Model
	OrderID   uint	`schema:"order_id"`
	ProductID uint	`schema:"product_id"`
	Quatity   uint	`schema:"quatity"`
}
