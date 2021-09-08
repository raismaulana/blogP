package uploaduserphotoprofile

import (
	"context"
	"mime/multipart"
)

// Inport of UploadUserPhotoProfile
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase UploadUserPhotoProfile
type InportRequest struct {
	ID           int64                 `uri:"id_user" binding:"required"`
	PhotoProfile *multipart.FileHeader `form:"photo_profile" binding:"required"` //
}

// InportResponse is response payload after running the usecase UploadUserPhotoProfile
type InportResponse struct {
}
