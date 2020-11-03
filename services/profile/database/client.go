package database

import (
	"time"

	pg "github.com/go-pg/pg/v10"

	"in-backend/services/profile"
	"in-backend/services/profile/configs"
)

// constants for db client config
const (
	ReadTimeout  = 30 * time.Second
	WriteTimeout = 30 * time.Second
	PoolSize     = 10
	MinIdleConns = 10
)

type postgresClient struct {
	DB *pg.DB
}

func (p postgresClient) GetConnection() *pg.DB {
	return p.DB
}

func (p postgresClient) Close() error {
	return p.DB.Close()
}

// NewClient returns a new PostgresClient
func NewClient(cfg configs.Config) profile.PostgresClient {
	opt := GetPgConnectionOptions(cfg)
	db := pg.Connect(opt)
	return postgresClient{
		DB: db,
	}
}

// GetPgConnectionOptions returns pg Options based on config
func GetPgConnectionOptions(cfg configs.Config) *pg.Options {
	return &pg.Options{
		Addr:            cfg.Database.Address,
		User:            cfg.Database.Username,
		Password:        cfg.Database.Password,
		Database:        cfg.Database.Database,
		ApplicationName: cfg.AppName,
		ReadTimeout:     ReadTimeout,
		WriteTimeout:    WriteTimeout,
		PoolSize:        PoolSize,
		MinIdleConns:    MinIdleConns,
	}
}
