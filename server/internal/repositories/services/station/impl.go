package station

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"
	"sigs.k8s.io/yaml"

	type_ "github.com/quocbang/data-flow-sync/server/assets/protobuf/types"
	"github.com/quocbang/data-flow-sync/server/internal/repositories"
	repo "github.com/quocbang/data-flow-sync/server/internal/repositories/orm/models"
	"github.com/quocbang/data-flow-sync/server/swagger/models"
	"github.com/quocbang/data-flow-sync/server/swagger/restapi/operations/station"
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

func (s service) UpsertStationDataStorage(params station.CreateStationDataStorageParams, principal *models.Principal) middleware.Responder {
	errorMessage := ""
	data, err := yaml.YAMLToJSON([]byte(params.Body.Content))
	if err != nil {
		errorMessage += fmt.Sprintf("%s station is not a valid YAML: %v\n", params.Body.ID, err)
	}

	if errorMessage != "" {
		// return repositories.CreateStationDataStorageReply{}, fmt.Errorf(errorMessage)
		return station.NewCreateStationDataStorageDefault(http.StatusBadRequest).WithPayload(&station.CreateStationDataStorageDefaultBody{
			Details: errorMessage,
		})
	}

	reply := s.pg.Create(&repo.DataStorage{
		ID:        params.Body.ID,
		Type:      type_.Type_STATION,
		Content:   data,
		CreatedBy: "tester_AI",
		CreatedAt: 454545,
	})

	if reply.Error != nil {
		return station.NewCreateStationDataStorageDefault(http.StatusBadRequest).WithPayload(&station.CreateStationDataStorageDefaultBody{
			Details: reply.Error.Error(),
		})
	}

	return station.NewCreateStationDataStorageOK()
}
