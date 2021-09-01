package activationuser

import (
	"context"
	"time"

	"github.com/raismaulana/blogP/application/apperror"
	"github.com/raismaulana/blogP/domain/repository"
	"gopkg.in/guregu/null.v4"
)

//go:generate mockery --name Outport -output mocks/

type activationUserInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase ActivationUser
func NewUsecase(outputPort Outport) Inport {
	return &activationUserInteractor{
		outport: outputPort,
	}
}

// Execute the usecase ActivationUser
func (r *activationUserInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	err := repository.WithTransaction(ctx, r.outport, func(ctx context.Context) error {
		userObj, err := r.outport.FindUserByID(ctx, req.ID)
		if err != nil {
			return apperror.ObjectNotFound.Var(userObj)
		}

		err = userObj.ValidateActivation(req.Email, req.ActivationCode)
		if err != nil {
			return err
		}

		key := userObj.RDBKeyActivationUser()
		RDBActivationCode, _ := r.outport.RDBGet(ctx, key)

		if RDBActivationCode != req.ActivationCode {
			return apperror.ActivationCodeIsIncorrectOrExpired
		}

		userObj.ActivatedAt = null.NewTime(time.Now(), true)

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
