package showallposts

import "github.com/raismaulana/blogP/domain/repository"

// Outport of ShowAllPosts
type Outport interface {
	repository.ReadOnlyDB
	repository.FetchPostsRepo
}
