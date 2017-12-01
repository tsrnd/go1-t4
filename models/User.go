package models

import (
	"log"

	"github.com/goweb4/database"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UID      string `schema:"uid"`
	UserName string `schema:"username"`
	Password string `schema:"password"`
	Email    string `schema:"email"`
	Gender   string `schema:"gender"`
	Role     string `schema:"role"`
	Avatar   string `schema:"avatar"`
	Phone    string `schema:"phone"`
	Provider string `schema:"provider"`
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
