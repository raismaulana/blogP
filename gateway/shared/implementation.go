package shared

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/raismaulana/blogP/domain/service"
	"github.com/raismaulana/blogP/infrastructure/envconfig"
	"github.com/raismaulana/blogP/infrastructure/log"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

type SharedGateway struct {
	Env *envconfig.EnvConfig
}

func NewSharedGateway(env *envconfig.EnvConfig) *SharedGateway {
	return &SharedGateway{
		Env: env,
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

func (r *SharedGateway) SendMail(ctx context.Context, req service.SendMailServiceRequest) {
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
		return
	}
	log.Info(ctx, "Mail sent")
}

func (r *SharedGateway) BuildMailActivationAccount(ctx context.Context, req service.BuildMailActivationAccountServiceRequest) *service.BuildMailActivationAccountServiceResponse {
	log.Info(ctx, "called")

	var mail service.BuildMailActivationAccountServiceResponse
	mail.To = req.To
	mail.Subject = "Account Activation"
	mail.Body = fmt.Sprintf("<p>Hello %s, your activation code is %s or click link below </p><p><a href=\"%susers/%v/activation?email=%s&activation_code=%s\">click me.</a></p><p>This link will expire in 3 days.</p>",
		req.Name,
		req.ActivationToken,
		r.Env.AppBaseURL,
		req.ID,
		req.To,
		req.ActivationToken,
	)

	return &mail
}
