package updatepassword

import (
	"github.com/raismaulana/blogP/domain/repository"
	"github.com/raismaulana/blogP/domain/service"
)

// Outport of UpdatePassword
type Outport interface {
	repository.FindUserByIDRepo
	repository.SaveUserRepo
	repository.TransactionDB
	service.HashPasswordService
	service.VerifyPasswordService
}
