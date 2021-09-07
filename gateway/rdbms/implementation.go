package rdbms

import (
	"github.com/raismaulana/blogP/infrastructure/database"
	"github.com/raismaulana/blogP/infrastructure/envconfig"
	"github.com/raismaulana/blogP/infrastructure/migration"
	"gorm.io/gorm"
)

type RDBMSGateway struct {
	database.GormReadOnlyImpl
	database.GormTransactionImpl
}

// NewRDBMSGateway ...
func NewRDBMSGateway(env *envconfig.EnvConfig, db *gorm.DB) (*RDBMSGateway, error) {
	err := migration.RDBMSMigration(db, env)
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
