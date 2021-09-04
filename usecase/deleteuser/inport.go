package deleteuser

import (
	"context"
)

// Inport of DeleteUser
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase DeleteUser
type InportRequest struct {
	ID int64 `json:"id_user" binding:"required"` //
}

// InportResponse is response payload after running the usecase DeleteUser
type InportResponse struct {
}
