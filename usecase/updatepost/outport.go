package updatepost

import "github.com/raismaulana/blogP/domain/repository"

// Outport of UpdatePost
type Outport interface {
	repository.FindCategoriesByIDsRepo
	repository.FindTagsByIDsRepo
	repository.FindPostByIDRepo
	repository.FindPostBySlugRepo
	repository.ReadOnlyDB
	repository.SavePostRepo
	repository.TransactionDB
}
