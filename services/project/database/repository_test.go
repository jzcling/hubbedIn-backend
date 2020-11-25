package database

import (
	"context"
	"fmt"
	"in-backend/services/project"
	"in-backend/services/project/configs"
	"in-backend/services/project/models"
	"strings"
	"testing"
	"time"

	pg "github.com/go-pg/pg/v10"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	ctx context.Context = context.Background()
	now time.Time       = time.Now()
)

func TestNewRepository(t *testing.T) {
	want := &repository{
		DB: &pg.DB{},
	}

	got := NewRepository(&pg.DB{})

	require.EqualValues(t, want, got)
}

func TestAllCRUD(t *testing.T) {
	testConfig, err := configs.LoadConfig(configs.TestFileName)
	require.NoError(t, err)

	opt := GetPgConnectionOptions(testConfig)

	c, err := setupPGContainer(opt)
	require.NoError(t, err)

	db, err := setupDB(c, opt, "../scripts/migrations/")
	require.NoError(t, err)

	r := NewRepository(db)

	// List of all tests to run in test suite
	testCreateProject(t, r, db)

	cleanDb(db)
	cleanContainer(c)
}

/* --------------- Project --------------- */

func testCreateProject(t *testing.T, r project.Repository, db *pg.DB) {
	testNoName := &models.Project{
		RepoURL:   "repo",
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	test := *testNoName
	test.Name = "name"

	testDupRepoURL := test

	type args struct {
		ctx   context.Context
		input *models.Project
	}

	type expect struct {
		output *models.Project
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"nil", args{ctx, nil}, expect{nil, nilErr("project")}},
		{"failed not null", args{ctx, testNoName}, expect{nil, failedToInsertErr("project")}},
		{"valid", args{ctx, &test}, expect{&test, nil}},
		{"failed unique", args{ctx, &testDupRepoURL}, expect{nil, failedToInsertErr("project")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.CreateProject(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func failedToInsertErr(s string) error {
	return errors.New(fmt.Sprintf("Failed to insert %s", s))
}

func nilErr(s string) error {
	return errors.New(fmt.Sprintf("Input parameter %s is nil", s))
}
