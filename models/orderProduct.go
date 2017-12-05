package models

import (
	"github.com/jinzhu/gorm"
	"github.com/goweb4/database"
)

type OrderProduct struct {
	gorm.Model
	OrderID   uint		`schema:"order_id"`
	ProductID uint		`schema:"product_id"`
	Quantity  uint		`schema:"quantity"`
	Product	  Product	//belong to product
	Order	    Order		//belong to order
}

func (orderProduct *OrderProduct) GetRelationship() map[string]interface{}{
	relationship := map[string]interface{} {
		"Product": &orderProduct.Product,
		"Order": &orderProduct.Order,
	}
	return relationship
}

func CreateOrderProduct(orderProduct *OrderProduct) (orderProID uint, err error){
	err = database.Tx.Create(&orderProduct).Error
	if orderProID == 0 && err == nil {
		orderProID = orderProduct.ID
	}
	return orderProID, err
}
