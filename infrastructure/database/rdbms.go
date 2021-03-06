package database

import (
	"context"
	"math"

	"github.com/raismaulana/blogP/application/apperror"
	"github.com/raismaulana/blogP/infrastructure/log"
	"gorm.io/gorm"
)

type contextDBType string

var ContextDBValue contextDBType = "DB"

type GormTransactionImpl struct {
	DB *gorm.DB
}

func (r *GormTransactionImpl) BeginTransaction(ctx context.Context) (context.Context, error) {
	log.Info(ctx, "Called")

	dbTrx := r.DB.Begin()

	trxCtx := context.WithValue(ctx, ContextDBValue, dbTrx)

	return trxCtx, nil
}

func (r *GormTransactionImpl) CommitTransaction(ctx context.Context) error {
	log.Info(ctx, "Called")

	db, err := ExtractDB(ctx)
	if err != nil {
		return err
	}

	return db.Commit().Error
}

func (r *GormTransactionImpl) RollbackTransaction(ctx context.Context) error {
	log.Info(ctx, "Called")

	db, err := ExtractDB(ctx)
	if err != nil {
		return err
	}

	return db.Rollback().Error
}

type GormReadOnlyImpl struct {
	DB *gorm.DB
}

func (r *GormReadOnlyImpl) GetDatabase(ctx context.Context) (context.Context, error) {
	log.Info(ctx, "Called")

	trxCtx := context.WithValue(ctx, ContextDBValue, r.DB)

	return trxCtx, nil
}

// ExtractDB is used by other repo to extract the database from context
func ExtractDB(ctx context.Context) (*gorm.DB, error) {

	db, ok := ctx.Value(ContextDBValue).(*gorm.DB)
	if !ok {
		return nil, apperror.DatabaseNotFoundInContextError
	}

	return db, nil
}

type PaginateRequest struct {
	Page       int    `json:"page" form:"page" binding:"numeric"`           //
	PageSize   int    `json:"page_size" form:"page_size" binding:"numeric"` //
	Sort       string `json:"sort" form:"sort"`                             //
	TotalRows  int64  `json:"total_rows"`
	TotalPages int    `json:"total_pages"`
}

// Paginate is pagination offset technique
func Paginate(req *PaginateRequest, db *gorm.DB, table interface{}) func(*gorm.DB) *gorm.DB {
	if req.Page <= 0 && req.PageSize <= 0 {
		req.Page = 0
		req.PageSize = 0
	} else {

		if req.Page <= 0 {
			req.Page = 1
		}

		if req.PageSize <= 0 {
			req.PageSize = 10
		}
		var totalRows int64
		db.Model(table).Count(&totalRows)
		req.TotalRows = totalRows
		totalPages := int(math.Ceil(float64(totalRows) / float64(req.PageSize)))
		req.TotalPages = totalPages
	}
	return func(db *gorm.DB) *gorm.DB {
		if req.Page <= 0 && req.PageSize <= 0 {
			return db
		}

		offset := (req.Page - 1) * req.PageSize
		return db.Offset(offset).Limit(req.PageSize)
	}
}
