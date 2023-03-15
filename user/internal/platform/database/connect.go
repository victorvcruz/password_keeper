package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"os"
)

type DatabaseClient interface {
	Connect() (*gorm.DB, error)
}

type database struct {
	localhost string
	user      string
	password  string
	port      string
	dbname    string
	sslmode   string
	timeZone  string
	schema    string
}

func NewDatabase() DatabaseClient {
	return &database{
		localhost: os.Getenv("POSTGRES_HOST"),
		user:      os.Getenv("POSTGRES_USER"),
		password:  os.Getenv("POSTGRES_PASSWORD"),
		port:      os.Getenv("POSTGRES_PORT"),
		dbname:    os.Getenv("POSTGRES_NAME"),
		sslmode:   os.Getenv("POSTGRES_SLLMODE"),
		timeZone:  os.Getenv("POSTGRES_TIMEZONE"),
		schema:    os.Getenv("POSTGRES_SCHEMA"),
	}
}

func (d *database) Connect() (*gorm.DB, error) {

	postgresDSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		d.localhost, d.user, d.password, d.dbname, d.port, d.sslmode, d.timeZone)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  postgresDSN,
		PreferSimpleProtocol: true,
	}), &gorm.Config{Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{TablePrefix: fmt.Sprintf("%s.", d.schema), SingularTable: false}})

	if err != nil {
		return nil, err
	}

	return db, err
}
