package psql

import (
	"database/sql"

	// model "github.com/goweb4/bird"
	repo "github.com/goweb4/bird/repository"
)

type birdRepository struct {
	DB *sql.DB
}

func (m *birdRepository) Create(title, description string, userID int64) (int64, error) {
	const query = `
		insert into birds (
			title,
			description,
			user_id
		) values (
			$1,
			$2,
			$3
		) returning id
	`
	var id int64
	err := m.DB.QueryRow(query, title, description, userID).Scan(&id)
	return id, err
}

func (m *productRepository) GetByID(id int64) (*model.Product, error) {
	const query = `
		select
			id,
			title,
			description,
			user_id
		from
			products
		where
			id = $1
	`
	var product model.Product
	err := m.DB.QueryRow(query, id).Scan(&product.ID, &product.Title, &product.Description, &product.UserID)
	return &product, err
}

func (m *birdRepository) GetByTitle(title string) (*model.Bird, error) {
	const query = `
		select
			id,
			title,
			description,
			user_id
		from
			birds
		where
			title = $1
	`
	var bird model.Bird
	err := m.DB.QueryRow(query, title).Scan(&bird.ID, &bird.Title, &bird.Description, &bird.UserID)
	return &bird, err
}


func NewBirdRepository(DB *sql.DB) repo.BirdRepository {
	return &birdRepository{DB}
}
