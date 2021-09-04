package deletetag

import (
	"context"
)

// Inport of DeleteTag
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase DeleteTag
type InportRequest struct {
	ID int64 `json:"id_tag" binding:"required"` //
}

// InportResponse is response payload after running the usecase DeleteTag
type InportResponse struct {
}
