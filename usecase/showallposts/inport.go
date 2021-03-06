package showallposts

import (
	"context"
	"time"

	"github.com/raismaulana/blogP/infrastructure/database"
)

// Inport of ShowAllPosts
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase ShowAllPosts
type InportRequest struct {
	PaginateRequest database.PaginateRequest `binding:"dive"`
}

// InportResponse is response payload after running the usecase ShowAllPosts
type InportResponse struct {
	database.PaginateRequest ``             //
	Posts                    []PostResponse `json:"posts"` //
}

type PostResponse struct {
	ID          int64              `json:"id_post"`     //
	Title       string             `json:"title"`       //
	Description string             `json:"description"` //
	Cover       string             `json:"cover"`       //
	Slug        string             `json:"slug"`        //
	ViewCount   int64              `json:"view_count"`  //
	Categories  []CategoryResponse `json:"categories"`  //
	Tags        []TagResponse      `json:"tags"`        //
	UserID      int64              `json:"id_user"`     //
	AuthorName  string             `json:"author_name"` //
	CreatedAt   time.Time          `json:"created_at"`  //
	UpdatedAt   time.Time          `json:"updated_at"`  //
}
type CategoryResponse struct {
	ID       int64  `json:"id_category"` //
	Category string `json:"category"`
}

type TagResponse struct {
	ID  int64  `json:"id_tag"` //
	Tag string `json:"tag"`
}
