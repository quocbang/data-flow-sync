package station

import (
	"context"
	"fmt"
	"time"

	"github.com/go-openapi/runtime/middleware"

	"github.com/quocbang/data-flow-sync/server/internal/repositories"
	"github.com/quocbang/data-flow-sync/server/internal/services"
	"github.com/quocbang/data-flow-sync/server/swagger/models"
	"github.com/quocbang/data-flow-sync/server/swagger/restapi/operations/station"
	"github.com/quocbang/data-flow-sync/server/utils"
	"github.com/quocbang/data-flow-sync/server/utils/function"
	"github.com/quocbang/data-flow-sync/server/utils/roles"
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
	ctx, cancel := context.WithTimeout(params.HTTPRequest.Context(), (time.Minute * 1))
	defer cancel()
	if params.Body.Files == nil {
		return station.NewCreateStationMergeRequestBadRequest().WithPayload(&models.ErrorResponse{
			Code:    0,
			Details: "missing station body",
		})
	}

	for idx, b := range params.Body.Files {
		fmt.Println(idx, b)
	}
	reply, err := s.Repository.Station().CreateMergeRequest(ctx)
	if err != nil {
		return utils.ParseError(ctx, station.NewGetStationMergeRequestDefault(0), err)
	}

	return station.NewCreateStationMergeRequestOK().WithPayload(&station.CreateStationMergeRequestOKBody{
		MergeRequestID: reply.MergeRequestID,
	})
}

func (s Station) GetStationMergeRequest(params station.GetStationMergeRequestParams, principal *models.Principal) middleware.Responder {
	return nil
}
