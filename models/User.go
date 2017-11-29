package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
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
}

var users = []User{}

func GetUsers() []User {
	return users
}

func GetUserByID(ID uint) (*User, error) {
	for _, user := range users {
		if user.ID == ID {
			return &user, nil
		}
	}
	return nil, fmt.Errorf(fmt.Sprintf("User with ID %v not found", ID))
}
