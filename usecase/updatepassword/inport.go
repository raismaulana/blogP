package updatepassword

import (
	"context"
)

// Inport of UpdatePassword
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase UpdatePassword
type InportRequest struct {
	ID          int64  `uri:"id_user" binding:"required"`
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" bbinding:"required,min=6,max=255"`
}

// InportResponse is response payload after running the usecase UpdatePassword
type InportResponse struct {
}
