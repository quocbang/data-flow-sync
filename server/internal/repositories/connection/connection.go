package connection

import (
	"crypto/tls"
	"fmt"
	"net/smtp"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/quocbang/data-flow-sync/server/config"
	"github.com/quocbang/data-flow-sync/server/internal/repositories"
	"github.com/quocbang/data-flow-sync/server/internal/repositories/connection/logging"
	"github.com/quocbang/data-flow-sync/server/internal/repositories/orm/models"
)

type DB struct {
	Postgres *gorm.DB
	TxFlag   bool
}

type options struct {
	IsMigrate bool
}

type Options func(*options)

func MaybeMigrate() Options {
	return func(o *options) {
		o.IsMigrate = true
	}
}

type Database config.DatabaseGroup

func parseOption(opts ...Options) *options {
	options := &options{}
	for _, o := range opts {
		o(options)
	}

	return options
}

func NewPostgres(pgCf config.PostgresConfig) (*gorm.DB, error) {
	connectString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d search_path=%s",
		pgCf.Address,
		pgCf.UserName,
		pgCf.Password,
		pgCf.Name,
		pgCf.Port,
		pgCf.Schema,
	)
	db, err := gorm.Open(postgres.Open(connectString), &gorm.Config{
		Logger: logging.NewGormLogger(),
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewRepository(db Database) (*DB, error) {
	pg, err := NewPostgres(db.Postgres)
	if err != nil {
		return nil, err
	}

	return &DB{
		Postgres: pg,
	}, nil
}

func New(db Database, opts ...Options) (repositories.Repositories, error) {
	o := parseOption(opts...)

	database, err := NewRepository(db)
	if err != nil {
		return nil, err
	}
	if o.IsMigrate {
		if err := database.maybeMigrate(); err != nil {
			return nil, err
		}
	}

	return database, nil
}

// maybeMigrate is migrate table, trigger,...
func (d *DB) maybeMigrate() error {
	tableList := models.GetModelList()
	if err := migrateTable(d.Postgres, tableList...); err != nil {
		return err
	}
	return nil
}

// migrateTable is automatically create table if implement
// model.Models interface
func migrateTable(pg *gorm.DB, ms ...models.Models) error {
	dst := []any{}
	for _, m := range ms {
		dst = append(dst, m)
	}
	return pg.AutoMigrate(dst...)
}

func NewSMTPConnection(config config.SmtpConfig) (*smtp.Client, error) {
	// Create an authentication mechanism
	auth := smtp.PlainAuth("", config.SenderEmail, config.Password, config.SmtpServer)

	// Create a TLS configuration
	tlsConfig := &tls.Config{
		InsecureSkipVerify: false, // You might want to set this to false in production
		ServerName:         config.SmtpServer,
	}

	// Connect to the SMTP server with a TLS connection
	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", config.SmtpServer, config.SmtpPort), tlsConfig)
	if err != nil {
		return &smtp.Client{}, err
	}

	// Connect to the SMTP server
	// Establish the SMTP client
	client, err := smtp.NewClient(conn, config.SmtpServer)
	if err != nil {
		return &smtp.Client{}, err
	}

	// Authenticate
	if err := client.Auth(auth); err != nil {
		return &smtp.Client{}, err
	}

	return client, nil
}
