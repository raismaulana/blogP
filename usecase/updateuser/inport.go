package updateuser

import (
	"context"
	"time"

	"gopkg.in/guregu/null.v4"
)

// Inport of UpdateUser
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase UpdateUser
type InportRequest struct {
	ID         int64       `json:"id_user" binding:"required,numeric"`              //
	Name       string      `json:"name" binding:"required,min=4,max=25"`            //
	City       string      `json:"city" binding:"required,min=1,max=25"`            //
	Country    string      `json:"country" binding:"required,min=1,max=25"`         //
	Birthday   time.Time   `json:"birthday" form:"birthday" binding:"required,lte"` //
	WebProfile null.String `json:"web_profile" binding:"omitempty,url"`             //
}

// InportResponse is response payload after running the usecase UpdateUser
type InportResponse struct {
}
