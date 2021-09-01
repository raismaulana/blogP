package deletecategory

import "github.com/raismaulana/blogP/domain/repository"

// Outport of DeleteCategory
type Outport interface {
	repository.DeleteCategoryRepo
	repository.FindCategoryByIDRepo
	repository.TransactionDB
}
