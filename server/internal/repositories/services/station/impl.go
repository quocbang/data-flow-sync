package station

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"
	"sigs.k8s.io/yaml"

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

func (s service) UpsertStationDataStorage(ctx context.Context, req *models.CreateStationDataStorage) (repositories.CreateStationDataStorageReply, error) {
	errorMessage := ""
	data, err := yaml.YAMLToJSON([]byte(req.Content))
	if err != nil {
		errorMessage += fmt.Sprintf("%s station is not a valid YAML: %v\n", req.ID, err)
	}

	if errorMessage != "" {
		return repositories.CreateStationDataStorageReply{}, fmt.Errorf(errorMessage)
	}

	reply := s.pg.Create(&station.DataStorage{
		ID:        req.ID,
		Type:      type_.Type_STATION,
		Content:   data,
		CreatedBy: "tester_AI",
		CreatedAt: 454545,
	})
	return repositories.CreateStationDataStorageReply{RowsAffected: reply.RowsAffected}, reply.Error
}
