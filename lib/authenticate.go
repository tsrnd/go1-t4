package lib

import (
	"errors"
	"github.com/goweb4/models"
	"golang.org/x/crypto/bcrypt"
)

func CheckCredentials(username string, password string) (user *models.User, err error) {
	user = &models.User{UserName: username}
	err = user.ReadByUsername()
	if err != nil {
		if err.Error() == "<QuerySeter> no row found" {
			err = errors.New("This username is not exist")
		}
		return user, err
	} else if user.Id < 1 {
		err = errors.New("No user founded")
		return user, err
	} else {
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			err = errors.New("Password not matched")
			return user, err
		}
		return user, err
	}
}

func HashPass(password string) (hashPass string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
