package showuserbyusername

import "github.com/raismaulana/blogP/domain/repository"

// Outport of ShowUserByUsername
type Outport interface {
	repository.ReadOnlyDB
	repository.FindUserByUsernameRepo
}
