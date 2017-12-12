package models

import (
	"fmt"

	"github.com/goweb4/database"
)

type User struct {
	Model
	ProId    string  `schema:"pro_id"`
	UserName string  `schema:"username"`
	Password string  `schema:"password"`
	Email    string  `schema:"email"`
	Gender   string  `schema:"gender"`
	Role     string  `schema:"role"`
	Avatar   string  `schema:"avatar"`
	Phone    string  `schema:"phone"`
	Address  string  `schema:"address"`
	Provider string  `schema:"provider"`
	Orders   []Order //has many order
}

const ADMIN_ROLE = "admin"
const USER_ROLE = "normal_user"

// func GetUser(userInfo User) (user User) {
// 	database.DBCon.Where(&userInfo).First(&user)
// 	return user
// }

// func GetUserById(id uint) (user User, err error) {
// 	err = database.DBCon.Where("id = ?", id).Find(&user).Error
// 	return user, err
// }

func GetUserByUserName(name string) (user User, err error) {
	err = database.DBCon.Db.QueryRow("SELECT password, email, role FROM users where user_name = $1", name).Scan(&user.Password, &user.Email, &user.Role)
	if err != nil {
		fmt.Println("get user by name has an error: ", err)
	}
	return user, err
}

func CreateUser(user User) (err error) {
	err = database.DBCon.Db.
		QueryRow("INSERT INTO users (user_name, email, password, phone, address, created_at, role) VALUES($1,$2,$3,$4,$5,$6,$7) returning id;",
			user.UserName, user.Email, user.Password, user.Phone, user.Address, user.CreatedAt, "user").Scan(&user.ID)
	return err
}

// func UpdateUser(user *User) (errUpdate error) {
// 	errUpdate = database.DBCon.Save(&user).Error
// 	return errUpdate
// }

// func DeleteUser(id uint) (err error) {
// 	user := User{}
// 	user, err = GetUserById(id)
// 	if err == nil {
// 		err = database.DBCon.Delete(&user).Error
// 	}
// 	return err
// }
