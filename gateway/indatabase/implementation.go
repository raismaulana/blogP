package indatabase

import (
	"context"

	"github.com/raismaulana/blogP/domain/entity"
	"github.com/raismaulana/blogP/gateway/shared"
	"github.com/raismaulana/blogP/infrastructure/database"
	"github.com/raismaulana/blogP/infrastructure/envconfig"
	"github.com/raismaulana/blogP/infrastructure/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type inDatabaseGateway struct {
	database.GormReadOnlyImpl
	database.GormTransactionImpl
	shared.SharedGateway
}

// NewInDatabaseGateway ...
func NewInDatabaseGateway(env *envconfig.EnvConfig, db *gorm.DB) (*inDatabaseGateway, error) {
	err := db.AutoMigrate(&entity.User{})
	if err != nil {
		return nil, err
	}

	return &inDatabaseGateway{
		GormReadOnlyImpl: database.GormReadOnlyImpl{
			DB: db,
		},
		GormTransactionImpl: database.GormTransactionImpl{
			DB: db,
		},
		SharedGateway: shared.SharedGateway{
			Env: env,
		},
	}, nil
}

func (r *inDatabaseGateway) SaveUser(ctx context.Context, obj *entity.User) error {
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

func (r *inDatabaseGateway) HashPassword(ctx context.Context, plainPassword string) (string, error) {
	log.Info(ctx, "called")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), 10)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (r *inDatabaseGateway) FindUserByUsername(ctx context.Context, username string, scope bool) (*entity.User, error) {
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

func (r *inDatabaseGateway) FindUserByEmail(ctx context.Context, email string, scope bool) (*entity.User, error) {
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

func (r *inDatabaseGateway) FindUserByID(ctx context.Context, ID int64, scope bool) (*entity.User, error) {
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

func (r *inDatabaseGateway) FetchUsers(ctx context.Context, scope bool) ([]*entity.User, error) {
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

func (r *inDatabaseGateway) DeleteUser(ctx context.Context, obj *entity.User) error {
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
