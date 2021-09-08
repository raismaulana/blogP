package showallusers

import (
	"context"
	"time"

	"gopkg.in/guregu/null.v4"
)

// Inport of ShowAllUsers
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase ShowAllUsers
type InportRequest struct {
}

// InportResponse is response payload after running the usecase ShowAllUsers
type InportResponse struct {
	Users []UsersResponse `json:"users"` //
}

type UsersResponse struct {
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
