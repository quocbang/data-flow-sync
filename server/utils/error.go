package utils

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	repoErrors "github.com/quocbang/data-flow-sync/server/internal/repositories/errors"
	"github.com/quocbang/data-flow-sync/server/swagger/models"
)

type defaultError interface {
	middleware.Responder

	SetStatusCode(int)
	SetPayload(*models.ErrorResponse)
}

func ParseError(ctx context.Context, d defaultError, err error) middleware.Responder {
	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		d.SetStatusCode(http.StatusRequestTimeout)
		d.SetPayload(&models.ErrorResponse{
			Details: err.Error(),
		})
		return d
	}
	if e, ok := repoErrors.As(err); ok {
		d.SetStatusCode(http.StatusBadRequest)
		d.SetPayload(&models.ErrorResponse{
			Code:    int64(e.Code),
			Details: e.Details,
		})
	} else {
		d.SetStatusCode(http.StatusInternalServerError)
		d.SetPayload(&models.ErrorResponse{
			Details: err.Error(),
		})
	}

	return d
}
