package redisconn

import (
	"context"

	"github.com/quocbang/data-flow-sync/server/config"
	"github.com/quocbang/data-flow-sync/server/internal/repositories/connection/logging"
	"github.com/redis/go-redis/v9"
)

// NewRedis is connect to redis database.
func NewRedis(rdCf config.RedisConfig) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     rdCf.Address,
		Password: rdCf.Password,
	})
	redis.SetLogger(logging.NewRedisLogger())

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return rdb, nil
}

func NewRDConn(rdCf config.RedisConfig) (*RDConn, error) {
	rd, err := NewRedis(rdCf)
	return &RDConn{
		rd: rd,
	}, err
}
