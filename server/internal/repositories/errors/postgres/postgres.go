package postgres

import (
	"github.com/jackc/pgx/v5/pgconn"
)

const (
	NotNullViolation    string = "23502"
	ForeignKeyViolation string = "23503"
	UniqueViolation     string = "23505"
)

func ErrorIs(err error, code string) bool {
	if e, ok := err.(*pgconn.PgError); ok {
		return e.Code == code
	}
	return false
}
