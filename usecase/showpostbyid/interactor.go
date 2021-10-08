package showpostbyid

import (
	"context"
	"encoding/json"

	"github.com/raismaulana/blogP/application/apperror"
	"github.com/raismaulana/blogP/domain/repository"
)

//go:generate mockery --name Outport -output mocks/

type showPostByIDInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase ShowPostByID
func NewUsecase(outputPort Outport) Inport {
	return &showPostByIDInteractor{
		outport: outputPort,
	}
}

// Execute the usecase ShowPostByID
func (r *showPostByIDInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}
	err := repository.WithTransaction(ctx, r.outport, func(ctx context.Context) error {
		postObj, err := r.outport.FindPostByID(ctx, req.ID)
		if err != nil {
			return apperror.ObjectNotFound.Var(postObj)
		}

		postObj.UpdateViewCount()
		err = r.outport.SavePost(ctx, postObj)
		if err != nil {
			return err
		}

		var vCategories []CategoryResponse
		for _, v := range postObj.Categories {
			vCategories = append(vCategories, CategoryResponse{
				ID:       v.ID,
				Category: v.Category,
			})
		}
		var vTags []TagResponse
		for _, v := range postObj.Tags {
			vTags = append(vTags, TagResponse{
				ID:  v.ID,
				Tag: v.Tag,
			})
		}
		res = &InportResponse{
			ID:          postObj.ID,
			Title:       postObj.Title,
			Description: postObj.Description,
			Content:     json.RawMessage(postObj.Content),
			Cover:       postObj.Cover,
			Slug:        postObj.Slug,
			ViewCount:   postObj.ViewCount,
			Categories:  vCategories,
			Tags:        vTags,
			UserID:      postObj.UserID,
			AuthorName:  postObj.User.Name,
			CreatedAt:   postObj.CreatedAt,
			UpdatedAt:   postObj.UpdatedAt,
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
