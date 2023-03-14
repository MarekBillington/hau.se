package dto

type UpdateHouse struct {
	PortfolioID uint   `json:"portfolio_id"`
	Street1     string `json:"street_1"`
	Street2     string `json:"street_2"`
	Postcode    string `json:"postcode"`
	Town        string `json:"town"`
	CountryID   uint   `json:"country_id"`
	Bedrooms    int    `json:"bedrooms"`
	Bathrooms   int    `json:"bathrooms"`
	Garage      int    `json:"garage"`
	Floorspace  int    `json:"floorspace"`
	Landarea    int    `json:"landarea"`
	PropertyID  uint   `json:"property_id"`
}
