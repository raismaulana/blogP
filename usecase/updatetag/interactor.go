package updatetag

import (
	"context"

	"github.com/raismaulana/blogP/application/apperror"
	"github.com/raismaulana/blogP/domain/repository"
)

//go:generate mockery --name Outport -output mocks/

type updateTagInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase UpdateTag
func NewUsecase(outputPort Outport) Inport {
	return &updateTagInteractor{
		outport: outputPort,
	}
}

// Execute the usecase UpdateTag
func (r *updateTagInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}
	err := repository.ReadOnly(ctx, r.outport, func(ctx context.Context) error {
		tagObj, err := r.outport.FindTagByTag(ctx, req.Tag)
		if tagObj != nil || err == nil {
			return apperror.TagAlreadyExsist
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	err = repository.WithTransaction(ctx, r.outport, func(ctx context.Context) error {
		tagObj, err := r.outport.FindTagByID(ctx, req.ID)
		if err != nil {
			return apperror.ObjectNotFound.Var(tagObj)
		}

		tagObj.Tag = req.Tag
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
