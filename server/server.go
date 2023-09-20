package server

import (
	"github.com/quocbang/data-flow-sync/server/config"
	"github.com/quocbang/data-flow-sync/server/internal/mailserver"
	mail_connection "github.com/quocbang/data-flow-sync/server/internal/mailserver/connection"
	rd_connection "github.com/quocbang/data-flow-sync/server/internal/redis_conn"
	redisconn "github.com/quocbang/data-flow-sync/server/internal/redis_conn/connection"
	"github.com/quocbang/data-flow-sync/server/internal/repositories"
	"github.com/quocbang/data-flow-sync/server/internal/repositories/connection"
)

func RegisterRepository(cfs config.DatabaseGroup) (repositories.Repositories, error) {
	opts := []connection.Options{
		connection.MaybeMigrate(),
	}

	return connection.New(connection.Database(cfs), opts...)
}

func RegisterSmtp(cfs config.SmtpConfig) (mailserver.MailServer, error) {
	return mail_connection.NewSMTP(cfs)
}

func RegisterRedis(cfs config.RedisConfig) (rd_connection.RedisConn, error) {
	return redisconn.NewRDConn(cfs)
}
