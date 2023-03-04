package entity

import (
	"html"
	"strings"

	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Active    bool      `json:"active" gorm:"notNull;default:true"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updateAt"`
	Email     string    `json:"email" gorm:"notNull;uniqueIndex;"`
	Password  string    `json:"password" gorm:"notNull"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
}

// Gorm functionality will trigger this during saving process
func (u *User) BeforeSave(tx *gorm.DB) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)

	u.Email = html.EscapeString(strings.TrimSpace(u.Email))

	return nil
}
