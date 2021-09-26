package entity

import (
	"time"

	"gorm.io/gorm"
)

type PostTag struct {
	PostID    int64     `gorm:"primary_key"`                        //
	TagID     int64     `gorm:"primary_key"`                        //
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"` //
}

func (PostTag) BeforeCreate(db *gorm.DB) error {
	return nil
}
