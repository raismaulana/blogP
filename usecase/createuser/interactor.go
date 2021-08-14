package createuser

import (
	"context"

	"github.com/raismaulana/blogP/application/apperror"
	"github.com/raismaulana/blogP/domain/entity"
	"github.com/raismaulana/blogP/domain/repository"
)

//go:generate mockery --name Outport -output mocks/

type createUserInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase CreateUser
func NewUsecase(outputPort Outport) Inport {
	return &createUserInteractor{
		outport: outputPort,
	}
}

// Execute the usecase CreateUser
func (r *createUserInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	err := repository.ReadOnly(ctx, r.outport, func(ctx context.Context) error {
		existingUser, err := r.outport.FindUserByUsername(ctx, req.Username)
		if existingUser != nil || err == nil {
			return apperror.UsernameAlreadyUsed
		}

		existingUser, err = r.outport.FindUserByEmail(ctx, req.Email)
		if existingUser != nil || err == nil {
			return apperror.EmailAlreadyUsed
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	err = repository.WithTransaction(ctx, r.outport, func(ctx context.Context) error {

		hashedPassword, err := r.outport.HashPassword(ctx, req.Password)
		if err != nil {
			return err
		}

		userObj, err := entity.NewUser(entity.UserRequest{
			Username:   req.Username,
			Name:       req.Name,
			Email:      req.Email,
			Password:   hashedPassword,
			City:       req.City,
			Country:    req.Country,
			Birthday:   req.Birthday,
			WebProfile: req.WebProfile,
		})
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
