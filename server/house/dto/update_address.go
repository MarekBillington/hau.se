package dto

type UpdateAddress struct {
	ID        uint   `json:"id" binding:"required"`
	Street1   string `json:"street_1"`
	Street2   string `json:"street_2"`
	Postcode  string `json:"postcode"`
	Town      string `json:"town"`
	CountryID uint   `json:"country_id"`
}
