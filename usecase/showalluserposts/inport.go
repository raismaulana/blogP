package showalluserposts

import (
	"context"
	"time"
)

// Inport of ShowAllUserPosts
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase ShowAllUserPosts
type InportRequest struct {
	Username string `uri:"username" binding:"required"`
}

// InportResponse is response payload after running the usecase ShowAllUserPosts
type InportResponse struct {
	Posts []PostResponse `json:"posts"` //
}

type PostResponse struct {
	ID          int64              `json:"id_post"`     //
	Title       string             `json:"title"`       //
	Description string             `json:"description"` //
	Cover       string             `json:"cover"`       //
	Slug        string             `json:"slug"`        //
	Categories  []CategoryResponse `json:"categories"`  //
	Tags        []TagResponse      `json:"tags"`        //
	UserID      int64              `json:"id_user"`     //
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
