package connection

import (
	"gitlab.com/quocbang/data-flow-sync/server/internal/repositories"
	"gitlab.com/quocbang/data-flow-sync/server/internal/repositories/services/account"
	"gitlab.com/quocbang/data-flow-sync/server/internal/repositories/services/station"
	"gitlab.com/quocbang/data-flow-sync/server/internal/repositories/services/stationgroup"
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
