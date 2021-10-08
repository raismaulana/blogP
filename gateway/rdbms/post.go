package rdbms

import (
	"context"

	"github.com/raismaulana/blogP/domain/entity"
	"github.com/raismaulana/blogP/infrastructure/database"
	"github.com/raismaulana/blogP/infrastructure/log"
)

func (r *RDBMSGateway) FetchPosts(ctx context.Context, paginate *database.PaginateRequest) ([]*entity.Post, error) {
	log.Info(ctx, "called")

	db, err := database.ExtractDB(ctx)
	if err != nil {
		return nil, err
	}

	var objs []*entity.Post
	key, val := entity.PostSortableColumn()(paginate.Sort)
	paginate.Sort = key
	err = db.
		Scopes(database.Paginate(paginate, db, objs)).
		Preload("User").
		Preload("Categories").
		Preload("Tags").
		Order(val).
		Find(&objs).
		Error

	if err != nil {
		log.Error(ctx, err.Error())
		return nil, err
	}
	return objs, nil
}

func (r *RDBMSGateway) CreatePost(ctx context.Context, obj *entity.Post) error {
	log.Info(ctx, "called")
	db, err := database.ExtractDB(ctx)
	if err != nil {
		return err
	}

	err = db.Omit("Categories.*,Tags.*").Save(obj).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *RDBMSGateway) SavePost(ctx context.Context, obj *entity.Post) error {
	log.Info(ctx, "called")
	db, err := database.ExtractDB(ctx)
	if err != nil {
		return err
	}

	err = db.Model(&obj).Association("Categories").Replace(obj.Categories)
	if err != nil {
		return err
	}
	err = db.Model(&obj).Association("Tags").Replace(obj.Tags)
	if err != nil {
		return err
	}
	err = db.Omit("Categories.*,Tags.*").Save(obj).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *RDBMSGateway) FindPostBySlug(ctx context.Context, slug string) (*entity.Post, error) {
	log.Info(ctx, "called")
	db, err := database.ExtractDB(ctx)
	if err != nil {
		return nil, err
	}

	var postObj entity.Post
	err = db.Where("slug = ?", slug).Preload("User").Preload("Categories").Preload("Tags").First(&postObj).Error
	if err != nil {
		log.Error(ctx, err.Error())
		return nil, err
	}

	return &postObj, nil
}

func (r *RDBMSGateway) FindPostByID(ctx context.Context, id int64) (*entity.Post, error) {
	log.Info(ctx, "called")
	db, err := database.ExtractDB(ctx)
	if err != nil {
		return nil, err
	}

	var postObj entity.Post
	err = db.Where("id_post = ?", id).Preload("User").Preload("Categories").Preload("Tags").First(&postObj).Error
	if err != nil {
		log.Error(ctx, err.Error())
		return nil, err
	}

	return &postObj, nil
}

func (r *RDBMSGateway) DeletePost(ctx context.Context, obj *entity.Post) error {
	log.Info(ctx, "called")
	db, err := database.ExtractDB(ctx)
	if err != nil {
		return err
	}
	err = db.Model(&obj).Association("Categories").Replace(obj.Categories)
	if err != nil {
		log.Error(ctx, err.Error())
		return err
	}
	err = db.Model(&obj).Association("Tags").Replace(obj.Tags)
	if err != nil {
		log.Error(ctx, err.Error())
		return err
	}
	err = db.Delete(&obj).Error
	if err != nil {
		log.Error(ctx, err.Error())
		return err
	}
	return nil
}
