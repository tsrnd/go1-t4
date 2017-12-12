package utils

import (
	"github.com/goweb4/models"
	"golang.org/x/crypto/bcrypt"
)

func CheckCredential(info models.User) bool {
	user, err := models.GetUserByUserName(info.UserName)
	if err != nil {
		return false
	}
	if CheckPasswordHash(info.Password, user.Password) {
		return true
	}
	return false
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CheckUserExist(name string) bool {
	_, err := models.GetUserByUserName(name)
	if err != nil {
		return false
	}
	return true
}

func IsAdminRole(userName string) bool {
	user, err := models.GetUserByUserName(userName)
	if err != nil {
		return false
	}
	if user.Role == "admin" {
		return true
	}
	return false
}
