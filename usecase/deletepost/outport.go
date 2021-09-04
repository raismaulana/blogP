package deletepost

import "github.com/raismaulana/blogP/domain/repository"

// Outport of DeletePost
type Outport interface {
	repository.FindPostByIDRepo
	repository.DeletePostRepo
	repository.TransactionDB
}
