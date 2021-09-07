package deletepost

import (
	"context"
)

// Inport of DeletePost
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase DeletePost
type InportRequest struct {
	ID     int64  `json:"id_post" binding:"required"` //
	UserID int64  `json:"id_user" binding:"required"` //
	Role   string `json:"role" binding:"required"`    //
}

// InportResponse is response payload after running the usecase DeletePost
type InportResponse struct {
}
