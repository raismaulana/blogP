package updatecategory

import (
	"context"
)

// Inport of UpdateCategory
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase UpdateCategory
type InportRequest struct {
	ID       int64  `json:"id_tag" binding:"required"`                //
	Category string `json:"category" binding:"required,min=1,max=15"` //
}

// InportResponse is response payload after running the usecase UpdateCategory
type InportResponse struct {
}
