package showuserbyemail

import (
	"github.com/raismaulana/blogP/domain/repository"
	"github.com/raismaulana/blogP/domain/service"
)

// Outport of ShowUserByEmail
type Outport interface {
	repository.FindUserByEmailRepo
	repository.ReadOnlyDB
	service.GetBaseURLRepo
}
