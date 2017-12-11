package models

import (
	
)

type Payment struct {
	Model
	Method string
}

// func GetPayments() (payments []Payment, err error) {
// 	err = database.DBCon.Find(&payments).Error
// 	return payments, err
// }
