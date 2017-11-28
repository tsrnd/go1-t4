package models

import (
	"fmt"
	"time"
)

type User struct {
	ID        int
	UID       string
	FirstName string
	LastName  string
	Password  string
	Email     string
	Gender    string
	Role      string
	Avatar    string
	Phone     string
	Provider  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

var users = []User{}

func GetUsers() []User {
	return users
}

func GetUserByID(ID int) (*User, error) {
	for _, user := range users {
		if user.ID == ID {
			return &user, nil
		}
	}
	return nil, fmt.Errorf(fmt.Sprintf("User with ID %v not found", ID))
}
