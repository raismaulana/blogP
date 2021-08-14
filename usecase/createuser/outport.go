package createuser

import (
	"github.com/raismaulana/blogP/domain/repository"
	"github.com/raismaulana/blogP/domain/service"
)

// Outport of CreateUser
type Outport interface {
	repository.FindUserByUsernameRepo
	repository.FindUserByEmailRepo
	repository.ReadOnlyDB
	repository.SaveUserRepo
	repository.TransactionDB
	service.HashPasswordService
}
