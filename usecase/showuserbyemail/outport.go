package showuserbyemail

import "github.com/raismaulana/blogP/domain/repository"

// Outport of ShowUserByEmail
type Outport interface {
	repository.FindUserByEmailRepo
	repository.ReadOnlyDB
}
