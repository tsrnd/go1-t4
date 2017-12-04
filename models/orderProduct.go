package models

import (
	"github.com/jinzhu/gorm"
)

type OrderProduct struct {
	gorm.Model
	OrderID   uint		`schema:"order_id"`
	ProductID uint		`schema:"product_id"`
	Quatity   uint		`schema:"quatity"`
	Product	  Product	//belong to product
	Order	  Order		//belong to order
}

func (orderProduct *OrderProduct) GetRelationship() map[string]interface{}{
	relationship := map[string]interface{} {
		"Product": &orderProduct.Product,
		"Order": &orderProduct.Order,
	}
	return relationship
}
