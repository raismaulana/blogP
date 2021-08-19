package loginuser

import (
	"context"
	"strconv"

	"github.com/raismaulana/blogP/application/apperror"
	"github.com/raismaulana/blogP/domain/repository"
	"github.com/raismaulana/blogP/domain/service"
)

//go:generate mockery --name Outport -output mocks/

type loginUserInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase LoginUser
func NewUsecase(outputPort Outport) Inport {
	return &loginUserInteractor{
		outport: outputPort,
	}
}

// Execute the usecase LoginUser
func (r *loginUserInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	err := repository.ReadOnly(ctx, r.outport, func(ctx context.Context) error {
		userObj, err := r.outport.FindUserByUsername(ctx, req.Username, true)
		if err != nil {
			return apperror.InvalidCredential
		}

		err = r.outport.VerifyPassword(ctx, service.VerifyPasswordServiceRequest{
			PlainPassword:  req.Password,
			HashedPassword: userObj.Password,
		})
		if err != nil {
			return apperror.InvalidCredential
		}

		token, err := r.outport.GenerateJWTToken(ctx, service.GenerateJWTTokenServiceRequest{
			ID:    strconv.FormatInt(userObj.ID, 10),
			Email: userObj.Email,
			Role:  "",
		})
		if err != nil {
			return apperror.FailedGenerateAuthToken
		}

		res.Token = token

		return nil
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}
