package db

import (
	"golang.org/x/crypto/bcrypt"
	"project/go-vets-backend/models"
)

func Login(email string, password string) (models.User, bool) {
	user, found := CheckUserExists(email)
	if !found {
		return user, false
	}
	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return user, false
	}
	return user, true
}
