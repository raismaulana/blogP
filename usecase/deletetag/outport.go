package deletetag

import "github.com/raismaulana/blogP/domain/repository"

// Outport of DeleteTag
type Outport interface {
	repository.DeleteTagRepo
	repository.FindTagByIDRepo
	repository.TransactionDB
}
