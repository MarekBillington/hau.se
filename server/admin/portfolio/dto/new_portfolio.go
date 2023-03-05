package dto

type NewPortfolio struct {
	Name      string `json:"name" binding:"required"`
	UserID    uint   `json:"user_id" binding:"required"`
	CountryID uint   `json:"country_id" binding:"required"`
}
