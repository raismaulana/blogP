package rdbms

import (
	"context"

	"github.com/raismaulana/blogP/domain/entity"
	"github.com/raismaulana/blogP/infrastructure/database"
	"github.com/raismaulana/blogP/infrastructure/log"
)

func (r *RDBMSGateway) FetchTags(ctx context.Context) ([]*entity.Tag, error) {
	log.Info(ctx, "called")

	db, err := database.ExtractDB(ctx)
	if err != nil {
		return nil, err
	}

	var objs []*entity.Tag
	err = db.Find(&objs).Error
	if err != nil {
		log.Error(ctx, err.Error())
		return nil, err
	}
	return objs, nil
}

func (r *RDBMSGateway) SaveTag(ctx context.Context, obj *entity.Tag) error {
	log.Info(ctx, "called")
	db, err := database.ExtractDB(ctx)
	if err != nil {
		return err
	}

	err = db.Save(obj).Error
	if err != nil {
		log.Error(ctx, err.Error())
		return err
	}

	return nil
}

func (r *RDBMSGateway) FindTagByTag(ctx context.Context, tag string) (*entity.Tag, error) {
	log.Info(ctx, "called")
	db, err := database.ExtractDB(ctx)
	if err != nil {
		return nil, err
	}

	var tagObj entity.Tag
	err = db.Where("tag = ?", tag).First(&tagObj).Error
	if err != nil {
		log.Error(ctx, err.Error())
		return nil, err
	}

	return &tagObj, nil
}

func (r *RDBMSGateway) FindTagByID(ctx context.Context, id int64) (*entity.Tag, error) {
	log.Info(ctx, "called")
	db, err := database.ExtractDB(ctx)
	if err != nil {
		return nil, err
	}

	var tagObj entity.Tag
	err = db.Where("id_tag = ?", id).First(&tagObj).Error
	if err != nil {
		log.Error(ctx, err.Error())
		return nil, err
	}

	return &tagObj, nil
}

func (r *RDBMSGateway) DeleteTag(ctx context.Context, obj *entity.Tag) error {
	log.Info(ctx, "called")
	db, err := database.ExtractDB(ctx)
	if err != nil {
		return err
	}

	err = db.Delete(&obj).Error
	if err != nil {
		log.Error(ctx, err.Error())
		return err
	}
	return nil
}
