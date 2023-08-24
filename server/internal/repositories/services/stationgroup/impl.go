package stationgroup

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/quocbang/data-flow-sync/server/internal/repositories"
)

type service struct {
	pg    *gorm.DB
	redis *redis.Client
}

func NewService(pg *gorm.DB, rd *redis.Client) repositories.StationGroupServices {
	return service{
		pg:    pg,
		redis: rd,
	}
}
