package redis

import (
	"github.com/go-redis/redis/v8"
	"github.com/raismaulana/blogP/infrastructure/database"
)

type RedisGateway struct {
	database.RedisClientImpl
}

func NewRedisGateway(rdb *redis.Client) *RedisGateway {
	return &RedisGateway{
		RedisClientImpl: database.RedisClientImpl{
			Rdb: rdb,
		},
	}
}
