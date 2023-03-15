package helper

import "gorm.io/gorm"

func JoinPortfolio(tx *gorm.DB) *gorm.DB {
	return tx.Joins("JOIN user_links ON user_links.user_id = users.id").
		Joins("JOIN portfolios ON user_links.portfolio_id = portfolios.id")
}
