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
	ID         int64       `json:"id_user"`     //
	Name       string      `json:"name"`        //
	City       string      `json:"city"`        //
	Country    string      `json:"country"`     //
	Birthday   time.Time   `json:"birthday"`    //
	WebProfile null.String `json:"web_profile"` //
}

// InportResponse is response payload after running the usecase UpdateUser
type InportResponse struct {
}
