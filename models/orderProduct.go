package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/goweb4/database"
)

type OrderProduct struct {
	gorm.Model
	OrderID   uint		`schema:"order_id"`
	ProductID uint		`schema:"product_id"`
	Quantity   uint		`schema:"quantity"`
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

func CreateOrderProduct(orderProduct *OrderProduct, lastItem bool) (orderProID uint, err error){
	WithConnectionDB(func(db *database.DB) {
		err = db.Create(&orderProduct).Error
		if orderProID == 0 && err == nil {
			orderProID = orderProduct.ID
		}
	})
	return orderProID, err
}
