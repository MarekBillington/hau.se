package dto

type NewHouse struct {
	PortfolioID uint       `json:"portfolio_id" binding:"required"`
	Bedrooms    int        `json:"bedrooms"`
	Bathrooms   int        `json:"bathrooms"`
	Garage      int        `json:"garage"`
	Floorspace  int        `json:"floorspace"`
	Landarea    int        `json:"landarea"`
	PropertyID  uint       `json:"property_id"`
	NewAddress  NewAddress `json:"address" binding:"required"`
}
