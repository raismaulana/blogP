package rdbms

import (
	"context"

	"github.com/raismaulana/blogP/domain/entity"
	"github.com/raismaulana/blogP/infrastructure/database"
	"github.com/raismaulana/blogP/infrastructure/log"
)

func (r *RDBMSGateway) FetchCategories(ctx context.Context) ([]*entity.Category, error) {
	log.Info(ctx, "called")

	db, err := database.ExtractDB(ctx)
	if err != nil {
		return nil, err
	}

	var objs []*entity.Category
	err = db.Find(&objs).Error
	if err != nil {
		log.Error(ctx, err.Error())
		return nil, err
	}
	return objs, nil
}

func (r *RDBMSGateway) SaveCategory(ctx context.Context, obj *entity.Category) error {
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

func (r *RDBMSGateway) FindCategoryByCategory(ctx context.Context, category string) (*entity.Category, error) {
	log.Info(ctx, "called")
	db, err := database.ExtractDB(ctx)
	if err != nil {
		return nil, err
	}

	var obj entity.Category
	err = db.Where("category = ?", category).First(&obj).Error
	if err != nil {
		log.Error(ctx, err.Error())
		return nil, err
	}

	return &obj, nil
}

func (r *RDBMSGateway) FindCategoryByID(ctx context.Context, id int64) (*entity.Category, error) {
	log.Info(ctx, "called")
	db, err := database.ExtractDB(ctx)
	if err != nil {
		return nil, err
	}

	var obj entity.Category
	err = db.Where("id_category = ?", id).First(&obj).Error
	if err != nil {
		log.Error(ctx, err.Error())
		return nil, err
	}

	return &obj, nil
}

func (r *RDBMSGateway) DeleteCategory(ctx context.Context, obj *entity.Category) error {
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

func (r *RDBMSGateway) FindCategoriesByIDs(ctx context.Context, ids []int64) ([]*entity.Category, error) {
	log.Info(ctx, "called")
	db, err := database.ExtractDB(ctx)
	if err != nil {
		return nil, err
	}
	var objs []*entity.Category
	err = db.Where(ids).Find(&objs).Error
	if err != nil {
		return nil, err
	}
	return objs, nil
}
