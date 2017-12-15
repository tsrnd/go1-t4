package psql

import (
	"database/sql"

	model "github.com/goweb4/book"
	repo "github.com/goweb4/book/repository"
)

type bookRepository struct {
	DB *sql.DB
}

func (m *bookRepository) Create(name, content string, userID int64) (int64, error) {
	const query = `
		insert into books (
			name,
			content,
			user_id
		) values (
			$1,
			$2,
			$3
		) returning id
	`
	var id int64
	err := m.DB.QueryRow(query, name, content, userID).Scan(&id)
	return id, err
}

func (m *bookRepository) GetByID(id int64) (*model.Book, error) {
	const query = `
		select
			id,
			name,
			content,
			user_id
		from
			books
		where
			id = $1
	`
	var book model.Book
	err := m.DB.QueryRow(query, id).Scan(&book.ID, &book.Name, &book.Content, &book.UserID)
	return &book, err
}

func (m *bookRepository) GetByName(name string) (*model.Book, error) {
	const query = `
		select
			id,
			name,
			content,
			user_id
		from
			books
		where
			name = $1
	`
	var book model.Book
	err := m.DB.QueryRow(query, name).Scan(&book.ID, &book.Name, &book.Content, &book.UserID)
	return &book, err
}

func (m *bookRepository) Update(bookID int64, name, content string) error {
	const query = `
    update books set
			name = $1,
			content = $2
    where
      id = $3
	`
	_, err := m.DB.Exec(query, name, content, bookID)
	return err
}

func (m *bookRepository) Delete(id int64) error {
	const query = `delete from books where id = $1`
	_, err := m.DB.Exec(query, id)
	return err
}

func (m *bookRepository) Fetch(offset, limit int64) ([]*model.Book, error) {
	const query = `
		select
			id,
			name,
			content,
			user_id
		from
			books
		limit $1 offset $2
	`
	books := make([]*model.Book, 0)

	rows, err := m.DB.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var book model.Book
		err = rows.Scan(&book.ID, &book.Name, &book.Content, &book.UserID)
		if err != nil {
			return nil, err
		}
		books = append(books, &book)
	}
	return books, err
}

// NewBookRepository func
func NewBookRepository(DB *sql.DB) repo.BookRepository {
	return &bookRepository{DB}
}
