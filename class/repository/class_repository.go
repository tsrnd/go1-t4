package repository

import (
	"github.com/goweb4/class"
)

type ClassRepository interface {
	Create(string, int) error
	Delete(int) error
	Update(class.Class) error
}
