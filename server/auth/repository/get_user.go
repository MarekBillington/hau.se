package repository

import (
	"hause/entity"

	"gorm.io/gorm"
)

func GetUserFromEmail(email string, db *gorm.DB) entity.User {

	var user entity.User

	db.Preload("UserRole").First(&user, "email = ?", email)

	return user
}
