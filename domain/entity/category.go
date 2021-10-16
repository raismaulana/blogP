package entity

import (
	"time"

	"github.com/raismaulana/blogP/application/apperror"
)

type Category struct {
	ID        int64     `gorm:"primary_key:auto_increment;column:id_category"`            //
	Category  string    `gorm:"type:varchar(200);not null;uniqueIndex:tags_category_key"` //
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`                       //
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`                       //
}
type CategoryRequest struct {
	Category string `` //
}

func NewCategory(req CategoryRequest) (*Category, error) {
	var categoryObj Category
	if req.Category == "" {
		return nil, apperror.CategoryMustNotEmpty
	}
	categoryObj.Category = req.Category
	return &categoryObj, nil
}
