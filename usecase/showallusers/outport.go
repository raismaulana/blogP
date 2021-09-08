package showallusers

import (
	"github.com/raismaulana/blogP/domain/repository"
	"github.com/raismaulana/blogP/domain/service"
)

// Outport of ShowAllUsers
type Outport interface {
	repository.FetchUsersRepo
	repository.ReadOnlyDB
	service.GetBaseURLRepo
}
