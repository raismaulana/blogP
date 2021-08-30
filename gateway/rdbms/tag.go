package rdbms

import (
	"context"

	"github.com/raismaulana/blogP/domain/entity"
	"github.com/raismaulana/blogP/infrastructure/database"
	"github.com/raismaulana/blogP/infrastructure/log"
)

func (r *RDBMSGateway) FetchTags(ctx context.Context, scope bool) ([]*entity.Tag, error) {
	log.Info(ctx, "called")

	db, err := database.ExtractDB(ctx)
	if err != nil {
		return nil, err
	}

	var objs []*entity.Tag
	if scope {
		err = db.Find(&objs).Error
	} else {
		err = db.Unscoped().Find(&objs).Error
	}
	if err != nil {
		log.Error(ctx, err.Error())
		return nil, err
	}
	return objs, nil
}
