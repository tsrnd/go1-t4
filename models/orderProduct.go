package models

import (
	// "database/sql"
)

type OrderProduct struct {
	Model
	OrderID   uint     `schema:"order_id"`
	ProductID uint     `schema:"product_id"`
	Quantity  uint     `schema:"quantity"`
	Product   *Product //belong to product
	Order     *Order   //belong to order
}
