package repositorieserror

import (
	"errors"

	"gorm.io/gorm"

	repoErr "github.com/quocbang/data-flow-sync/server/internal/repositories/errors"
	"github.com/quocbang/data-flow-sync/server/internal/repositories/errors/postgres"
)

// transformErr return transformed error if satisfy the condition, if not return origin error.
func transformErr(condition func(error) bool, transformedErr error, originErr error) error {
	if condition(originErr) {
		return transformedErr
	}
	return originErr
}

func transformPostgresExistedErr(e error) error {
	return transformErr(isPostgresExistedErr, repoErr.ErrDataExisted, e)
}

// This function checks id given error is Postgres data existed error.
// return true if is Postgres data existed error, if not return false.
func isPostgresExistedErr(e error) bool {
	return postgres.ErrorIs(e, postgres.UniqueViolation)
}

func transformGormRecordNotFoundErr(e error) error {
	return transformErr(isGormRecordNotFoundErr, repoErr.ErrDataNotFound, e)
}

// This function checks if given error is GORM record not found error.
// return true if is GORM record not found error else case is false.
func isGormRecordNotFoundErr(e error) bool {
	return errors.Is(e, gorm.ErrRecordNotFound)
}

func MapError(err error) error {
	return transformPostgresExistedErr(transformGormRecordNotFoundErr(err))
}
