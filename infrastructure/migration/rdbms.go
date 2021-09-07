package migration

import (
	"context"
	"errors"
	"time"

	"github.com/raismaulana/blogP/domain/entity"
	"github.com/raismaulana/blogP/infrastructure/envconfig"
	"github.com/raismaulana/blogP/infrastructure/log"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

func RDBMSMigration(db *gorm.DB, env *envconfig.EnvConfig) error {
	log.Info(context.Background(), "Migrate RDBMS")
	if err := db.SetupJoinTable(&entity.Post{}, "Categories", &entity.PostCategory{}); err != nil {
		return err
	}
	if err := db.SetupJoinTable(&entity.Post{}, "Tags", &entity.PostTag{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(
		&entity.User{},
		&entity.Post{},
		&entity.Category{},
		&entity.Tag{}); err != nil {
		return err
	}
	// this transaction will always make user with id 1 become a king
	if err := db.Transaction(func(tx *gorm.DB) error {
		var user entity.User

		err := tx.Where("id_user = ? AND role = ?", 1, "king").First(&user).Error
		if err == nil {
			return errors.New("super user is already exsist")
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(env.SuperPassword), 10)
		if err != nil {
			return err
		}

		user = entity.User{
			ID:          1,
			Username:    env.SuperUsername,
			Name:        "King",
			Email:       "king@mize.com",
			Password:    string(hashedPassword),
			City:        "Jakarta",
			Country:     "Indonesia",
			Birthday:    time.Date(2020, 8, 2, 0, 0, 0, 0, time.UTC),
			WebProfile:  null.String{},
			Role:        "king",
			ActivatedAt: null.NewTime(time.Now(), true),
		}
		err = tx.Save(&user).Error
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		log.Error(context.Background(), err.Error())
	}

	return nil
}
