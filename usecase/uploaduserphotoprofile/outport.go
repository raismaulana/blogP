package uploaduserphotoprofile

import (
	"github.com/raismaulana/blogP/domain/repository"
	"github.com/raismaulana/blogP/domain/service"
)

// Outport of UploadUserPhotoProfile
type Outport interface {
	service.GenerateRandomStringService
	repository.FindUserByIDRepo
	repository.SaveUserRepo
	repository.TransactionDB
	repository.ReadOnlyDB
}
