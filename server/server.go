package server

import (
	"github.com/quocbang/data-flow-sync/server/config"
	"github.com/quocbang/data-flow-sync/server/internal/repositories"
	"github.com/quocbang/data-flow-sync/server/internal/repositories/connection"
)

func RegisterRepository(cfs config.DatabaseGroup) (repositories.Repositories, error) {
	opts := []connection.Options{
		connection.MaybeMigrate(),
	}

	return connection.New(connection.Database(cfs), opts...)
}
