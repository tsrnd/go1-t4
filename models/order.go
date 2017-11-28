package models

import (
	"time"
	"github.com/goweb4/database"
)

type Order struct {
	ID         int
	UserID     int
	TotalMoney float64
	OrderDate  time.Time
	Status     bool
}

func GetOrder(id int) (order Order, err error) {
  order = Order{}
  err = database.Db.Where("id = ?", id).Find(&order).Error

  return order, err
}
