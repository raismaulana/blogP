package createpost

import "github.com/raismaulana/blogP/domain/repository"

// Outport of CreatePost
type Outport interface {
	repository.FindCategoriesByIDsRepo
	repository.FindTagsByIDsRepo
	repository.FindPostBySlugRepo
	repository.ReadOnlyDB
	repository.SavePostRepo
	repository.TransactionDB
}
