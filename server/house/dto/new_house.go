package dto

type NewHouse struct {
	PortfolioID uint       `json:"portfolio_id" binding:"required"`
	Street1     string     `json:"street_1" binding:"required"`
	Street2     string     `json:"street_2"`
	Postcode    string     `json:"postcode" binding:"required"`
	Town        string     `json:"town" binding:"required"`
	CountryID   uint       `json:"country_id" binding:"required"`
	Bedrooms    int        `json:"bedrooms"`
	Bathrooms   int        `json:"bathrooms"`
	Garage      int        `json:"garage"`
	Floorspace  int        `json:"floorspace"`
	Landarea    int        `json:"landarea"`
	PropertyID  uint       `json:"property_id"`
}
