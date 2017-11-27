package models

import (
	"time"
)

type Order struct {
	ID         int
	UserID     int
	TotalMoney float64
	OrderDate  time.Time
	Status     bool
}
