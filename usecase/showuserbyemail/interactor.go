package showuserbyemail

import (
	"context"

	"github.com/raismaulana/blogP/application/apperror"
	"github.com/raismaulana/blogP/domain/repository"
)

//go:generate mockery --name Outport -output mocks/

type showUserByEmailInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase ShowUserByEmail
func NewUsecase(outputPort Outport) Inport {
	return &showUserByEmailInteractor{
		outport: outputPort,
	}
}

// Execute the usecase ShowUserByEmail
func (r *showUserByEmailInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	// code your usecase definition here ...
	err := repository.ReadOnly(ctx, r.outport, func(ctx context.Context) error {
		userObj, err := r.outport.FindUserByEmail(ctx, req.Email)
		if err != nil {
			return apperror.ObjectNotFound.Var(userObj)
		}

		res = &InportResponse{
			ID:           userObj.ID,
			Username:     userObj.Username,
			Name:         userObj.Name,
			Email:        userObj.Email,
			City:         userObj.City,
			Country:      userObj.Country,
			Birthday:     userObj.Birthday,
			WebProfile:   userObj.WebProfile,
			PhotoProfile: r.outport.GetBaseURL(ctx) + userObj.PhotoProfile,
			Activated:    userObj.ActivatedAt.Valid,
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}
