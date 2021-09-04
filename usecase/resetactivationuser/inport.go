package resetactivationuser

import (
	"context"
)

// Inport of ResetActivationUser
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase ResetActivationUser
type InportRequest struct {
	ID int64 `json:"id_user" binding:"required"` //
}

// InportResponse is response payload after running the usecase ResetActivationUser
type InportResponse struct {
}
