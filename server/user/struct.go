package user

import "time"

type user struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Active    bool      `json:"active" gorm:"notNull;default:true"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updateAt"`
	Email     string    `json:"email" gorm:"notNull;uniqueIndex;"`
	Password  string    `json:"password" gorm:"notNull"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
}

type register struct {
	Email    string `json"email" binding:"required"`
	Password string `json"password" binding:"required"`
}

type refresh struct {
	Token string `json"token" binding:"required"`
}
