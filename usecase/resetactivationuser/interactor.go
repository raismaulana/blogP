package resetactivationuser

import (
	"context"
	"time"

	"github.com/raismaulana/blogP/application/apperror"
	"github.com/raismaulana/blogP/domain/repository"
	"github.com/raismaulana/blogP/domain/service"
	"github.com/raismaulana/blogP/infrastructure/log"
)

//go:generate mockery --name Outport -output mocks/

type resetActivationUserInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase ResetActivationUser
func NewUsecase(outputPort Outport) Inport {
	return &resetActivationUserInteractor{
		outport: outputPort,
	}
}

// Execute the usecase ResetActivationUser
func (r *resetActivationUserInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}
	mail := &service.BuildMailActivationAccountServiceResponse{}

	err := repository.ReadOnly(ctx, r.outport, func(ctx context.Context) error {
		userObj, err := r.outport.FindUserByID(ctx, req.ID, true)
		if err != nil {
			return apperror.ObjectNotFound.Var(userObj)
		}
		if userObj.ActivatedAt.Valid {
			return apperror.UserIsAlreadyActivated
		}

		RDBkey := userObj.RDBKeyActivationUser()
		RDBvalue := r.outport.GenerateRandomString(ctx)

		err = r.outport.RDBSet(ctx, RDBkey, RDBvalue, time.Hour*72)
		if err != nil {
			log.Error(ctx, err.Error())
		}

		mail = r.outport.BuildMailActivationAccount(ctx, service.BuildMailActivationAccountServiceRequest{
			ID:              userObj.ID,
			To:              userObj.Email,
			Name:            userObj.Name,
			ActivationToken: RDBvalue,
		})

		return nil
	})

	if err != nil {
		return nil, err
	}

	go r.outport.SendMail(ctx, service.SendMailServiceRequest{
		To:      mail.To,
		Subject: mail.Subject,
		Body:    mail.Body,
	})

	return res, nil
}
