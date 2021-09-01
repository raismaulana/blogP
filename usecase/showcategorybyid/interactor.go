package showcategorybyid

import (
	"context"

	"github.com/raismaulana/blogP/application/apperror"
	"github.com/raismaulana/blogP/domain/repository"
)

//go:generate mockery --name Outport -output mocks/

type showCategoryByIDInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase ShowCategoryByID
func NewUsecase(outputPort Outport) Inport {
	return &showCategoryByIDInteractor{
		outport: outputPort,
	}
}

// Execute the usecase ShowCategoryByID
func (r *showCategoryByIDInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	err := repository.ReadOnly(ctx, r.outport, func(ctx context.Context) error {
		categoryObj, err := r.outport.FindCategoryByID(ctx, req.ID)
		if err != nil {
			return apperror.ObjectNotFound.Var(categoryObj)
		}

		res.ID = categoryObj.ID
		res.Category = categoryObj.Category

		return nil
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}
