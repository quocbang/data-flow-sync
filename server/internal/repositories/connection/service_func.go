package connection

import (
	"github.com/quocbang/data-flow-sync/server/internal/repositories"
	"github.com/quocbang/data-flow-sync/server/internal/repositories/services/account"
	"github.com/quocbang/data-flow-sync/server/internal/repositories/services/station"
	"github.com/quocbang/data-flow-sync/server/internal/repositories/services/stationgroup"
)

func (s *DB) Station() repositories.StationServices {
	return station.NewService(s.Postgres, s.Redis)
}

func (s *DB) Close() error {
	return nil
}

func (s *DB) StationGroup() repositories.StationGroupServices {
	return stationgroup.NewService(s.Postgres, s.Redis)
}

func (s *DB) Account() repositories.AccountServices {
	return account.NewService(s.Postgres, s.Redis)
}
