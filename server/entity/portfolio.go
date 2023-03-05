package entity

import "time"

type Portfolio struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Active    bool      `json:"active" gorm:"notNull;default:true"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updateAt"`
	Name      string    `json:"name"`
	CountryID uint      `json:"country_id"`
	Country   Country   `json:"-" gorm:"foreignKey:CountryID"`
}
