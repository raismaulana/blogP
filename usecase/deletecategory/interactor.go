package deletecategory

import (
	"context"

	"github.com/raismaulana/blogP/application/apperror"
	"github.com/raismaulana/blogP/domain/repository"
)

//go:generate mockery --name Outport -output mocks/

type deleteCategoryInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase DeleteCategory
func NewUsecase(outputPort Outport) Inport {
	return &deleteCategoryInteractor{
		outport: outputPort,
	}
}

// Execute the usecase DeleteCategory
func (r *deleteCategoryInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	err := repository.WithTransaction(ctx, r.outport, func(ctx context.Context) error {
		categoryObj, err := r.outport.FindCategoryByID(ctx, req.ID)
		if err != nil {
			return apperror.ObjectNotFound.Var(categoryObj)
		}

		err = r.outport.DeleteCategory(ctx, categoryObj)
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
