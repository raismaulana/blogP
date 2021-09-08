package showuserbyid

import (
	"github.com/raismaulana/blogP/domain/repository"
	"github.com/raismaulana/blogP/domain/service"
)

// Outport of ShowUserByID
type Outport interface {
	repository.FindUserByIDRepo
	repository.ReadOnlyDB
	service.GetBaseURLRepo
}
