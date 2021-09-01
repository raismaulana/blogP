package showcategorybyid

import (
	"context"
)

// Inport of ShowCategoryByID
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase ShowCategoryByID
type InportRequest struct {
	ID int64 `json:"id_category"`
}

// InportResponse is response payload after running the usecase ShowCategoryByID
type InportResponse struct {
	ID       int64  `json:"id_category"`
	Category string `json:"category"`
}
