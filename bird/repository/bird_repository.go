package repository

import (
	// model "github.com/goweb4/bird"
)

type BirdRepository interface {
	Create(title, description string, userID int64) (int64, error)
}
