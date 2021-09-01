package showtagbyid

import (
	"context"

	"github.com/raismaulana/blogP/application/apperror"
	"github.com/raismaulana/blogP/domain/repository"
)

//go:generate mockery --name Outport -output mocks/

type showTagByIDInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase ShowTagByID
func NewUsecase(outputPort Outport) Inport {
	return &showTagByIDInteractor{
		outport: outputPort,
	}
}

// Execute the usecase ShowTagByID
func (r *showTagByIDInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	err := repository.ReadOnly(ctx, r.outport, func(ctx context.Context) error {
		tagObj, err := r.outport.FindTagByID(ctx, req.ID)
		if err != nil {
			return apperror.ObjectNotFound.Var(tagObj)
		}

		res.ID = tagObj.ID
		res.Tag = tagObj.Tag

		return nil
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}
