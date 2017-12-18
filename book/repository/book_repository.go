package repository

import (
	model "github.com/goweb4/book"
)

// BookRepository interface
type BookRepository interface {
	Create(name, content string, userID int64) (int64, error)
	Update(BookID int64, name, content string) error
	Delete(id int64) error
	GetByID(id int64) (*model.Book, error)
	GetByName(name string) ([]*model.Book, error)
	Fetch(offset, limit int64) ([]*model.Book, error)
}
