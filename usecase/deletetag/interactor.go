package deletetag

import (
	"context"

	"github.com/raismaulana/blogP/application/apperror"
	"github.com/raismaulana/blogP/domain/repository"
)

//go:generate mockery --name Outport -output mocks/

type deleteTagInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase DeleteTag
func NewUsecase(outputPort Outport) Inport {
	return &deleteTagInteractor{
		outport: outputPort,
	}
}

// Execute the usecase DeleteTag
func (r *deleteTagInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	err := repository.WithTransaction(ctx, r.outport, func(ctx context.Context) error {
		tagObj, err := r.outport.FindTagByID(ctx, req.ID)
		if err != nil {
			return apperror.ObjectNotFound.Var(tagObj)
		}

		err = r.outport.DeleteTag(ctx, tagObj)
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
