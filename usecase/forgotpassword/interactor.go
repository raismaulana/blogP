package forgotpassword

import (
	"context"
	"time"

	"github.com/raismaulana/blogP/application/apperror"
	"github.com/raismaulana/blogP/domain/repository"
	"github.com/raismaulana/blogP/domain/service"
	"github.com/raismaulana/blogP/infrastructure/log"
)

//go:generate mockery --name Outport -output mocks/

type forgetPasswordInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase ForgetPassword
func NewUsecase(outputPort Outport) Inport {
	return &forgetPasswordInteractor{
		outport: outputPort,
	}
}

// Execute the usecase ForgetPassword
func (r *forgetPasswordInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	mail := &service.BuildMailServiceResponse{}
	err := repository.ReadOnly(ctx, r.outport, func(ctx context.Context) error {
		userObj, err := r.outport.FindUserByEmail(ctx, req.Email)
		if err != nil {
			return apperror.ObjectNotFound.Var(userObj)
		}
		RDBkey := userObj.RDBKeyForgotPassword()
		RDBvalue := r.outport.GenerateRandomString(ctx)

		err = r.outport.RDBSet(ctx, RDBkey, RDBvalue, time.Hour)
		if err != nil {
			log.Error(ctx, err.Error())
		}

		mail = r.outport.BuildMailForgotPasswordAccount(ctx, service.BuildMailForgotPasswordAccountServiceRequest{
			ID:                  userObj.ID,
			To:                  userObj.Email,
			Username:            userObj.Username,
			ForgotPasswordToken: RDBvalue,
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
