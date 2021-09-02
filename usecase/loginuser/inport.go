package loginuser

import (
	"context"
)

// Inport of LoginUser
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase LoginUser
type InportRequest struct {
	Username string `json:"username" binding:"required"` //
	Password string `json:"password" binding:"required"` //
}

// InportResponse is response payload after running the usecase LoginUser
type InportResponse struct {
	Token string `json:"token"` //
}
