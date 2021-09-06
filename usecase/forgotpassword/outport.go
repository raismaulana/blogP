package forgotpassword

import (
	"github.com/raismaulana/blogP/domain/repository"
	"github.com/raismaulana/blogP/domain/service"
)

// Outport of ForgetPassword
type Outport interface {
	repository.RDBSetRepo
	repository.FindUserByEmailRepo
	repository.ReadOnlyDB
	service.BuildMailForgotPasswordAccountService
	service.GenerateRandomStringService
	service.SendMailService
}
