package entity

import "time"

type Address struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Active    bool      `json:"active" gorm:"notNull;default:true"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
	Street1   string    `json:"street_1"`
	Street2   string    `json:"street_2"`
	Postcode  string    `json:"postcode"`
	Town      string    `json:"town"`
	CountryID uint      `json:"country_id"`
	Country   Country   `json:"-" gorm:"foreignKey:CountryID"`
	HouseID   uint      `json:"house_id"`
	House     House     `json:"-" gorm:"foreignKey:HouseID"`
}
