package dto

type UpdateHouse struct {
	PortfolioID   uint          `json:"portfolio_id"`
	Bedrooms      int           `json:"bedrooms"`
	Bathrooms     int           `json:"bathrooms"`
	Garage        int           `json:"garage"`
	Floorspace    int           `json:"floorspace"`
	Landarea      int           `json:"landarea"`
	PropertyID    uint          `json:"property_id"`
	UpdateAddress UpdateAddress `json:"address"`
}
