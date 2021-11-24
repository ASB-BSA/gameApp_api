package models

import (
	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
	Model
	User     string `json:"user"`
	Password []byte `json:"-"`
}

func (user *Admin) SetPassword(password string) {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	user.Password = hashPassword
}

func (user *Admin) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
