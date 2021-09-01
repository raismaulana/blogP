package createcategory

import (
	"context"

	"github.com/raismaulana/blogP/application/apperror"
	"github.com/raismaulana/blogP/domain/entity"
	"github.com/raismaulana/blogP/domain/repository"
)

//go:generate mockery --name Outport -output mocks/

type createCategoryInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase CreateCategory
func NewUsecase(outputPort Outport) Inport {
	return &createCategoryInteractor{
		outport: outputPort,
	}
}

// Execute the usecase CreateCategory
func (r *createCategoryInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	err := repository.WithTransaction(ctx, r.outport, func(ctx context.Context) error {
		existingCategory, err := r.outport.FindCategoryByCategory(ctx, req.Category)
		if existingCategory != nil || err == nil {
			return apperror.CategoryAlreadyExsist
		}

		categoryObj, err := entity.NewCategory(entity.CategoryRequest{
			Category: req.Category,
		})
		if err != nil {
			return err
		}
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
