package showuserbyusername

import (
	"context"
	"time"

	"gopkg.in/guregu/null.v4"
)

// Inport of ShowUserByUsername
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase ShowUserByUsername
type InportRequest struct {
	Username string `json:"username"` //
}

// InportResponse is response payload after running the usecase ShowUserByUsername
type InportResponse struct {
	Username   string      `json:"username"`    //
	Name       string      `json:"name"`        //
	Email      string      `json:"email"`       //
	City       string      `json:"city"`        //
	Country    string      `json:"country"`     //
	Birthday   time.Time   `json:"birthday"`    //
	WebProfile null.String `json:"web_profile"` //
}
