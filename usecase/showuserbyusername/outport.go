package showuserbyusername

import (
	"github.com/raismaulana/blogP/domain/repository"
	"github.com/raismaulana/blogP/domain/service"
)

// Outport of ShowUserByUsername
type Outport interface {
	repository.ReadOnlyDB
	repository.FindUserByUsernameRepo
	service.GetBaseURLRepo
}
