package updatepost

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

type updatePostInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase UpdatePost
func NewUsecase(outputPort Outport) Inport {
	return &updatePostInteractor{
		outport: outputPort,
	}
}

// Execute the usecase UpdatePost
func (r *updatePostInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	var postObj *entity.Post

	err := repository.ReadOnly(ctx, r.outport, func(ctx context.Context) error {

		postIDObj, err := r.outport.FindPostByID(ctx, req.ID)
		if err != nil {
			return apperror.ObjectNotFound.Var(postObj)
		}
		if postIDObj.UserID != req.UserID && req.Role != "king" {
			return apperror.Forbidden
		}

		postSlugObj, err := r.outport.FindPostBySlug(ctx, req.Slug)
		if postIDObj.Slug != req.Slug && (postSlugObj != nil || err == nil) {
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

		postObj = postIDObj
		return nil
	})
	if err != nil {
		return nil, err
	}

	err = repository.WithTransaction(ctx, r.outport, func(ctx context.Context) error {
		var vCategory []entity.Category
		for _, v := range req.Categories {
			vCategory = append(vCategory, entity.Category{
				ID: v,
			})
		}
		var vTag []entity.Tag
		for _, v := range req.Tags {
			vTag = append(vTag, entity.Tag{
				ID: v,
			})
		}
		err = postObj.UpdatePost(entity.UpdatePostRequest{
			Title:       req.Title,
			Description: req.Description,
			Content:     datatypes.JSON([]byte(req.Content)),
			Cover:       postObj.Cover,
			Slug:        req.Slug,
			Categories:  vCategory,
			Tags:        vTag,
		})
		if err != nil {
			return err
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
