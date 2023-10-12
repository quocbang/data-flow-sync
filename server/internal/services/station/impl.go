package station

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-openapi/runtime/middleware"

	"github.com/quocbang/data-flow-sync/server/internal/repositories"
	repoErr "github.com/quocbang/data-flow-sync/server/internal/repositories/errors"
	m "github.com/quocbang/data-flow-sync/server/internal/repositories/orm/models"
	"github.com/quocbang/data-flow-sync/server/internal/services"
	"github.com/quocbang/data-flow-sync/server/swagger/models"
	"github.com/quocbang/data-flow-sync/server/swagger/restapi/operations/station"
	"github.com/quocbang/data-flow-sync/server/utils"
	"github.com/quocbang/data-flow-sync/server/utils/diff"
	"github.com/quocbang/data-flow-sync/server/utils/function"
	"github.com/quocbang/data-flow-sync/server/utils/roles"
	"github.com/quocbang/data-flow-sync/server/utils/types"
)

type Station struct {
	Repository             repositories.Repositories
	MergeRequestExpiryTime int64
	HasPermission          func(function.FuncName, roles.Roles) bool
}

func NewStation(repo repositories.Repositories, mrExpiryTime int64,
	hasPermission func(function.FuncName, roles.Roles) bool) services.StationServices {
	return Station{
		Repository:             repo,
		MergeRequestExpiryTime: mrExpiryTime,
		HasPermission:          hasPermission,
	}
}

func (s Station) CreateStationMergeRequest(params station.CreateStationMergeRequestParams, principal *models.Principal) middleware.Responder {
	if !s.HasPermission(function.FuncName_CREATE_STATION_MERGE_REQUEST, roles.Roles(principal.Role)) {
		return station.NewCreateStationMergeRequestDefault(http.StatusForbidden).WithPayload(&models.ErrorResponse{
			Code:    int64(repoErr.Code_FORBIDDEN),
			Details: "permission denied",
		})
	}
	ctx := params.HTTPRequest.Context()

	if len(params.Body.Files) == 0 {
		return station.NewCreateStationMergeRequestBadRequest().WithPayload(&models.ErrorResponse{
			Code:    0,
			Details: "missing request body",
		})
	}

	createMergeRequest := repositories.CreateMRRequest{}
	year, month, day := time.Now().Local().Date()
	user := principal.ID
	createMergeRequest.MergeRequest = m.MergeRequest{
		Type: types.Types_STATION,
		File: make(m.Files, len(params.Body.Files)),
		Information: m.MergeRequestInfo{
			HistoryChanged: []string{fmt.Sprintf("Create a merge request at %d-%d-%d, by %s", year, month, day, user)},
			Status: m.MergeRequestStatus{
				IsApproved: false,
				IsMerged:   false,
				IsOpen:     true,
			},
		},
		CreatedBy: user,
	}

	for idx, f := range params.Body.Files {
		// unmarshal file to struct
		stationData := repositories.Station{}
		if err := json.Unmarshal([]byte(f), &stationData); err != nil {
			return station.NewCreateStationMergeRequestBadRequest().WithPayload(&models.ErrorResponse{
				Details: fmt.Sprintf("failed to unmarshal file [%d], error: %v", idx, err),
			})
		}

		// check whether the file was requested or not, if yes return the conflict error.
		_, err := s.Repository.MergeRequest().GetMergeRequestOpeningByFileID(ctx, stationData.ID)
		if err == nil { // err == nil => record is found
			return station.NewCreateStationMergeRequestBadRequest().WithPayload(&models.ErrorResponse{
				Code:    int64(repoErr.Code_FILE_CONFLICTED),
				Details: fmt.Sprintf("file [%s] was conflicted", stationData.ID),
			})
		}

		// get old file data
		oldFile, err := s.Repository.File().GetFile(ctx, repositories.GetFileRequest{
			ID:   stationData.ID,
			Type: types.Types_STATION,
		})
		if err != nil {
			if errors.Is(err, repoErr.ErrDataNotFound) {
				oldFile.File.FileContent = []byte(`{}`)
			} else {
				return utils.ParseError(ctx, station.NewCreateStationMergeRequestDefault(0), err)
			}
		}

		// compare file
		added, deleted, err := diff.FindDiff[repositories.Station](oldFile.File.FileContent, []byte(f))
		if err != nil {
			return station.NewCreateStationMergeRequestBadRequest().WithPayload(&models.ErrorResponse{
				Details: err.Error(),
			})
		}

		createMergeRequest.MergeRequest.File[idx] = m.FileMergeRequest{
			FileID: stationData.ID,
			FileCompare: m.FileCompares{
				FileBeforeMerge: oldFile.File.FileContent,
				FileAfterMerge:  []byte(f),
			},
			FileChanged: m.Changed{
				Added:   added,
				Deleted: deleted,
			},
		}
	}

	reply, err := s.Repository.MergeRequest().CreateMergeRequest(ctx, createMergeRequest)
	if err != nil {
		return utils.ParseError(ctx, station.NewGetStationMergeRequestDefault(0), err)
	}

	return station.NewCreateStationMergeRequestOK().WithPayload(&station.CreateStationMergeRequestOKBody{
		MergeRequestID: reply.MergeRequestID,
	})
}

func (s Station) GetStationMergeRequest(params station.GetStationMergeRequestParams, principal *models.Principal) middleware.Responder {
	if !s.HasPermission(function.FuncName_GET_STATION_MERGE_REQUEST, roles.Roles(principal.Role)) {
		return station.NewCreateStationMergeRequestDefault(http.StatusForbidden).WithPayload(&models.ErrorResponse{
			Code:    int64(repoErr.Code_FORBIDDEN),
			Details: "permission denied",
		})
	}
	if params.ID == 0 {
		return station.NewGetStationMergeRequestBadRequest().WithPayload(&models.ErrorResponse{
			Details: "missing merge request id",
		})
	}

	ctx := params.HTTPRequest.Context()
	reply, err := s.Repository.MergeRequest().GetMergeRequest(ctx, repositories.GetMRRequest{
		MergeRequestID: params.ID,
	})
	if err != nil {
		return utils.ParseError(ctx, station.NewGetStationMergeRequestDefault(0), err)
	}

	stationMergeRequest := make([]*models.StationMergeRequest, len(reply.MergeRequest.File))
	for i, f := range reply.MergeRequest.File {
		added, deleted, file := &models.Station{}, &models.Station{}, &models.Station{}
		if err := json.Unmarshal(f.FileChanged.Added, added); err != nil {
			return station.NewGetStationMergeRequestInternalServerError().WithPayload(&models.ErrorResponse{
				Details: fmt.Sprintf("failed to unmarshal added, error: %v", err),
			})
		}
		if err := json.Unmarshal(f.FileChanged.Deleted, deleted); err != nil {
			return station.NewGetStationMergeRequestInternalServerError().WithPayload(&models.ErrorResponse{
				Details: fmt.Sprintf("failed to unmarshal deleted, error: %v", err),
			})
		}
		if err := json.Unmarshal(f.FileCompare.FileAfterMerge, file); err != nil {
			return station.NewGetStationMergeRequestInternalServerError().WithPayload(&models.ErrorResponse{
				Details: fmt.Sprintf("failed to unmarshal file after merge, error: %v", err),
			})
		}

		stationMergeRequest[i] = &models.StationMergeRequest{
			Added:   added,
			Deleted: deleted,
			File:    file,
		}
	}

	return station.NewGetStationMergeRequestOK().WithPayload(&station.GetStationMergeRequestOKBody{
		Data: stationMergeRequest,
		MergeRequestInfo: &station.GetStationMergeRequestOKBodyMergeRequestInfo{
			HistoryChanged: reply.MergeRequest.Information.HistoryChanged,
		},
		MergeRequestStatus: &station.GetStationMergeRequestOKBodyMergeRequestStatus{
			IsApproved: &reply.MergeRequest.Information.Status.IsApproved,
			IsMerged:   &reply.MergeRequest.Information.Status.IsMerged,
			IsOpening:  &reply.MergeRequest.Information.Status.IsOpen,
		},
	})
}
