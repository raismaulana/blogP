package entity

import (
	"time"

	"gorm.io/gorm"
)

type PostCategory struct {
	PostID     int64     `gorm:"primary_key"` //
	CategoryID int64     `gorm:"primary_key"` //
	CreatedAt  time.Time ``                   //
}

func (PostCategory) BeforeCreate(db *gorm.DB) error {
	return nil
}
