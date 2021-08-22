package createuser

import (
	"context"
	"time"
)

// Inport of CreateUser
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase CreateUser
type InportRequest struct {
	Username   string    `json:"username" form:"username" binding:"required,min=6,max=12" `  //
	Name       string    `json:"name" form:"name" binding:"required,min=4,max=25" `          //
	Email      string    `json:"email" form:"email" binding:"required,email" `               //
	Password   string    `json:"password" form:"password" binding:"required,min=6,max=255" ` //
	City       string    `json:"city" form:"city" binding:"required,min=1,max=25"`           //
	Country    string    `json:"country" form:"country" binding:"required,min=1,max=25" `    //
	Birthday   time.Time `json:"birthday" form:"birthday" binding:"required,lte" `           //
	WebProfile *string   `json:"web_profile" form:"web_profile" binding:"omitempty,url"`     //
}

// InportResponse is response payload after running the usecase CreateUser
type InportResponse struct {
}
