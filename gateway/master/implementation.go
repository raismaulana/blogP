package master

import (
	"github.com/go-redis/redis/v8"
	"github.com/raismaulana/blogP/gateway/rdbms"
	redisGateway "github.com/raismaulana/blogP/gateway/redis"
	"github.com/raismaulana/blogP/gateway/shared"
	"github.com/raismaulana/blogP/infrastructure/auth"
	"github.com/raismaulana/blogP/infrastructure/envconfig"
	"gorm.io/gorm"
)

type masterGateway struct {
	rdbms.RDBMSGateway
	redisGateway.RedisGateway
	shared.SharedGateway
}

func NewMasterGateway(env *envconfig.EnvConfig, db *gorm.DB, rdb *redis.Client, jwtToken *auth.JWTToken) (*masterGateway, error) {
	rdbmsG, err := rdbms.NewRDBMSGateway(env, db)
	if err != nil {
		return nil, err
	}

	redisG := redisGateway.NewRedisGateway(rdb)
	sharedG := shared.NewSharedGateway(env, jwtToken)
	return &masterGateway{
		RDBMSGateway:  *rdbmsG,
		RedisGateway:  *redisG,
		SharedGateway: *sharedG,
	}, nil
}
