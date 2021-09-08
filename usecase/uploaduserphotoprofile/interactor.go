package uploaduserphotoprofile

import (
	"context"
	"io"
	"os"

	"github.com/raismaulana/blogP/application/apperror"
	"github.com/raismaulana/blogP/domain/entity"
	"github.com/raismaulana/blogP/domain/repository"
)

//go:generate mockery --name Outport -output mocks/

type uploadUserPhotoProfileInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase UploadUserPhotoProfile
func NewUsecase(outputPort Outport) Inport {
	return &uploadUserPhotoProfileInteractor{
		outport: outputPort,
	}
}

// Execute the usecase UploadUserPhotoProfile
func (r *uploadUserPhotoProfileInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}
	var userObj *entity.User
	err := repository.ReadOnly(ctx, r.outport, func(ctx context.Context) error {
		obj, err := r.outport.FindUserByID(ctx, req.ID)
		userObj = obj
		if err != nil {
			return apperror.ObjectNotFound.Var(userObj)
		}
		return nil
	})
	if err != nil {
		return nil, apperror.ServerError.Var(err)
	}
	src, err := req.PhotoProfile.Open()
	if err != nil {
		return nil, apperror.ServerError.Var(err)
	}
	defer src.Close()
	filename := r.outport.GenerateRandomString(ctx)
	path := "/public/images/" + filename + ".jpg"
	out, err := os.Create("." + path)
	if err != nil {
		return nil, apperror.ServerError.Var(err)
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	if err != nil {
		return nil, apperror.ServerError.Var(err)
	}

	err = repository.WithTransaction(ctx, r.outport, func(ctx context.Context) error {
		userObj.PhotoProfile = path
		err = r.outport.SaveUser(ctx, userObj)
		if err != nil {
			return apperror.ServerError.Var(err)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}
