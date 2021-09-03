package createpost

import (
	"context"
)

// Inport of CreatePost
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase CreatePost
type InportRequest struct {
	Title       string  `json:"title" binding:"required,min=5,max=20"`       //
	Description string  `json:"description" binding:"required,min=5,max=50"` //
	Content     string  `json:"content" binding:"required,json"`             //
	Cover       string  `json:"cover" binding:"required,url"`                //
	Slug        string  `json:"slug" binding:"required"`                     //
	Categories  []int64 `json:"categories" binding:"unique"`                 //
	Tags        []int64 `json:"tags" binding:"unique"`                       //
	UserID      int64   `json:"id_user" binding:"required,numeric"`          //
}

// InportResponse is response payload after running the usecase CreatePost
type InportResponse struct {
}

type CategoryRequest struct {
	ID int64 `json:"id_category" binding:"unique"`
}

type TagRequest struct {
	ID int64 `json:"id_tag" binding:"unique"`
}
