package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UID       string	`schema:"uid"`
	Username  string	`schema:"username"`
	Password  string	`schema:"password"`
	Email     string	`schema:"email"`
	Gender    string	`schema:"gender"`
	Role      string	`schema:"role"`
	Avatar    string	`schema:"avatar"`
	Phone     string	`schema:"phone"`
	Provider  string	`schema:"provider"`
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
