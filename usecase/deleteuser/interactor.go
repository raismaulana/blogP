package deleteuser

import (
	"context"

	"github.com/raismaulana/blogP/domain/repository"
)

//go:generate mockery --name Outport -output mocks/

type deleteUserInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase DeleteUser
func NewUsecase(outputPort Outport) Inport {
	return &deleteUserInteractor{
		outport: outputPort,
	}
}

// Execute the usecase DeleteUser
func (r *deleteUserInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	// code your usecase definition here ...
	err := repository.WithTransaction(ctx, r.outport, func(ctx context.Context) error {
		userObj, err := r.outport.FindUserByID(ctx, req.ID)
		if err != nil {
			return err
		}

		err = r.outport.DeleteUser(ctx, userObj)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	return res, nil
}
