package models

type User struct {
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

func (user *User) GetRelationship() map[string]interface{} {
	relationship := map[string]interface{}{
		"Orders": &user.Orders,
	}
	return relationship
}

// func GetUser(userInfo User) (user User) {
// 	database.DBCon.Where(&userInfo).First(&user)
// 	return user
// }

// func GetUserById(id uint) (user User, err error) {
// 	err = database.DBCon.Where("id = ?", id).Find(&user).Error
// 	return user, err
// }

// func GetUserByUserName(name string) (user User, err error) {
// 	err = database.DBCon.Where("user_name = ?", name).Find(&user).Error
// 	return user, err
// }

// func CreateUser(user User) (err error) {
// 	err = database.DBCon.Create(&user).Error
// 	return err
// }

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
