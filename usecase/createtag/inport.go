package createtag

import (
	"context"
)

// Inport of CreateTag
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase CreateTag
type InportRequest struct {
	Tag string `json:"tag" binding:"required,min=1,max=15"` //
}

// InportResponse is response payload after running the usecase CreateTag
type InportResponse struct {
}
