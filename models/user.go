package models

import (
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

func GetUserById(id uint) (user User, err error) {
	err = database.DBCon.Db.QueryRow("SELECT password, email, role FROM users where id = $1", id).Scan(&user.Password, &user.Email, &user.Role)
	return user, err
}

func GetUserByUserName(name string) (user User, err error) {
	err = database.DBCon.Db.QueryRow("SELECT id, password, user_name, email, phone, address, role FROM users where user_name = $1", name).Scan(&user.ID, &user.Password, &user.UserName, &user.Email, &user.Phone, &user.Address, &user.Role)
	return user, err
}

func CreateUser(user User) (err error) {
	err = database.DBCon.Db.
		QueryRow("INSERT INTO users (user_name, email, password, phone, address, created_at, role) VALUES($1,$2,$3,$4,$5,$6,$7) returning id;",
			user.UserName, user.Email, user.Password, user.Phone, user.Address, user.CreatedAt, "user").Scan(&user.ID)
	return err
}

func UpdateUser(user *User) (errUpdate error) {
	_, errUpdate = database.DBCon.Db.
		Exec("UPDATE users SET user_name=$1, email=$2, password=$3, phone=$4, address=$5, updated_at=$6, role=$7 WHERE id = $8",
			user.UserName, user.Email, user.Password, user.Phone, user.Address, user.UpdatedAt, user.Role, user.ID)
	return
}

func DeleteUser(id uint) (err error) {
	_, err = GetUserById(id)
	if err != nil {
		return
	}
	_, err = database.DBCon.Db.
		Exec("DELETE FROM users WHERE id = $1", id)
	return err
}
