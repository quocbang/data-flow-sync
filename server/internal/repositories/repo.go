package repositories

import (
	"context"
	"database/sql"
)

type Repositories interface {
	Close() error
	Services
	Begin(context.Context, ...*sql.TxOptions) (Repositories, error)
	Commit() error
	RollBack() error
}
