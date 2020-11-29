package database

import (
	"fmt"
	"in-backend/services/joblisting/configs"
	"time"

	pg "github.com/go-pg/pg/v10"
)

// constants for db client config
const (
	ReadTimeout  = 30 * time.Second
	WriteTimeout = 30 * time.Second
	PoolSize     = 10
	MinIdleConns = 10
)

// NewDatabase returns a new PostgresDB
func NewDatabase(opt *pg.Options) *pg.DB {
	return pg.Connect(opt)
}

// GetPgConnectionOptions returns pg Options based on config
func GetPgConnectionOptions(cfg configs.Config) *pg.Options {
	return &pg.Options{
		Addr:            fmt.Sprintf("%s:%s", cfg.Database.Address, cfg.Database.Port),
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
