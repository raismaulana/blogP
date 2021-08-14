package createuser

import (
	"context"
	"time"

	"gopkg.in/guregu/null.v4"
)

// Inport of CreateUser
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase CreateUser
type InportRequest struct {
	Username   string      `json:"username" form:"username" binding:"required"` //
	Name       string      `json:"name" form:"name" binding:"required"`         //
	Email      string      `json:"email" form:"email" binding:"required,email"` //
	Password   string      `json:"password" form:"password" binding:"required"` //
	City       string      `json:"city" form:"city" binding:"required"`         //
	Country    string      `json:"country" form:"country" binding:"required"`   //
	Birthday   time.Time   `json:"birthday" form:"birthday" binding:"required"` //
	WebProfile null.String `json:"web_profile" form:"web_profile" `             //
}

// InportResponse is response payload after running the usecase CreateUser
type InportResponse struct {
}
