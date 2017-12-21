package usecase

import (
	model "github.com/goweb4/book"
	repos "github.com/goweb4/book/repository"
)

// BookUsecase interface
type BookUsecase interface {
	Create(name, content string, userID int64) (*model.Book, error)
	GetByID(id int64) (*model.Book, error)
	GetByName(name string) ([]*model.Book, error)
	Update(bookID int64, name, content string) (*model.Book, error)
	Delete(id int64) error
	Fetch(offset, limit int64) ([]*model.Book, error)
}

type bookUsecase struct {
	repo repos.BookRepository
}

func (uc *bookUsecase) Create(name, content string, userID int64) (*model.Book, error) {
	exist, _ := uc.GetByName(name)
	if exist != nil {
		return nil, model.ConflictError
	}

	id, err := uc.repo.Create(name, content, userID)
	if err != nil {
		return nil, err
	}

	return uc.GetByID(id)
}

func (uc *bookUsecase) GetByID(id int64) (*model.Book, error) {
	return uc.repo.GetByID(id)
}

func (uc *bookUsecase) GetByName(name string) ([]*model.Book, error) {
	result, err := uc.repo.GetByName(name)
	return result, err
}

func (uc *bookUsecase) Update(bookID int64, name, content string) (*model.Book, error) {
	err := uc.repo.Update(bookID, name, content)
	if err != nil {
		return nil, err
	}
	return uc.GetByID(bookID)
}

func (uc *bookUsecase) Delete(id int64) error {
	existedBook, _ := uc.GetByID(id)
	if existedBook == nil {
		return model.NotFoundError
	}
	return uc.repo.Delete(id)
}

func (uc *bookUsecase) Fetch(offset, limit int64) ([]*model.Book, error) {
	if limit == 0 {
		limit = 10
	}
	return uc.repo.Fetch(offset, limit)
}

// NewBookUsecase func
func NewBookUsecase(uc repos.BookRepository) BookUsecase {
	return &bookUsecase{uc}
}
