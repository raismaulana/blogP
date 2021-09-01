package deletecategory

import (
	"context"
)

// Inport of DeleteCategory
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase DeleteCategory
type InportRequest struct {
	ID int64 `json:"id_category"`
}

// InportResponse is response payload after running the usecase DeleteCategory
type InportResponse struct {
}
