package service

import "context"

type HashPasswordService interface {
	HashPassword(ctx context.Context, plainPassword string) (string, error)
}
type SendMailService interface {
	SendMail(ctx context.Context, req SendMailServiceRequest) error
}

type SendMailServiceRequest struct {
	To      string `` //
	Subject string `` //
	Body    string `` //
}

type GenerateRandomStringService interface {
	GenerateRandomString(ctx context.Context) string
}
type BuildMailActivationAccountService interface {
	BuildMailActivationAccount(ctx context.Context, req BuildMailActivationAccountServiceRequest) *BuildMailActivationAccountServiceResponse
}

type BuildMailActivationAccountServiceRequest struct {
	ID              int64  `` //
	To              string `` //
	Name            string `` //
	ActivationToken string `` //
}

type BuildMailActivationAccountServiceResponse struct {
	To      string `` //
	Subject string `` //
	Body    string `` //
}
