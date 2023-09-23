package mr

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"github.com/quocbang/data-flow-sync/server/internal/repositories"
	"github.com/quocbang/data-flow-sync/server/internal/repositories/errors/repositorieserror"
	"github.com/quocbang/data-flow-sync/server/internal/repositories/orm/models"
)

type service struct {
	pg *gorm.DB
}

func NewService(pg *gorm.DB) repositories.MergeRequestServices {
	return service{
		pg: pg,
	}
}

// CreateMergeRequest is create merge request and then reply merge request id
func (s service) CreateMergeRequest(ctx context.Context, req repositories.CreateMRRequest) (repositories.CreateMRReply, error) {
	mergeRequest := req.MergeRequest
	err := s.pg.Create(&mergeRequest).Error
	return repositories.CreateMRReply{
		MergeRequestID: int64(mergeRequest.ID),
	}, repositorieserror.MapError(err)
}

// GetMergeRequest is get merge request information.
func (s service) GetMergeRequest(ctx context.Context, req repositories.GetMRRequest) (repositories.GetMRReply, error) {
	var getMergeRequest models.MergeRequest
	err := s.pg.Where("id=?", req.MergeRequestID).Take(&getMergeRequest).Error
	return repositories.GetMRReply{
		MergeRequest: getMergeRequest,
	}, repositorieserror.MapError(err)
}

func (s service) GetMergeRequestOpeningByFileID(ctx context.Context, fileID string) (repositories.GetMergeRequestOpeningByFileIDReply, error) {
	var mergeRequest models.MergeRequest
	err := s.pg.Where(fmt.Sprintf(`file @> '[{"file_id":"%s"}]' and information @> '{"status": {"is_open": true}}'`, fileID)).Take(&mergeRequest).Error
	return repositories.GetMergeRequestOpeningByFileIDReply{
		MergeRequest: mergeRequest,
	}, repositorieserror.MapError(err)
}
