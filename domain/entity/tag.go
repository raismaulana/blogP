package entity

import (
	"strings"
	"time"

	"github.com/raismaulana/blogP/application/apperror"
)

type Tag struct {
	ID        int64     `gorm:"primary_key:auto_increment;column:id_tag"` //
	Tag       string    `gorm:"type: varchar(200) not null unique"`       //
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`       //
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`       //
}
type TagRequest struct {
	Tag string `` //
}

func NewTag(req TagRequest) (*Tag, error) {

	//validate
	if strings.TrimSpace(req.Tag) == "" || len(strings.TrimSpace(req.Tag)) > 200 {
		return nil, apperror.TagMustNotEmpty
	}

	obj := Tag{
		Tag: req.Tag,
	}

	return &obj, nil
}
