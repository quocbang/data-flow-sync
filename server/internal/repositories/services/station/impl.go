package station

import (
	"context"

	"gorm.io/gorm"

	"github.com/quocbang/data-flow-sync/server/internal/repositories"
)

type service struct {
	pg *gorm.DB
}

func NewService(pg *gorm.DB) repositories.StationServices {
	return service{
		pg:    pg,
		redis: redis,
	}
}

// CreateMergeRequest is create merge request and then reply merge request id
func (s service) CreateMergeRequest(ctx context.Context) (repositories.CreateStationMRReply, error) {

	return repositories.CreateStationMRReply{}, nil
}

// GetMergeRequest is get merge request information.
func (s service) GetMergeRequest(ctx context.Context, req repositories.GetMRRequest) (repositories.GetMRReply, error) {
	return repositories.GetMRReply{}, nil
}
