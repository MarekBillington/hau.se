package entity

import (
	"time"

	"gorm.io/datatypes"
)

type UserRole struct {
	ID          uint           `json:"id"`
	Active      bool           `json:"active" gorm:"notNull;default:true"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"update_at"`
	Name        string         `json:"name"`
	Permissions datatypes.JSON `json:"permissions"`
}
