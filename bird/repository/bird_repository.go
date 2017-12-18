package repository

import (
	model "github.com/goweb4/bird"
)

type BirdRepository interface {
	Create(name, color string, description string) (int64, error)
	GetByID(id int64) (*model.Bird, error)
}
