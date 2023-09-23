package file

import (
	"context"

	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"

	"github.com/quocbang/data-flow-sync/server/internal/repositories"
	"github.com/quocbang/data-flow-sync/server/internal/repositories/errors/repositorieserror"
	"github.com/quocbang/data-flow-sync/server/internal/repositories/orm/models"
)

type service struct {
	pg    *gorm.DB
	redis *redis.Client
}

func NewService(pg *gorm.DB, redis *redis.Client) repositories.FileServices {
	return service{
		pg:    pg,
		redis: redis,
	}
}

func (s service) GetFile(ctx context.Context, req repositories.GetFileRequest) (repositories.GetFileReply, error) {
	var file models.File
	err := s.pg.Where("id=? and type=?", req.ID, req.Type).Take(&file).Error
	return repositories.GetFileReply{
		File: file,
	}, repositorieserror.MapError(err)
}
