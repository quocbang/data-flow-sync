package server

import (
	"gitlab.com/quocbang/data-flow-sync/server/config"
	"gitlab.com/quocbang/data-flow-sync/server/internal/repositories"
	"gitlab.com/quocbang/data-flow-sync/server/internal/repositories/connection"
)

func RegisterRepository(cfs config.DatabaseGroup) (repositories.Repositories, error) {
	opts := []connection.Options{
		connection.MaybeMigrate(),
	}

	return connection.New(connection.Database(cfs), opts...)
}
