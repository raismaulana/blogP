package updatepassword

import (
	"context"

	"github.com/raismaulana/blogP/application/apperror"
	"github.com/raismaulana/blogP/domain/repository"
	"github.com/raismaulana/blogP/domain/service"
)

//go:generate mockery --name Outport -output mocks/

type updatePasswordInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase UpdatePassword
func NewUsecase(outputPort Outport) Inport {
	return &updatePasswordInteractor{
		outport: outputPort,
	}
}

// Execute the usecase UpdatePassword
func (r *updatePasswordInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	err := repository.WithTransaction(ctx, r.outport, func(ctx context.Context) error {
		userObj, err := r.outport.FindUserByID(ctx, req.ID)
		if err != nil {
			return apperror.ObjectNotFound.Var(userObj)
		}
		err = r.outport.VerifyPassword(ctx, service.VerifyPasswordServiceRequest{
			PlainPassword:  req.OldPassword,
			HashedPassword: userObj.Password,
		})
		if err != nil {
			return apperror.InvalidCredential
		}
		newHashedPassword, err := r.outport.HashPassword(ctx, req.NewPassword)
		if err != nil {
			return err
		}
		err = userObj.ChangePassword(newHashedPassword)
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
