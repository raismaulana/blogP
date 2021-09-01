package createtag

import (
	"context"

	"github.com/raismaulana/blogP/application/apperror"
	"github.com/raismaulana/blogP/domain/entity"
	"github.com/raismaulana/blogP/domain/repository"
)

//go:generate mockery --name Outport -output mocks/

type createTagInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase CreateTag
func NewUsecase(outputPort Outport) Inport {
	return &createTagInteractor{
		outport: outputPort,
	}
}

// Execute the usecase CreateTag
func (r *createTagInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	err := repository.WithTransaction(ctx, r.outport, func(ctx context.Context) error {
		existingTag, err := r.outport.FindTagByTag(ctx, req.Tag)
		if existingTag != nil || err == nil {
			return apperror.TagAlreadyExsist
		}

		tagObj, err := entity.NewTag(entity.TagRequest{
			Tag: req.Tag,
		})
		if err != nil {
			return err
		}

		err = r.outport.SaveTag(ctx, tagObj)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}
