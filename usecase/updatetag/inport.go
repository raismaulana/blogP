package updatetag

import (
	"context"
)

// Inport of UpdateTag
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase UpdateTag
type InportRequest struct {
	ID  int64  `json:"id_tag" binding:"required,numeric"`   //
	Tag string `json:"tag" binding:"required,min=1,max=15"` //
}

// InportResponse is response payload after running the usecase UpdateTag
type InportResponse struct {
}
