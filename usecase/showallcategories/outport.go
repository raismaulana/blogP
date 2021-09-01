package showallcategories

import "github.com/raismaulana/blogP/domain/repository"

// Outport of ShowAllCategories
type Outport interface {
	repository.FetchCategoriesRepo
	repository.ReadOnlyDB
}
