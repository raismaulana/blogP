package createcategory

import (
	"context"
)

// Inport of CreateCategory
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase CreateCategory
type InportRequest struct {
	Category string `json:"category" binding:"required,min=1,max=50"`
}

// InportResponse is response payload after running the usecase CreateCategory
type InportResponse struct {
}
