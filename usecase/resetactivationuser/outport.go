package resetactivationuser

import (
	"github.com/raismaulana/blogP/domain/repository"
	"github.com/raismaulana/blogP/domain/service"
)

// Outport of ResetActivationUser
type Outport interface {
	repository.RDBSetRepo
	repository.RDBGetRepo
	repository.ReadOnlyDB
	repository.FindUserByIDRepo
	service.GenerateRandomStringService
	service.BuildMailActivationAccountService
	service.SendMailService
}
