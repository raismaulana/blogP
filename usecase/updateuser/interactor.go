package updateuser

import (
	"context"

	"github.com/raismaulana/blogP/domain/entity"
	"github.com/raismaulana/blogP/domain/repository"
	"github.com/raismaulana/blogP/infrastructure/log"
	"github.com/raismaulana/blogP/infrastructure/util"
)

//go:generate mockery --name Outport -output mocks/

type updateUserInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase UpdateUser
func NewUsecase(outputPort Outport) Inport {
	return &updateUserInteractor{
		outport: outputPort,
	}
}

// Execute the usecase UpdateUser
func (r *updateUserInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	// code your usecase definition here ...
	err := repository.WithTransaction(ctx, r.outport, func(ctx context.Context) error {
		userObj, err := r.outport.FindUserByID(ctx, req.ID)
		if err != nil {
			return err
		}

		err = userObj.UpdateUser(entity.UserUpdateRequest{
			Name:       req.Name,
			City:       req.City,
			Country:    req.Country,
			Birthday:   req.Birthday,
			WebProfile: req.WebProfile,
		})
		log.Info(ctx, util.MustJSON(userObj))
		if err != nil {
			return err
		}

		err = r.outport.SaveUser(ctx, userObj)
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
