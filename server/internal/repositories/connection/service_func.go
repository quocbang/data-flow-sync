package connection

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/quocbang/data-flow-sync/server/internal/repositories"
	"github.com/quocbang/data-flow-sync/server/internal/repositories/errors"
	"github.com/quocbang/data-flow-sync/server/internal/repositories/services/account"
	"github.com/quocbang/data-flow-sync/server/internal/repositories/services/file"
	"github.com/quocbang/data-flow-sync/server/internal/repositories/services/mr"
	"github.com/quocbang/data-flow-sync/server/internal/repositories/services/station"
	"github.com/quocbang/data-flow-sync/server/internal/repositories/services/stationgroup"
)

func (s *DB) Station() repositories.StationServices {
	return station.NewService(s.Postgres)
}

func (s *DB) Close() error {
	// close postgresql
	db, err := s.Postgres.DB()
	if err != nil {
		return err
	}
	if err := db.Close(); err != nil {
		return err
	}

	return nil
}

func (s *DB) StationGroup() repositories.StationGroupServices {
	return stationgroup.NewService(s.Postgres)
}

func (s *DB) Account() repositories.AccountServices {
	return account.NewService(s.Postgres)
}

func (d *DB) Begin(ctx context.Context, opts ...*sql.TxOptions) (repositories.Repositories, error) {
	if d.TxFlag {
		return d, errors.Error{Code: errors.Code_ALREADY_IN_TRANSACTION}
	}

	newHandlerPtr := &DB{
		Postgres: d.Postgres.Begin(opts...),
		TxFlag:   true,
	}
	return newHandlerPtr, nil
}

func (d *DB) Commit() error {
	if !d.TxFlag {
		return fmt.Errorf("not in transaction")
	}
	return d.Postgres.Commit().Error
}

func (d *DB) RollBack() error {
	if !d.TxFlag {
		return fmt.Errorf("not in transaction")
	}
	return d.Postgres.Rollback().Error
}

func (s *DB) File() repositories.FileServices {
	return file.NewService(s.Postgres, s.Redis)
}

func (s *DB) MergeRequest() repositories.MergeRequestServices {
	return mr.NewService(s.Postgres, s.Redis)
}
