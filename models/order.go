package models

import (
	"fmt"
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

func TestDB()  string {
	if database.Db != nil {
		//use db. query here
		fmt.Println("DB is OKE")
    } else {
        fmt.Println("DB object is NULL")
	}
	return "Oke"
}
