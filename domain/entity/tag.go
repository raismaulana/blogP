package entity

import (
	"time"

	"gorm.io/gorm"
)

type Tag struct {
	ID        int64          `gorm:"primary_key:auto_increment;column:id_tag"` //
	Tag       string         `gorm:"type: varchar(200) not null unique"`
	CreatedAt time.Time      ``             //
	UpdatedAt time.Time      ``             //
	DeletedAt gorm.DeletedAt `gorm:"index"` //
}
