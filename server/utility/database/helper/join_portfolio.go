package helper

import "gorm.io/gorm"

func JoinPortfolio(db *gorm.DB, table interface{}, user_id uint, preload string) *gorm.DB {
	// Might need to make the preloading a []string for larger objects
	if preload != "" {
		db = db.Preload(preload)
	}

	// @todo create view for simpler joining
	db = db.Joins("Portfolio").
		Joins("JOIN user_links ON user_links.portfolio_id = \"Portfolio\".id AND user_links.user_id = ?", user_id)

	return db
}
