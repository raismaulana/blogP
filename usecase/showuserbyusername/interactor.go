package showuserbyusername

import (
	"context"

	"github.com/raismaulana/blogP/application/apperror"
	"github.com/raismaulana/blogP/domain/repository"
)

//go:generate mockery --name Outport -output mocks/

type showUserByUsernameInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase ShowUserByUsername
func NewUsecase(outputPort Outport) Inport {
	return &showUserByUsernameInteractor{
		outport: outputPort,
	}
}

// Execute the usecase ShowUserByUsername
func (r *showUserByUsernameInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	// code your usecase definition here ...
	err := repository.ReadOnly(ctx, r.outport, func(ctx context.Context) error {
		userObj, err := r.outport.FindUserByUsername(ctx, req.Username, true)
		if err != nil {
			return apperror.ObjectNotFound.Var(userObj)
		}

		res = &InportResponse{
			Username:   userObj.Username,
			Name:       userObj.Name,
			Email:      userObj.Email,
			City:       userObj.City,
			Country:    userObj.Country,
			Birthday:   userObj.Birthday,
			WebProfile: userObj.WebProfile,
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}
