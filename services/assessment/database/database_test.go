package database

import (
	"fmt"
	"strings"
	"testing"

	"github.com/go-pg/migrations/v8"
	pg "github.com/go-pg/pg/v10"
	"github.com/ory/dockertest/v3"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"in-backend/services/assessment/configs"
)

// container declares the model for the docker container set up
type container struct {
	Pool     *dockertest.Pool
	Resource *dockertest.Resource
}

// setupPGContainer sets up a docker container from the postgres image
func setupPGContainer(opt *pg.Options) (*container, error) {
	var err error

	pool, err := dockertest.NewPool("")
	if err != nil {
		return nil, errors.Wrap(err, "Could not connect to docker")
	}

	resource, err := pool.Run(
		"postgres", "13-alpine",
		[]string{
			"POSTGRES_USER=" + opt.User,
			"POSTGRES_DB=" + opt.Database,
			"POSTGRES_PASSWORD=" + opt.Password,
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "Could not start resource")
	}

	return &container{
		Pool:     pool,
		Resource: resource,
	}, nil
}

func setupDB(c *container, opt *pg.Options, migrationsDir string) (*pg.DB, error) {
	var db *pg.DB
	err := c.Pool.Retry(func() error {
		options := *opt
		options.Addr = fmt.Sprintf("%s:%s", strings.Split(options.Addr, ":")[0], c.Resource.GetPort("5432/tcp"))
		db = NewDatabase(&options)
		_, err := db.Exec("select 1")
		return errors.Wrap(err, "Test query failed")
	})
	if err != nil {
		cleanContainer(c)
		return nil, errors.Wrap(err, "Could not start postgres")
	}

	mc := migrations.NewCollection()

	_, _, err = mc.Run(db, "init")
	if err != nil {
		cleanDb(db)
		cleanContainer(c)
		return nil, errors.Wrap(err, "Could not init migrations")
	}

	err = mc.DiscoverSQLMigrations(migrationsDir)
	if err != nil {
		cleanDb(db)
		cleanContainer(c)
		return nil, errors.Wrap(err, "Failed to read migrations")
	}

	_, _, err = mc.Run(db, "up")
	if err != nil {
		cleanDb(db)
		cleanContainer(c)
		return nil, errors.Wrap(err, "Could not migrate")
	}

	return db, nil
}

func cleanContainer(c *container) error {
	// kill and remove the container
	err := c.Pool.Purge(c.Resource)
	if err != nil {
		return errors.Wrap(err, "Failed to purge docker pool")
	}
	return nil
}

func cleanDb(db *pg.DB) error {
	err := db.Close()
	if err != nil {
		return errors.Wrap(err, "Failed to clean DB")
	}
	return nil
}

func TestGetPgConnectionOptions(t *testing.T) {
	cfg := configs.Config{
		AppName: "app",
		Database: configs.DbConfig{
			Address:  "address",
			Port:     "5432",
			Username: "user",
			Password: "password",
			Database: "database",
		},
	}

	want := &pg.Options{
		Addr:            "address:5432",
		User:            "user",
		Password:        "password",
		Database:        "database",
		ApplicationName: "app",
		ReadTimeout:     ReadTimeout,
		WriteTimeout:    WriteTimeout,
		PoolSize:        PoolSize,
		MinIdleConns:    MinIdleConns,
	}

	got := GetPgConnectionOptions(cfg)

	require.EqualValues(t, want, got)
}

func TestNewDatabase(t *testing.T) {
	testConfig, err := configs.LoadConfig(configs.TestFileName)
	require.NoError(t, err)

	opt := GetPgConnectionOptions(testConfig)

	c, err := setupPGContainer(opt)
	require.NoError(t, err)

	db, err := setupDB(c, opt, "../scripts/migrations/")
	require.NoError(t, err)

	var num int
	_, err = db.Query(pg.Scan(&num), "SELECT ?", 42)

	cleanDb(db)
	cleanContainer(c)

	require.NoError(t, err)
	require.Equal(t, 42, num)
}
