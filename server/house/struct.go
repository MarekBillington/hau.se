package house

import "time"

// @todo custom validators https://pkg.go.dev/github.com/gin-gonic/gin#readme-custom-validators
type house struct {
	ID         int       `json:"id" gorm:"primaryKey"`
	Active     bool      `json:"active" gorm:"notNull;default:true"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdateAt   time.Time `json:"updateAt"`
	Address    string    `json:"address" binding:"required" gorm:"index"`
	Bedrooms   int       `json:"bedrooms"`
	Bathrooms  int       `json:"bathrooms"`
	Garage     int       `json:"garage"`
	Floorspace int       `json:"floorspace"`
	Landarea   int       `json:"landarea"`
}
