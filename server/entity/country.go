package entity

import "time"

type Country struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Active    bool      `json:"active" gorm:"notNull;default:true"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
	Name      string    `json:"name"`
	Currency  string    `json:"currency"`
	HasStates bool      `json:"has_states" gorm:"default:false"`
}
