package entity

import (
	"time"

	"github.com/raismaulana/blogP/application/apperror"
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
	CreatedAt   time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP"`        //
	UpdatedAt   time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP"`        //
}

type PostRequest struct {
	Title       string         `` //
	Description string         `` //
	Content     datatypes.JSON `` //
	Cover       string         `` //
	Slug        string         `` //
	Categories  []Category     `` //
	Tags        []Tag          `` //
	UserID      int64          `` //
}
type UpdatePostRequest struct {
	Title       string         `` //
	Description string         `` //
	Content     datatypes.JSON `` //
	Cover       string         `` //
	Slug        string         `` //
	Categories  []Category     `` //
	Tags        []Tag          `` //
}

func NewPost(req PostRequest) (*Post, error) {
	if req.Title == "" {
		return nil, apperror.TitleMustNotEmpty
	}
	if req.Description == "" {
		return nil, apperror.DescriptionMustNotEmpty
	}
	if _, err := req.Content.MarshalJSON(); err != nil {
		return nil, apperror.ContentMustBeValidJSON
	}
	if req.Cover == "" {
		return nil, apperror.CoverMustNotEmpty
	}
	if req.Slug == "" {
		return nil, apperror.SlugMustNotEmpty
	}
	if req.UserID == 0 {
		return nil, apperror.AuthorIDMustNotEmpty
	}
	obj := Post{
		Title:       req.Title,
		Description: req.Description,
		Content:     req.Content,
		Cover:       req.Cover,
		Slug:        req.Slug,
		Categories:  req.Categories,
		Tags:        req.Tags,
		UserID:      req.UserID,
	}
	return &obj, nil
}

func (r *Post) UpdatePost(req UpdatePostRequest) error {
	if req.Title == "" {
		return apperror.TitleMustNotEmpty
	}
	if req.Description == "" {
		return apperror.DescriptionMustNotEmpty
	}
	if _, err := req.Content.MarshalJSON(); err != nil {
		return apperror.ContentMustBeValidJSON
	}
	if req.Cover == "" {
		return apperror.CoverMustNotEmpty
	}
	if req.Slug == "" {
		return apperror.SlugMustNotEmpty
	}

	r.Title = req.Title
	r.Description = req.Description
	r.Content = req.Content
	r.Cover = req.Cover
	r.Slug = req.Slug
	r.Categories = req.Categories
	r.Tags = req.Tags

	return nil
}
func (r *Post) DeleteAssociation() {
	r.Categories = []Category{}
	r.Tags = []Tag{}
}
