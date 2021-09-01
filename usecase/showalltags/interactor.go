package showalltags

import (
	"context"

	"github.com/raismaulana/blogP/application/apperror"
	"github.com/raismaulana/blogP/domain/repository"
)

//go:generate mockery --name Outport -output mocks/

type showAllTagsInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase ShowAllTags
func NewUsecase(outputPort Outport) Inport {
	return &showAllTagsInteractor{
		outport: outputPort,
	}
}

// Execute the usecase ShowAllTags
func (r *showAllTagsInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	err := repository.ReadOnly(ctx, r.outport, func(ctx context.Context) error {
		tagObjs, err := r.outport.FetchTags(ctx)

		if err != nil {
			return apperror.ObjectNotFound.Var(tagObjs)
		}

		for _, v := range tagObjs {
			res.Tags = append(res.Tags, Tag{
				ID:  v.ID,
				Tag: v.Tag,
			})
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}
