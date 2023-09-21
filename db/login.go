package db

import (
	"github.com/GonzaPiccinini/twitter-go/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(email string, password string) (models.User, bool) {
	u, ok, _ := UserExists(email)
	if !ok {
		return u, false
	}

	passwordReq := []byte(password)
	passwordDb := []byte(u.Password)
	err := bcrypt.CompareHashAndPassword(passwordDb, passwordReq)
	if err != nil {
		return u, false
	}

	return u, true
}
