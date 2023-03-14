package database

import (
	"hause/entity"

	"gorm.io/gorm"
)

func SetupDb(db *gorm.DB) {
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Country{})
	db.AutoMigrate(&entity.Portfolio{})
	db.AutoMigrate(&entity.UserLink{})
	db.AutoMigrate(&entity.House{})
}
