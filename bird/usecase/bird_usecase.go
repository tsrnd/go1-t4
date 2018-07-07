package usecase

import (
	"fmt"
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


func (b *birdUsecase) Create(name, color string, description string) (*model.Bird, error) {
	id, err := b.repo.Create(name, color, description)
	if err != nil {
		fmt.Println(err)
	}
	
	return b.repo.GetByID(id)
}

func (uc *birdUsecase) GetByID(id int64) (*model.Bird, error) {
	return uc.repo.GetByID(id)
}

// NewProductUsecase func
func NewBirdUsecase(uc repos.BirdRepository) BirdUsecase {
	return &birdUsecase{uc}
}
