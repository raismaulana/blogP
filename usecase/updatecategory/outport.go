package updatecategory

import "github.com/raismaulana/blogP/domain/repository"

// Outport of UpdateCategory
type Outport interface {
	repository.FindCategoryByCategoryRepo
	repository.FindCategoryByIDRepo
	repository.ReadOnlyDB
	repository.SaveCategoryRepo
	repository.TransactionDB
}
