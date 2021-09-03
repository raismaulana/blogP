package shared

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/raismaulana/blogP/domain/entity"
	"github.com/raismaulana/blogP/domain/service"
	"github.com/raismaulana/blogP/infrastructure/auth"
	"github.com/raismaulana/blogP/infrastructure/envconfig"
	"github.com/raismaulana/blogP/infrastructure/log"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

type SharedGateway struct {
	Env      *envconfig.EnvConfig
	JWTToken *auth.JWTToken
}

func NewSharedGateway(env *envconfig.EnvConfig, jwtToken *auth.JWTToken) *SharedGateway {
	return &SharedGateway{
		Env:      env,
		JWTToken: jwtToken,
	}
}

func (r *SharedGateway) HashPassword(ctx context.Context, plainPassword string) (string, error) {
	log.Info(ctx, "called")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), 10)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (r *SharedGateway) GenerateRandomString(ctx context.Context) string {
	log.Info(ctx, "called")

	return uuid.NewString()
}

func (r *SharedGateway) SendMail(ctx context.Context, req service.SendMailServiceRequest) error {
	log.Info(ctx, "called")

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", r.Env.SMTPSender)
	mailer.SetHeader("To", req.To)
	mailer.SetHeader("Subject", req.Subject)
	mailer.SetBody("text/html", req.Body)

	dialer := gomail.NewDialer(
		r.Env.SMTPHost,
		r.Env.SMTPPort,
		r.Env.SMTPEmail,
		r.Env.SMTPPassword,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Error(ctx, err.Error())
		return err
	}
	log.Info(ctx, "Mail sent")

	return nil
}

func (r *SharedGateway) BuildMailActivationAccount(ctx context.Context, req service.BuildMailActivationAccountServiceRequest) *service.BuildMailActivationAccountServiceResponse {
	log.Info(ctx, "called")

	var mail service.BuildMailActivationAccountServiceResponse
	mail.To = req.To
	mail.Subject = "Account Activation"
	mail.Body = fmt.Sprintf("<p>Hello %s, your activation code is %s or click link below </p><p><a href=\"%susers/%v/activation?email=%s&activation_code=%s\">click me.</a></p><p>This link will expire in 3 days.</p>",
		req.Name,
		req.ActivationToken,
		r.Env.AppBaseURLV1,
		req.ID,
		req.To,
		req.ActivationToken,
	)

	return &mail
}

func (r *SharedGateway) VerifyPassword(ctx context.Context, req service.VerifyPasswordServiceRequest) error {
	log.Info(ctx, "called")

	err := bcrypt.CompareHashAndPassword([]byte(req.HashedPassword), []byte(req.PlainPassword))
	if err != nil {
		return err
	}

	return nil
}

func (r *SharedGateway) GenerateJWTToken(ctx context.Context, userObj entity.User) (string, error) {
	log.Info(ctx, "called")

	token, err := r.JWTToken.GenerateToken(auth.GenerateTokenRequest{
		Subject:   userObj.Name,
		ID:        userObj.ID,
		Email:     userObj.Email,
		Activated: userObj.ActivatedAt.Valid,
		Role:      "",
	})
	if err != nil {
		log.Error(ctx, err.Error())
		return "", err
	}
	return token, nil
}
