package usecase

import (
	model "github.com/goweb4/gift"
	repos "github.com/goweb4/gift/repository"
)

type GiftUsecase interface {
	Create(fromUserID, toUserID, productID int64, message string) (*model.Gift, error)
	Update(id, productID int64, message string) (*model.Gift, error)
	Delete(id int64) error
	GetByID(id int64) (*model.Gift, error)
	GetByFromUserID(id int64) ([]*model.Gift, error)
	Fetch(offset, limit int64) ([]*model.Gift, error)
}

type giftUsecase struct {
	repo repos.repository
}

func (gU *giftUsecase) Create(fromUserID, toUserID, productID int64, message string) (*model.Gift, error) {
	id, err := gU.repo.Create(fromUserID, toUserID, productID, message)
	if err != nil {
		return nil, err
	}
	return gU.GetByID(id)
}

func (gU *giftUsecase) Update(id, productID int64, message string) (*model.Gift, error) {
	err := gU.repo.Update(id, productID, message)
	if err != nil {
		return nil, err
	}
	return gU.GetByID(id)
}

func (gU *giftUsecase) Delete(id int64) error {
	return gU.repo.Delete(id)
}

func (gU *giftUsecase) GetByID(id int64) (*model.Gift, error) {
	return gU.repo.GetByID(id)
}

func (gU *giftUsecase) GetByFromUserID(id int64) ([]*model.Gift, error) {
	return gU.repo.GetByFromUserID(id)
}

func (gU *giftUsecase) Fetch(offset, limit int64) ([]*model.Gift, error) {
	if limit == 0 {
		limit = 10
	}
	return gU.repo.Fetch(offset, limit)
}
