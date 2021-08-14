package deleteuser

import "github.com/raismaulana/blogP/domain/repository"

// Outport of DeleteUser
type Outport interface {
	repository.DeleteUserRepo
	repository.FindUserByIDRepo
	repository.TransactionDB
}
