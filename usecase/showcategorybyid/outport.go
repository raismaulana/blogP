package showcategorybyid

import "github.com/raismaulana/blogP/domain/repository"

// Outport of ShowCategoryByID
type Outport interface {
	repository.FindCategoryByIDRepo
	repository.ReadOnlyDB
}
