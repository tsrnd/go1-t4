package models

import (
	"log"
	"time"

	"github.com/goweb4/database"
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

func GetUser(id int) (user User, err error) {
	user = User{}
	db, errConnection := database.DBConnection()
	if errConnection != nil {
		log.Fatal(errConnection)
	}
	defer db.Close()
	err = db.Where("id = ?", id).Find(&user).Error

	return user, err
}

func GetUsers() (users []User, err error) {
	users = []User{}
	db, errConnection := database.DBConnection()
	if errConnection != nil {
		log.Fatal(errConnection)
	}
	defer db.Close()

	err = db.Find(&users).Error

	return users, err
}

func UpdateUser(id int, oldUser map[string]interface{}) error {
	user := User{}
	db, errConnection := database.DBConnection()
	if errConnection != nil {
		log.Fatal(errConnection)
	}
	defer db.Close()

	errFindUser := db.Where("id = ?", id).Find(&user).Error
	if errFindUser != nil {
		return errFindUser
	}
	errUpdate := db.Model(&user).Updates(oldUser).Error
	return errUpdate
}

func DeleteUser(id int) error {
	db, errConnection := database.DBConnection()
	if errConnection != nil {
		log.Fatal(errConnection)
	}
	defer db.Close()

	err := db.Where("id = ?", id).Delete(&User{}).Error

	return err
}

func CreateUser(user User) error {
	db, errConnection := database.DBConnection()
	if errConnection != nil {
		log.Fatal(errConnection)
	}
	defer db.Close()

	err := db.Create(&user).Error

	return err
}
