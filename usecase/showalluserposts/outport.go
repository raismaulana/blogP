package showalluserposts

import "github.com/raismaulana/blogP/domain/repository"

// Outport of ShowAllUserPosts
type Outport interface {
	repository.FetchPostsByUserUsernameRepo
	repository.FindUserByUsernameRepo
	repository.ReadOnlyDB
}
