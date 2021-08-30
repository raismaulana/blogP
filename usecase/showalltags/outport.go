package showalltags

import "github.com/raismaulana/blogP/domain/repository"

// Outport of ShowAllTags
type Outport interface {
	repository.FetchTagsRepo
	repository.ReadOnlyDB
}
