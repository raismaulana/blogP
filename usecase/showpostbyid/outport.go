package showpostbyid

import "github.com/raismaulana/blogP/domain/repository"

// Outport of ShowPostByID
type Outport interface {
	repository.FindPostByIDRepo
	repository.ReadOnlyDB
}
