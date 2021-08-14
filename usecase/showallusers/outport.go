package showallusers

import "github.com/raismaulana/blogP/domain/repository"

// Outport of ShowAllUsers
type Outport interface {
	repository.FetchUsersRepo
	repository.ReadOnlyDB
}
