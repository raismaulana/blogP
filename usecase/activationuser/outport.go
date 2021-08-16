package activationuser

import "github.com/raismaulana/blogP/domain/repository"

// Outport of ActivationUser
type Outport interface {
	repository.FindUserByIDRepo
	repository.ReadOnlyDB
	repository.SaveUserRepo
	repository.TransactionDB
}
