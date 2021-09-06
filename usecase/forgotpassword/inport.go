package forgotpassword

import (
	"context"
)

// Inport of ForgetPassword
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase ForgetPassword
type InportRequest struct {
	Email string `json:"email" binding:"required,email"`
}

// InportResponse is response payload after running the usecase ForgetPassword
type InportResponse struct {
}
