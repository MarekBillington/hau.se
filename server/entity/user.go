package entity

import (
	"html"
	"strings"

	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Active    bool      `json:"active" gorm:"notNull;default:true"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
	Email     string    `json:"email" gorm:"notNull;uniqueIndex;"`
	Password  string    `json:"-" gorm:"notNull"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Language  string    `json:"language"`
}

// Gorm functionality will trigger this during saving process
func (u *User) BeforeCreate(tx *gorm.DB) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)

	u.Email = html.EscapeString(strings.TrimSpace(u.Email))

	return nil
}
