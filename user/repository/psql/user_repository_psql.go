package psql

import (
	"database/sql"

	"github.com/goweb4/services/crypto"
  model "github.com/goweb4/user"
  repo  "github.com/goweb4/user/repository"
)

type userRepository struct {
	DB *sql.DB
}

func (m *userRepository) GetByID(id int) (*model.User, error) {
	const query = `
    select
      id,
      email,
      name
    from
      users
    where
      id = $1
  `
	var user model.User
	err := m.DB.QueryRow(query, id).Scan(&user.ID, &user.Email, &user.Name)
	return &user, err
}

func (m *userRepository) GetByEmail(email string) (*model.User, error) {
	const query = `
    select
      id,
      email,
      name
    from
      users
    where
      email = $1
  `
	var user model.User
	err := m.DB.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Name)
	return &user, err
}

func (m *userRepository) GetPrivateDetailsByEmail(email string) (*model.PrivateUserDetails, error) {
	const query = `
    select
      id,
      password,
      salt
    from
      users
    where
      email = $1
  `
	var u model.PrivateUserDetails
	err := m.DB.QueryRow(query, email).Scan(&u.ID, &u.Password, &u.Salt)
	return &u, err
}

func (m *userRepository) Create(email, name, password string) (int, error) {
	const query = `
    insert into users (
      email,
      name,
      password,
      salt
    ) values (
      $1,
      $2,
      $3,
      $4
    )
    returning id
  `
	salt := crypto.GenerateSalt()
	hashedPassword := crypto.HashPassword(password, salt)
	var id int
	err := m.DB.QueryRow(query, email, name, hashedPassword, salt).Scan(&id)
	return id, err
}

// NewUserRepository func
func NewUserRepository(DB *sql.DB) repo.UserRepository {
	return &userRepository{DB}
}
