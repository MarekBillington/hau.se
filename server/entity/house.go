package entity

import "time"

// @todo custom validators https://pkg.go.dev/github.com/gin-gonic/gin#readme-custom-validators
type House struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Active      bool      `json:"active" gorm:"notNull;default:true"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"update_at"`
	PortfolioID uint      `json:"portfolio_id"`
	Portfolio   Portfolio `json:"-" gorm:"foreignKey:PortfolioID"`
	Street1     string    `json:"street_1"`
	Street2     string    `json:"street_2"`
	Postcode    string    `json:"postcode"`
	Town        string    `json:"town"`
	CountryID   uint      `json:"country_id"`
	Country     Country   `json:"-" gorm:"foreignKey:CountryID"`
	Bedrooms    int       `json:"bedrooms"`
	Bathrooms   int       `json:"bathrooms"`
	Garage      int       `json:"garage"`
	Floorspace  int       `json:"floorspace"`
	Landarea    int       `json:"landarea"`
	PropertyID  uint      `json:"property_id" gorm:"default:null"`
	Property    *House    `json:"-" gorm:"foreignKey:PropertyID;default:null"`
}

// The Property field is a self-referencial nullable id
// i.e. a house can be inside a property (apartment building with flats)
// 		or a property of its own (standalone house)
