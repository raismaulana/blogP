package showpostbyslug

import "github.com/raismaulana/blogP/domain/repository"

// Outport of ShowPostBySlug
type Outport interface {
	repository.FindPostBySlugRepo
	repository.ReadOnlyDB
}
