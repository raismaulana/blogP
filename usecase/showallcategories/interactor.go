package showallcategories

import (
	"context"

	"github.com/raismaulana/blogP/domain/repository"
)

//go:generate mockery --name Outport -output mocks/

type showAllCategoriesInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase ShowAllCategories
func NewUsecase(outputPort Outport) Inport {
	return &showAllCategoriesInteractor{
		outport: outputPort,
	}
}

// Execute the usecase ShowAllCategories
func (r *showAllCategoriesInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	err := repository.ReadOnly(ctx, r.outport, func(ctx context.Context) error {
		categoryObjs, err := r.outport.FetchCategories(ctx)
		if err != nil {
			return err
		}

		for _, v := range categoryObjs {
			res.Categories = append(res.Categories, CategoryResponse{
				ID:       v.ID,
				Category: v.Category,
			})
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}
