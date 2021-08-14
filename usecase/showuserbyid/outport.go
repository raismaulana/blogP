package showuserbyid

import "github.com/raismaulana/blogP/domain/repository"

// Outport of ShowUserByID
type Outport interface {
	repository.FindUserByIDRepo
	repository.ReadOnlyDB
}
