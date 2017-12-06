package models

import (
	"github.com/jinzhu/gorm"
	"github.com/goweb4/database"
)

type Payment struct {
	gorm.Model
	Method string
}

func GetPayments() (payments []Payment, err error) {
	err = database.DBCon.Find(&payments).Error
	return payments, err
}
