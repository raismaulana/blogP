package updatetag

import "github.com/raismaulana/blogP/domain/repository"

// Outport of UpdateTag
type Outport interface {
	repository.FindTagByIDRepo
	repository.FindTagByTagRepo
	repository.ReadOnlyDB
	repository.SaveTagRepo
	repository.TransactionDB
}
