package indatabase

import (
	"context"

	"github.com/raismaulana/blogP/domain/entity"
	"github.com/raismaulana/blogP/infrastructure/database"
	"github.com/raismaulana/blogP/infrastructure/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type inDatabaseGateway struct {
	database.GormReadOnlyImpl
	database.GormTransactionImpl
}

// NewInDatabaseGateway ...
func NewInDatabaseGateway(db *gorm.DB) (*inDatabaseGateway, error) {
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

func (r *inDatabaseGateway) FindUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	log.Info(ctx, "called")

	db, err := database.ExtractDB(ctx)
	if err != nil {
		return nil, err
	}

	var user entity.User
	err = db.Where("username = ?", username).First(&user).Error
	if err != nil {
		log.Error(ctx, err.Error())
		return nil, err
	}

	return &user, nil
}

func (r *inDatabaseGateway) FindUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	log.Info(ctx, "called")

	db, err := database.ExtractDB(ctx)
	if err != nil {
		return nil, err
	}

	var user entity.User
	err = db.Where("email = ?", email).First(&user).Error
	if err != nil {
		log.Error(ctx, err.Error())
		return nil, err
	}

	return &user, nil
}

func (r *inDatabaseGateway) FindUserByID(ctx context.Context, ID int64) (*entity.User, error) {
	log.Info(ctx, "called")

	db, err := database.ExtractDB(ctx)
	if err != nil {
		return nil, err
	}

	var user entity.User
	err = db.Where("id_user = ?", ID).First(&user).Error
	if err != nil {
		log.Error(ctx, err.Error())
		return nil, err
	}

	return &user, nil
}

func (r *inDatabaseGateway) FetchUsers(ctx context.Context) ([]*entity.User, error) {
	log.Info(ctx, "called")

	db, err := database.ExtractDB(ctx)
	if err != nil {
		return nil, err
	}

	var objs []*entity.User
	err = db.Find(&objs).Error
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

	err = db.Unscoped().Delete(&obj).Error

	if err != nil {
		log.Error(ctx, err.Error())
		return err
	}
	return nil
}
