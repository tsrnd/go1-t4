package usecase

import (
	"github.com/goweb4/class"
	repo "github.com/goweb4/class/repository"
)

type ClassUsecase interface {
	GetClassById(id int) (class.Class, error)
	CreateClass(string, int) error
	UpdateClass() error
	DeleteClass(int) error
}

type classUsecase struct {
	Repo repo.ClassRepository
}

func (c classUsecase) CreateClass(name string, number int) error {
	err := c.Repo.Create(name, number)
	return err
}

func (c classUsecase) GetClassById(id int) (cl class.Class, err error) {
	return cl, err
}

func (c classUsecase) UpdateClass() error {
	return nil
}

func (c classUsecase) DeleteClass(int) error {
	return nil
}

func NewClassUsecase(r repo.ClassRepository) ClassUsecase {
	return &classUsecase{r}
}
