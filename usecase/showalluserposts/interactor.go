package showalluserposts

import (
	"context"

	"github.com/raismaulana/blogP/application/apperror"
	"github.com/raismaulana/blogP/domain/repository"
)

//go:generate mockery --name Outport -output mocks/

type showAllUserPostsInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase ShowAllUserPosts
func NewUsecase(outputPort Outport) Inport {
	return &showAllUserPostsInteractor{
		outport: outputPort,
	}
}

// Execute the usecase ShowAllUserPosts
func (r *showAllUserPostsInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	err := repository.ReadOnly(ctx, r.outport, func(ctx context.Context) error {
		userObj, err := r.outport.FetchPostsByUserUsername(ctx, req.Username)
		if err != nil {
			return apperror.ObjectNotFound.Var(userObj)
		}
		for _, v := range userObj.Posts {
			var vCategories []CategoryResponse
			for _, w := range v.Categories {
				vCategories = append(vCategories, CategoryResponse{
					ID:       w.ID,
					Category: w.Category,
				})
			}

			var vTags []TagResponse
			for _, w := range v.Tags {
				vTags = append(vTags, TagResponse{
					ID:  w.ID,
					Tag: w.Tag,
				})
			}
			res.Posts = append(res.Posts, PostResponse{
				ID:          v.ID,
				Title:       v.Title,
				Description: v.Description,
				Cover:       v.Cover,
				Slug:        v.Slug,
				ViewCount:   v.ViewCount,
				Categories:  vCategories,
				Tags:        vTags,
				UserID:      v.UserID,
				CreatedAt:   v.CreatedAt,
				UpdatedAt:   v.UpdatedAt,
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
