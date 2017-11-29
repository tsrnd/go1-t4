package models

import (
	"log"
  "github.com/goweb4/database"
	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	UserID     uint 			`schema:"user_id"`
	TotalMoney float64		`schema:"total_money"`
	Status     bool				`schema:"status"`
}

func GetOrder(id int) (order Order, err error) {
  order = Order{}
  db, errConnection := database.DBConnection(); if errConnection != nil {
	  log.Fatal(errConnection)
  }
  defer db.Close()
  err = db.Where("id = ?", id).Find(&order).Error

  return order, err
}

func GetOrdersByUser(id int) (orders []Order, err error) {
	orders = []Order{}
	db, errConnection := database.DBConnection(); if errConnection != nil {
		log.Fatal(errConnection)
	}
	defer db.Close()

	err = db.Where("user_id = ?", id).Find(&orders).Error

	return orders, err
}
