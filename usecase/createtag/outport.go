package createtag

import "github.com/raismaulana/blogP/domain/repository"

// Outport of CreateTag
type Outport interface {
	repository.TransactionDB
	repository.SaveTagRepo
	repository.FindTagByTagRepo
}
