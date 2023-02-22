package user

import (
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (u *user) BeforeSave(tx *gorm.DB) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)

	u.Email = html.EscapeString(strings.TrimSpace(u.Email))

	return nil
}
