package models

import (
	"github.com/goweb4/database"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	ProId    string `schema:"pro_id"`
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
	WithConnectionDB(func(db *database.DB) {
		db.Where(&userInfo).First(&user)
	})
	return user
}

func GetUserById(id uint) (user User, err error) {
	WithConnectionDB(func(db *database.DB) {
		err = db.Where("id = ?", id).Find(&user).Error
	})
	return user, err
}

func GetUserByUserName(name string) (user User, err error) {
	WithConnectionDB(func(db *database.DB) {
		err = db.Where("user_name = ?", name).Find(&user).Error
	})
	return user, err
}

func CreateUser(user User) (err error) {
	WithConnectionDB(func(db *database.DB) {
		err = db.Create(&user).Error
	})
	return err
}

func UpdateUser(user *User) (errUpdate error) {
	WithConnectionDB(func(db *database.DB) {
		errUpdate = db.Save(&user).Error
	})
	return errUpdate
}

func DeleteUser(id uint) (err error) {
	user := User{}
	WithConnectionDB(func(db *database.DB) {
		user, err = GetUserById(id)
		if err == nil {
			err = db.Delete(&user).Error
		}
	})
	return err
}
