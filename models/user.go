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

func GetUser(userInfo User) (user User) {
	db, errConnection := database.DBConnection()
	if errConnection != nil {
		log.Fatal(errConnection)
	}
	defer db.Close()
	db.Where(&userInfo).First(&user)
	return user
}

func GetUserById(id uint) (user User, err error) {
	user = User{}
	db, errConnection := database.DBConnection()
	if errConnection != nil {
		log.Fatal(errConnection)
	}
	defer db.Close()

	err = db.Where("id = ?", id).Find(&user).Error

	return user, err
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

func UpdateUser(user *User) error {
	db, errConnection := database.DBConnection()
	if errConnection != nil {
		log.Fatal(errConnection)
	}
	defer db.Close()
	errUpdate := db.Save(&user).Error
	return errUpdate
}

func DeleteUser(id uint) error {
	user := User{}
	db, errConnection := database.DBConnection()
	if errConnection != nil {
		log.Fatal(errConnection)
	}
	defer db.Close()

	user, errGet := GetUserById(id)
	if errGet != nil {
		return errGet
	}
	err := db.Delete(&user).Error
	return err
}
