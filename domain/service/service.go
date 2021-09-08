package service

import (
	"context"

	"github.com/raismaulana/blogP/domain/entity"
)

type GetBaseURLRepo interface {
	GetBaseURL(ctx context.Context) string
}

type HashPasswordService interface {
	HashPassword(ctx context.Context, plainPassword string) (string, error)
}

type VerifyPasswordService interface {
	VerifyPassword(ctx context.Context, req VerifyPasswordServiceRequest) error
}

type VerifyPasswordServiceRequest struct {
	PlainPassword  string
	HashedPassword string
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
	BuildMailActivationAccount(ctx context.Context, req BuildMailActivationAccountServiceRequest) *BuildMailServiceResponse
}

type BuildMailActivationAccountServiceRequest struct {
	ID              int64  `` //
	To              string `` //
	Name            string `` //
	ActivationToken string `` //
}

type BuildMailServiceResponse struct {
	To      string `` //
	Subject string `` //
	Body    string `` //
}

type GenerateJWTTokenService interface {
	GenerateJWTToken(ctx context.Context, userObj entity.User) (string, error)
}

type GenerateJWTTokenServiceRequest struct {
	ID    string
	Email string
	Role  string
}
type BuildMailForgotPasswordAccountService interface {
	BuildMailForgotPasswordAccount(ctx context.Context, req BuildMailForgotPasswordAccountServiceRequest) *BuildMailServiceResponse
}

type BuildMailForgotPasswordAccountServiceRequest struct {
	ID                  int64  `` //
	To                  string `` //
	Username            string `` //
	ForgotPasswordToken string `` //
}
