package house

// @todo custom validators https://pkg.go.dev/github.com/gin-gonic/gin#readme-custom-validators
type House struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	Address    string `json:"address" binding:"required"`
	Bedrooms   int    `json:"bedrooms"`
	Bathrooms  int    `json:"bathrooms"`
	Garage     int    `json:"garage"`
	Floorspace int    `json:"floorspace"`
	Landarea   int    `json:"landarea"`
}
