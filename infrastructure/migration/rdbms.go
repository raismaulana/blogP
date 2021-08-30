package migration

import (
	"context"

	"github.com/raismaulana/blogP/domain/entity"
	"github.com/raismaulana/blogP/infrastructure/log"
	"gorm.io/gorm"
)

func RDBMSMigration(db *gorm.DB) error {
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
	return nil
}
