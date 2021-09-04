package showalltags

import (
	"context"
)

// Inport of ShowAllTags
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase ShowAllTags
type InportRequest struct {
}

// InportResponse is response payload after running the usecase ShowAllTags
type InportResponse struct {
	Tags []Tag `json:"tags"` //
}

type Tag struct {
	ID  int64  `json:"id_tag"` //
	Tag string `json:"name"`   //
}
