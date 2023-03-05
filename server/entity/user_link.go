package entity

import "time"

type UserLink struct {
	ID          uint      ``
	Active      bool      `json:"active" gorm:"notNull;default:true"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"update_at"`
	UserID      uint      `json:"user_id"`
	User        User      `json:"-" gorm:"foreignKey:UserID"`
	PortfolioID uint      `json:"portfolio_id"`
	Portfolio   Portfolio `json:"-" gorm:"foreignKey:PortfolioID"`
}
