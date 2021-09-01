package showallcategories

import (
	"context"
)

// Inport of ShowAllCategories
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase ShowAllCategories
type InportRequest struct {
}

// InportResponse is response payload after running the usecase ShowAllCategories
type InportResponse struct {
	Categories []CategoryResponse `json:"categories"` //
}

type CategoryResponse struct {
	ID       int64  `json:"id_category"` //
	Category string `json:"category"`    //
}
