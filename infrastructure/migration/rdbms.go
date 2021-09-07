package migration

import (
	"context"
	"time"

	"github.com/raismaulana/blogP/domain/entity"
	"github.com/raismaulana/blogP/infrastructure/envconfig"
	"github.com/raismaulana/blogP/infrastructure/log"
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
	db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&entity.User{
			ID:          1,
			Username:    env.SuperUsername,
			Name:        "King",
			Email:       "king@mize.com",
			Password:    env.SuperUsername,
			City:        "Jakarta",
			Country:     "Indonesia",
			Birthday:    time.Date(2020, 8, 2, 0, 0, 0, 0, time.UTC),
			WebProfile:  null.String{},
			Role:        "king",
			ActivatedAt: null.NewTime(time.Now(), true),
		}).Error
		if err != nil {
			return err
		}
		return nil
	})
	return nil
}
