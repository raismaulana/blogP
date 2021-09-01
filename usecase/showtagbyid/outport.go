package showtagbyid

import "github.com/raismaulana/blogP/domain/repository"

// Outport of ShowTagByID
type Outport interface {
	repository.FindTagByIDRepo
	repository.ReadOnlyDB
}
