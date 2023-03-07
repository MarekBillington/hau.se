package helper

import "gorm.io/gorm"

func JoinPortfolio(db *gorm.DB, table interface{}, user_id uint, preload string) *gorm.DB {
	// Might need to make the preloading a []string for larger objects
	if preload != "" {
		db = db.Preload(preload)
	}

	// create view
	db.Joins("Portfolio")
	//Joins("UserLink")

	return db
}
