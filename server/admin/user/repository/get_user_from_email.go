package repository

import (
	"hause/entity"

	"gorm.io/gorm"
)

func GetUserFromEmail(email string, db *gorm.DB) entity.User {

	var user entity.User

	db.First(&user, "email = ?", email)

	return user
}
