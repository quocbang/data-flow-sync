package suite

import (
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/jessevdk/go-flags"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/quocbang/data-flow-sync/server/config"
	"github.com/quocbang/data-flow-sync/server/internal/repositories"
	"github.com/quocbang/data-flow-sync/server/internal/repositories/connection"
)

var postGresTest struct {
	DBSchema   string `long:"db_schema" description:"The test DB schema" env:"DB_SCHEMA_TEST"`
	DBAddress  string `long:"db_address" description:"The test DB address" env:"DB_ADDRESS_TEST"`
	DBPort     int    `long:"db_port" description:"The test DB port" env:"DB_PORT_TEST"`
	DBUsername string `long:"db_username" description:"The test DB username" env:"DB_USERNAME_TEST"`
	DBPassword string `long:"db_password" description:"The test DB password" env:"DB_PASSWORD_TEST"`
	DBDatabase string `long:"db_database" description:"The test DB database" env:"DB_DATABASE_TEST"`
}

var redisTest struct {
	RedisAddress  string `long:"redis_address" description:"The test Redis address" env:"REDIS_ADDRESS_TEST"`
	RedisPassword string `long:"redis_password" description:"The test Redis password " env:"REDIS_PASSWORD_TEST"`
	RedisDatabase int    `long:"redis_database" description:"The test Redis database" env:"REDIS_DATABASE_TEST"`
}

var smtpTest struct {
	SmtpServer  string `long:"smtp-server" description:"the smtp server" env:"SMTP_SERVER_TEST"`
	SmtpPort    int    `long:"smtp-port" description:"the smtp port" env:"SMTP_PORT_TEST"`
	SenderEmail string `long:"smtp-sender" description:"sender email" env:"SMTP_SENDER_TEST"`
	Password    string `long:"smtp-sender-password" description:"sender's password" env:"SMTP_PASSWORD_TEST"`
}

func InitializeDB() (repo repositories.Repositories, db *gorm.DB, rd *redis.Client, err error) {
	if err = parseFlags(); err != nil {
		return // return when parse failed.
	}

	repo, db, rd, err = newTestDataManager()
	return
}

func parseFlags() error {
	configurations := []swag.CommandLineOptionsGroup{
		{
			LongDescription:  "PostGres Configuration",
			ShortDescription: "PostGres Configuration",
			Options:          &postGresTest,
		},
		{
			LongDescription:  "Redis Configuration",
			ShortDescription: "Redis Configuration",
			Options:          &redisTest,
		},
		{
			LongDescription:  "SMTP Configuration",
			ShortDescription: "SMTP Configuration",
			Options:          &smtpTest,
		},
	}

	parse := flags.NewParser(nil, flags.IgnoreUnknown)
	for _, opt := range configurations {
		if _, err := parse.AddGroup(opt.LongDescription, opt.LongDescription, opt.Options); err != nil {
			return err
		}
	}

	if _, err := parse.Parse(); err != nil {
		return fmt.Errorf("failed to parse postgres flags")
	}

	return nil
}

func newTestDataManager() (dm repositories.Repositories, db *gorm.DB, rd *redis.Client, err error) {
	conn := connection.Database{
		Postgres: config.PostgresConfig{
			Address:  postGresTest.DBAddress,
			Port:     postGresTest.DBPort,
			UserName: postGresTest.DBUsername,
			Password: postGresTest.DBPassword,
			Name:     postGresTest.DBDatabase,
			Schema:   postGresTest.DBSchema,
		},
		Redis: config.RedisConfig{
			Address:  redisTest.RedisAddress,
			Password: redisTest.RedisPassword,
			Database: redisTest.RedisDatabase,
		},
		SMTP: config.SmtpConfig{
			SmtpServer:  smtpTest.SmtpServer,
			SmtpPort:    smtpTest.SmtpPort,
			SenderEmail: smtpTest.SenderEmail,
			Password:    smtpTest.Password,
		},
	}

	rd, err = connection.NewRedis(conn.Redis)
	if err != nil {
		return nil, nil, nil, err
	}

	dm, err = connection.New(conn, connection.MaybeMigrate())
	if err != nil {
		return nil, nil, nil, err
	}

	db, err = connection.NewPostgres(conn.Postgres)
	if err != nil {
		return nil, nil, nil, err
	}

	return
}
