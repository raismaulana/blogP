package activationuser

import (
	"context"
)

// Inport of ActivationUser
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase ActivationUser
type InportRequest struct {
	ID             int64  `json:"id_user"`                                                   //
	Email          string `json:"email" form:"email" binding:"required"`                     //
	ActivationCode string `json:"activation_code" form:"activation_code" binding:"required"` //
}

// InportResponse is response payload after running the usecase ActivationUser
type InportResponse struct {
}
