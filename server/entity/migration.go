package entity

import "gorm.io/gorm"

func InitDb(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Country{})
	db.AutoMigrate(&Portfolio{})
	db.AutoMigrate(&UserLink{})
	db.AutoMigrate(&House{})
	db.AutoMigrate(&Address{})
}
