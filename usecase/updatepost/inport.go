package updatepost

import (
	"context"
)

// Inport of UpdatePost
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase UpdatePost
type InportRequest struct {
	ID          int64   `json:"id_post" binding:"required"`
	Title       string  `json:"title" binding:"required,min=5,max=155"`       //
	Description string  `json:"description" binding:"required,min=5,max=255"` //
	Content     string  `json:"content" binding:"required,json"`              //
	Cover       string  `json:"cover" binding:"required,url"`                 //
	Slug        string  `json:"slug" binding:"required,min=5,max=100"`        //
	Categories  []int64 `json:"categories" binding:"unique"`                  //
	Tags        []int64 `json:"tags" binding:"unique"`                        //
	UserID      int64   `json:"id_user" binding:"required"`                   //
	Role        string  `json:"role" binding:"required"`                      //
}

// InportResponse is response payload after running the usecase UpdatePost
type InportResponse struct {
}
