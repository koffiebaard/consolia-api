package models

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/jinzhu/gorm"
	_"fmt"
)

type User struct {
	Username   	string
	Password	string
}

func (c *User) Authenticate (db gorm.DB, username string, password string) bool {

	user := User{}
	if db.Where("username = ?", username).First(&user).Error == nil {

		//hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

		if err == nil {
			return true
		}
	}

	return false;
}
