package models

import (
	"github.com/goweb4/database"
	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	UserID     		uint 						`schema:"user_id"`
	TotalMoney 		float64					`schema:"total_money"`
	Status     		bool						`schema:"status"`
	OrderProducts	[]OrderProduct	//has many order products
	User					User						//belong to user
}

func GetOrder(id int) (order Order, err error) {
	WithConnectionDB(func(db *database.DB) {
		err = db.Where("id = ?", id).Find(&order).Error
	})
	return order, err
}

func GetOrdersByUser(id int) (orders []Order, err error) {
	WithConnectionDB(func(db *database.DB) {
		err = db.Where("user_id = ?", id).Find(&orders).Error
	})
	return orders, err
}
