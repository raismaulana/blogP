package showuserbyid

import (
	"context"
	"time"

	"gopkg.in/guregu/null.v4"
)

// Inport of ShowUserByID
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase ShowUserByID
type InportRequest struct {
	ID int64 `json:"id_user"` //
}

// InportResponse is response payload after running the usecase ShowUserByID
type InportResponse struct {
	ID         int64       `json:"id_user"`     //
	Username   string      `json:"username"`    //
	Name       string      `json:"name"`        //
	Email      string      `json:"email"`       //
	City       string      `json:"city"`        //
	Country    string      `json:"country"`     //
	Birthday   time.Time   `json:"birthday"`    //
	WebProfile null.String `json:"web_profile"` //
	Activated  bool        `json:"activated"`   //
}
