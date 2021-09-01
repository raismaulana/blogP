package updatecategory

import (
	"context"

	"github.com/raismaulana/blogP/application/apperror"
	"github.com/raismaulana/blogP/domain/repository"
)

//go:generate mockery --name Outport -output mocks/

type updateCategoryInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase UpdateCategory
func NewUsecase(outputPort Outport) Inport {
	return &updateCategoryInteractor{
		outport: outputPort,
	}
}

// Execute the usecase UpdateCategory
func (r *updateCategoryInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}
	err := repository.ReadOnly(ctx, r.outport, func(ctx context.Context) error {
		categoryObj, err := r.outport.FindCategoryByCategory(ctx, req.Category)
		if categoryObj != nil || err == nil {
			return apperror.CategoryAlreadyExsist
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	err = repository.WithTransaction(ctx, r.outport, func(ctx context.Context) error {
		categoryObj, err := r.outport.FindCategoryByID(ctx, req.ID)
		if err != nil {
			return apperror.ObjectNotFound.Var(categoryObj)
		}

		categoryObj.Category = req.Category
		err = r.outport.SaveCategory(ctx, categoryObj)
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
