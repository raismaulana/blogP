package showtagbyid

import (
	"context"
)

// Inport of ShowTagByID
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase ShowTagByID
type InportRequest struct {
	ID int64 `json:"id_tag" binding:"required"` //
}

// InportResponse is response payload after running the usecase ShowTagByID
type InportResponse struct {
	ID  int64  `json:"id_tag"`
	Tag string `json:"tag"`
}
