package entity

import "time"

type Portfolio struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Active    bool      `json:"active" gorm:"notNull;default:true"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
	Name      string    `json:"name"`
	CountryID uint      `json:"country_id"`
	Country   Country   `json:"-" gorm:"foreignKey:CountryID"`
}
