package usecase

import (
	model "github.com/goweb4/bird"
	repos "github.com/goweb4/bird/repository"
)

// ProductUsecase interface
type BirdUsecase interface {
	Create(title, description string, userID int64) (*model.Bird, error)
}

type birdUsecase struct {
	repo repos.BirdRepository
}

func (uc *birdUsecase) Create(title, description string, userID int64) (*model.Bird, error) {
	// exist, _ := uc.GetByTitle(title)
	// if exist != nil {
	// 	return nil, model.ConflictError
	// }

	
	return nil, nil
	// return uc.GetByID(id)
}

// NewProductUsecase func
func NewBirdUsecase(uc repos.BirdRepository) BirdUsecase {
	return &birdUsecase{uc}
}
