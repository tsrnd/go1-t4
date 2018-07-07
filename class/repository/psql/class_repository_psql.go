package psql

import (
	"database/sql"

	"github.com/goweb4/class"
	repo "github.com/goweb4/class/repository"
)

type classRepository struct {
	DB *sql.DB
}

func (c *classRepository) Create(name string, number int) error {
	var id int
	query := "INSERT into class(name,student_number) values($1,$2) returning id;"
	err := c.DB.QueryRow(query, name, number).Scan(&id)
	return err
}

func (c *classRepository) Delete(int) error {
	return nil
}

func (c *classRepository) Update(class.Class) error {
	return nil
}

func NewClassRepository(DB *sql.DB) repo.ClassRepository {
	return &classRepository{DB}
}
