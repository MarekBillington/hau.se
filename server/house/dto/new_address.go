package dto

type NewAddress struct {
	Street1   string `json:"street_1" binding:"required"`
	Street2   string `json:"street_2"`
	Postcode  string `json:"postcode" binding:"required"`
	Town      string `json:"town" binding:"required"`
	CountryID uint   `json:"country_id" binding:"required"`
}
