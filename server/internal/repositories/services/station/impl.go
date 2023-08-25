package station

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis"
	"gorm.io/gorm"

	type_ "github.com/quocbang/data-flow-sync/server/assets/protobuf/types"
	"github.com/quocbang/data-flow-sync/server/internal/repositories"
	station "github.com/quocbang/data-flow-sync/server/internal/repositories/orm/models"
	"github.com/quocbang/data-flow-sync/server/swagger/models"
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

func (s service) UpsertStationDataStorage(ctx context.Context, req *models.Station) (repositories.CreateStationDataStorageReply, error) {
	jsonStationData, err := json.Marshal(req)
	if err != nil {
		return repositories.CreateStationDataStorageReply{}, err
	}

	reply := s.pg.Create(&station.DataStorage{
		ID:      req.ID,
		Type:    type_.Type_STATION,
		Content: jsonStationData,
	})
	return repositories.CreateStationDataStorageReply{RowsAffected: reply.RowsAffected}, reply.Error
}
