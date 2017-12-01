package models

import (
	"log"

	"github.com/goweb4/database"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Proid    string `schema:"proid"`
	UserName string `schema:"username"`
	Password string `schema:"password"`
	Email    string `schema:"email"`
	Gender   string `schema:"gender"`
	Role     string `schema:"role"`
	Avatar   string `schema:"avatar"`
	Phone    string `schema:"phone"`
	Address  string `schema:"address"`
	Provider string `schema:"provider"`
}

var users = []User{}

func GetUsers() []User {
	return users
}

func GetUser(userInfo User) (user User) {
	db, errConnection := database.DBConnection()
	if errConnection != nil {
		log.Fatal(errConnection)
	}
	defer db.Close()
	db.Where(&userInfo).First(&user)
	return user
}
