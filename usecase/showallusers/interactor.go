package showallusers

import (
	"context"

	"github.com/raismaulana/blogP/application/apperror"
	"github.com/raismaulana/blogP/domain/repository"
)

//go:generate mockery --name Outport -output mocks/

type showAllUsersInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase ShowAllUsers
func NewUsecase(outputPort Outport) Inport {
	return &showAllUsersInteractor{
		outport: outputPort,
	}
}

// Execute the usecase ShowAllUsers
func (r *showAllUsersInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	// code your usecase definition here ...
	err := repository.ReadOnly(ctx, r.outport, func(ctx context.Context) error {

		userObj, err := r.outport.FetchUsers(ctx)
		if err != nil {
			return err
		}
		if userObj == nil {
			return apperror.ObjectNotFound.Var(userObj)
		}

		for _, v := range userObj {
			res.Users = append(res.Users, UsersResponse{
				Username:   v.Username,
				Name:       v.Name,
				Email:      v.Email,
				City:       v.City,
				Country:    v.Country,
				Birthday:   v.Birthday,
				WebProfile: v.WebProfile,
			})
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}
