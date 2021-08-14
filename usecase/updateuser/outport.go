package updateuser

import "github.com/raismaulana/blogP/domain/repository"

// Outport of UpdateUser
type Outport interface {
	repository.FindUserByIDRepo
	repository.SaveUserRepo
	repository.TransactionDB
}
