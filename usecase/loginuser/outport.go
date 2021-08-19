package loginuser

import (
	"github.com/raismaulana/blogP/domain/repository"
	"github.com/raismaulana/blogP/domain/service"
)

// Outport of LoginUser
type Outport interface {
	repository.FindUserByUsernameRepo
	repository.ReadOnlyDB
	service.VerifyPasswordService
	service.GenerateJWTTokenService
}
