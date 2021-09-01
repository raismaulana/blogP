package createcategory

import "github.com/raismaulana/blogP/domain/repository"

// Outport of CreateCategory
type Outport interface {
	repository.FindCategoryByCategoryRepo
	repository.TransactionDB
	repository.SaveCategoryRepo
}
