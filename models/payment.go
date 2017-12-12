package models

import "github.com/goweb4/database"

type Payment struct {
	Model
	Method string
}

func GetPayments() (payments []Payment, err error) {
	rows, err := database.DBCon.Db.Query("SELECT method FROM payments")
	if err != nil {
		return
	}
	for rows.Next() {
		payment := Payment{}
		err = rows.Scan(&payment.Method)
		if err != nil {
			return
		}
		payments = append(payments, payment)
	}
	rows.Close()
	return payments, nil
}
