package repository

import (
	"hause/auth/dto"
	"hause/entity"

	"gorm.io/gorm"
)

func CreateUser(reg dto.Register, db *gorm.DB) (entity.User, error) {

	user := entity.User{
		Email:     reg.Email,
		Password:  reg.Password,
		FirstName: reg.FirstName,
		LastName:  reg.LastName,
	}

	err := db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
