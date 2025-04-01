package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID uint `gorm:"primaryKey"` // Redeclared to directly access ID field in repositories

	TelegramID int64   `gorm:"uniqueIndex;not null"`
	Username   *string `gorm:"size:32"` // Pointer type (*string) because it can be nil

	Status  int8 `gorm:"default:0;not null"` // -1 - banned, 0 - blocked, 1 - active
	IsAdmin bool `gorm:"default:false"`
}
