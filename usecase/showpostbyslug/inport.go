package showpostbyslug

import (
	"context"
	"encoding/json"
	"time"
)

// Inport of ShowPostBySlug
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase ShowPostBySlug
type InportRequest struct {
	Slug string `json:"slug" uri:"slug" binding:"required"` //
}

// InportResponse is response payload after running the usecase ShowPostBySlug
type InportResponse struct {
	ID          int64              `json:"id_post"`     //
	Title       string             `json:"title"`       //
	Description string             `json:"description"` //
	Content     json.RawMessage    `json:"content"`     //
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