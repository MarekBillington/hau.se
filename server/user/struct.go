package user

import "time"

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
