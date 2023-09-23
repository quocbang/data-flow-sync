package station

import (
	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"

	"github.com/quocbang/data-flow-sync/server/internal/repositories"
)

type service struct {
	pg    *gorm.DB
	redis *redis.Client
}

func NewService(pg *gorm.DB, redis *redis.Client) repositories.StationServices {
	return service{
		pg:    pg,
		redis: redis,
	}
}
