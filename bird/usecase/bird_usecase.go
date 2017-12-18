package usecase

import (
	model "github.com/goweb4/bird"
	repos "github.com/goweb4/bird/repository"
)

// ProductUsecase interface
type BirdUsecase interface {
	Create(name, color string, description string) (*model.Bird, error)
	GetByID(id int64) (*model.Bird, error)
	
}

type birdUsecase struct {
	repo repos.BirdRepository
}


func (uc *birdUsecase) Create(name, color string, description string) (*model.Bird, error) {
	// exist, _ := uc.GetByTitle(title)
	// if exist != nil {
	// 	return nil, model.ConflictError
	// }

	
	return nil, nil
	// return uc.GetByID(id)
}

func (uc *birdUsecase) GetByID(id int64) (*model.Bird, error) {
	return uc.repo.GetByID(id)
}

// NewProductUsecase func
func NewBirdUsecase(uc repos.BirdRepository) BirdUsecase {
	return &birdUsecase{uc}
}
