package showuserbyemail

import (
	"context"
	"time"

	"gopkg.in/guregu/null.v4"
)

// Inport of ShowUserByEmail
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase ShowUserByEmail
type InportRequest struct {
	Email string `uri:"email" binding:"required,email"` //
}

// InportResponse is response payload after running the usecase ShowUserByEmail
type InportResponse struct {
	ID           int64       `json:"id_user"`       //
	Username     string      `json:"username"`      //
	Name         string      `json:"name"`          //
	Email        string      `json:"email"`         //
	City         string      `json:"city"`          //
	Country      string      `json:"country"`       //
	Birthday     time.Time   `json:"birthday"`      //
	WebProfile   null.String `json:"web_profile"`   //
	PhotoProfile string      `json:"photo_profile"` //
	Activated    bool        `json:"activated"`     //
}
