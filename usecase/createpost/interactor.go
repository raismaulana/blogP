package createpost

import (
	"context"

	"github.com/raismaulana/blogP/application/apperror"
	"github.com/raismaulana/blogP/domain/entity"
	"github.com/raismaulana/blogP/domain/repository"
	"github.com/raismaulana/blogP/infrastructure/log"
	"github.com/raismaulana/blogP/infrastructure/util"
	"gorm.io/datatypes"
)

//go:generate mockery --name Outport -output mocks/

type createPostInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase CreatePost
func NewUsecase(outputPort Outport) Inport {
	return &createPostInteractor{
		outport: outputPort,
	}
}

// Execute the usecase CreatePost
func (r *createPostInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}
	err := repository.ReadOnly(ctx, r.outport, func(ctx context.Context) error {
		postObj, err := r.outport.FindPostBySlug(ctx, req.Slug)
		if postObj != nil || err == nil {
			return apperror.SlugAlreadyExsist
		}
		if req.Categories != nil {
			categoryObjs, err := r.outport.FindCategoriesByIDs(ctx, req.Categories)
			if err != nil || len(categoryObjs) != len(req.Categories) {
				log.Info(ctx, util.MustJSON(categoryObjs))
				return apperror.SomeCategoryDoesNotExist
			}
		}
		if req.Tags != nil {
			tagObjs, err := r.outport.FindTagsByIDs(ctx, req.Tags)
			if err != nil || len(tagObjs) != len(req.Tags) {
				log.Info(ctx, util.MustJSON(tagObjs))
				return apperror.SomeTagDoesNotExist
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	err = repository.WithTransaction(ctx, r.outport, func(ctx context.Context) error {
		postObj, err := entity.NewPost(entity.PostRequest{
			Title:       req.Title,
			Description: req.Description,
			Content:     datatypes.JSON([]byte(req.Content)),
			Cover:       req.Cover,
			Slug:        req.Slug,
			Categories:  []entity.Category{},
			Tags:        []entity.Tag{},
			UserID:      req.UserID,
		})
		if err != nil {
			return err
		}
		for _, v := range req.Categories {
			postObj.Categories = append(postObj.Categories, entity.Category{
				ID: v,
			})
		}
		for _, v := range req.Tags {
			postObj.Tags = append(postObj.Tags, entity.Tag{
				ID: v,
			})
		}

		err = r.outport.SavePost(ctx, postObj)
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
