package entity

import (
	"time"

	"gorm.io/datatypes"
)

type Post struct {
	ID          int64          `gorm:"primary_key:auto_increment;column:id_post"` //
	Title       string         `gorm:"type:varchar(100) not null"`                //
	Description string         `gorm:"type:varchar(255) not null"`                //
	Content     datatypes.JSON `gorm:"not null"`                                  //
	Cover       string         `gorm:"type:text not null"`                        //
	Slug        string         `gorm:"type:varchar(255) not null unique"`         //
	Categories  []Category     `gorm:"many2many:post_categories;"`                //
	Tags        []Tag          `gorm:"many2many:post_tags;"`                      //
	UserID      int64          `gorm:"not null"`                                  //
	CreatedAt   time.Time      ``                                                 //
	UpdatedAt   time.Time      ``                                                 //
}
