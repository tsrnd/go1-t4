package repository

import (
	model "github.com/goweb4/gift"
)

type GiftRepository interface {
	Create(fromUserID, toUserID, productID int64, message string) (int64, error)
	Update(id, productID int64, message string) (error)
	Delete(id int64) (error)
	GetByID(id int64) (*model.Gift, error)
	GetByFromUserID(id int64) (*model.Gift, error)
	Fetch(offset, limit int64) ([]*model.Gift, error)
}
