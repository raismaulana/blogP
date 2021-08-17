package rdbms

import (
	"context"

	"github.com/raismaulana/blogP/domain/entity"
	"github.com/raismaulana/blogP/infrastructure/database"
	"github.com/raismaulana/blogP/infrastructure/envconfig"
	"github.com/raismaulana/blogP/infrastructure/log"
	"gorm.io/gorm"
)

type RDBMSGateway struct {
	database.GormReadOnlyImpl
	database.GormTransactionImpl
}

// NewRDBMSGateway ...
func NewRDBMSGateway(env *envconfig.EnvConfig, db *gorm.DB) (*RDBMSGateway, error) {
	err := db.AutoMigrate(&entity.User{})
	if err != nil {
		return nil, err
	}

	return &RDBMSGateway{
		GormReadOnlyImpl: database.GormReadOnlyImpl{
			DB: db,
		},
		GormTransactionImpl: database.GormTransactionImpl{
			DB: db,
		},
	}, nil
}

func (r *RDBMSGateway) SaveUser(ctx context.Context, obj *entity.User) error {
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

func (r *RDBMSGateway) FindUserByUsername(ctx context.Context, username string, scope bool) (*entity.User, error) {
	log.Info(ctx, "called")

	db, err := database.ExtractDB(ctx)
	if err != nil {
		return nil, err
	}

	var user entity.User
	if scope {
		err = db.Where("username = ?", username).First(&user).Error
	} else {
		err = db.Unscoped().Where("username = ?", username).First(&user).Error
	}
	if err != nil {
		log.Error(ctx, err.Error())
		return nil, err
	}

	return &user, nil
}

func (r *RDBMSGateway) FindUserByEmail(ctx context.Context, email string, scope bool) (*entity.User, error) {
	log.Info(ctx, "called")

	db, err := database.ExtractDB(ctx)
	if err != nil {
		return nil, err
	}

	var user entity.User
	if scope {
		err = db.Where("email = ?", email).First(&user).Error
	} else {
		err = db.Unscoped().Where("email = ?", email).First(&user).Error

	}
	if err != nil {
		log.Error(ctx, err.Error())
		return nil, err
	}

	return &user, nil
}

func (r *RDBMSGateway) FindUserByID(ctx context.Context, ID int64, scope bool) (*entity.User, error) {
	log.Info(ctx, "called")

	db, err := database.ExtractDB(ctx)
	if err != nil {
		return nil, err
	}

	var user entity.User
	if scope {
		err = db.Where("id_user = ?", ID).First(&user).Error
	} else {
		err = db.Unscoped().Where("id_user = ?", ID).First(&user).Error
	}
	if err != nil {
		log.Error(ctx, err.Error())
		return nil, err
	}

	return &user, nil
}

func (r *RDBMSGateway) FetchUsers(ctx context.Context, scope bool) ([]*entity.User, error) {
	log.Info(ctx, "called")

	db, err := database.ExtractDB(ctx)
	if err != nil {
		return nil, err
	}

	var objs []*entity.User
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

func (r *RDBMSGateway) DeleteUser(ctx context.Context, obj *entity.User) error {
	log.Info(ctx, "called")

	db, err := database.ExtractDB(ctx)
	if err != nil {
		return err
	}

	err = db.Unscoped().Delete(&obj).Error // delete
	if err != nil {
		log.Error(ctx, err.Error())

		err = db.Delete(&obj).Error // soft delete
		if err != nil {
			log.Error(ctx, err.Error())
			return err
		}
	}
	return nil
}
