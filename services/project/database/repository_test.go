package database

import (
	"context"
	"in-backend/services/project"
	"in-backend/services/project/configs"
	"in-backend/services/project/models"
	testmodels "in-backend/services/project/tests/models"
	"strings"
	"testing"
	"time"

	pg "github.com/go-pg/pg/v10"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type test func(t *testing.T, r project.Repository, db *pg.DB)

var (
	ctx         context.Context = context.Background()
	now         time.Time       = time.Now()
	testRoutine []test          = []test{
		testCreateProject,
		testGetAllProjects,
		testGetProjectByID,
		testUpdateProject,
		testDeleteProject,

		testCreateCandidateProject,
		testDeleteCandidateProject,

		testCreateRating,
		testDeleteRating,
	}
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

	db, err := setupDB(c, opt, "../tests/migrations/")
	require.NoError(t, err)

	r := NewRepository(db)

	// run all tests in test suite
	for _, test := range testRoutine {
		tx, err := db.Begin()
		require.NoError(t, err)

		test(t, r, db)

		tx.Rollback()
	}

	cleanDb(db)
	cleanContainer(c)
}

/* --------------- Project --------------- */

func testCreateProject(t *testing.T, r project.Repository, db *pg.DB) {
	testNoName := testmodels.ProjectNoName
	test := testmodels.ProjectValid
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
		{"failed not null", args{ctx, testNoName}, expect{nil, failedToInsertErr(nil, "project", testNoName)}},
		{"valid", args{ctx, test}, expect{test, nil}},
		{"failed unique", args{ctx, testDupRepoURL}, expect{nil, errors.New("Failed to insert project")}},
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

func testGetAllProjects(t *testing.T, r project.Repository, db *pg.DB) {
	count, err := db.WithContext(ctx).Model((*models.Project)(nil)).Count()
	require.NoError(t, err)

	f := &models.ProjectFilters{
		ID:          []uint64{1, 2},
		CandidateID: 1,
		Name:        "test",
		RepoURL:     "repo",
	}

	type args struct {
		ctx context.Context
		f   *models.ProjectFilters
	}

	type expect struct {
		cnt int
		err error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"no filter", args{ctx, nil}, expect{count, nil}},
		{"all filters", args{ctx, f}, expect{2, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetAllProjects(tt.args.ctx, *tt.args.f)
			assert.Equal(t, tt.exp.cnt, len(got))
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetProjectByID(t *testing.T, r project.Repository, db *pg.DB) {
	existing := &models.Project{}
	err := db.WithContext(ctx).Model(existing).First()
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		id  uint64
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
		{"id exists", args{ctx, existing.ID}, expect{&models.Project{ID: existing.ID}, nil}},
		{"id 10000", args{ctx, 10000}, expect{nil, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetProjectByID(tt.args.ctx, tt.args.id)
			if tt.exp.output != nil && got != nil {
				assert.Equal(t, tt.exp.output.ID, got.ID)
			} else {
				assert.Equal(t, tt.exp.output, got)
			}
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testUpdateProject(t *testing.T, r project.Repository, db *pg.DB) {
	existing := &models.Project{}
	err := db.WithContext(ctx).Model(existing).First()
	require.NoError(t, err)

	updated := *existing
	updated.Name = "newname"

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
		{"id existing", args{ctx, &updated}, expect{&updated, nil}},
		{"id 10000", args{ctx, &models.Project{ID: 10000}}, expect{nil, updateErr(nil, "project", 10000)}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.UpdateProject(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testDeleteProject(t *testing.T, r project.Repository, db *pg.DB) {
	existing := &models.Project{}
	err := db.WithContext(ctx).Model(existing).First()
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		id  uint64
	}

	type expect struct {
		err error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id existing", args{ctx, existing.ID}, expect{nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := r.DeleteProject(tt.args.ctx, tt.args.id)
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

/* --------------- Candidate Project --------------- */

func testCreateCandidateProject(t *testing.T, r project.Repository, db *pg.DB) {
	test := testmodels.CandidateProjectValid

	type args struct {
		ctx   context.Context
		input *models.CandidateProject
	}

	type expect struct {
		output *models.CandidateProject
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"nil", args{ctx, nil}, expect{nil, nilErr("candidate project")}},
		{"valid", args{ctx, test}, expect{test, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := r.CreateCandidateProject(tt.args.ctx, tt.args.input)
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testDeleteCandidateProject(t *testing.T, r project.Repository, db *pg.DB) {
	existing := &models.CandidateProject{}
	err := db.WithContext(ctx).Model(existing).First()
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		id  uint64
	}

	type expect struct {
		err error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id existing", args{ctx, existing.ID}, expect{nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := r.DeleteCandidateProject(tt.args.ctx, tt.args.id)
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

/* --------------- Rating --------------- */

func testCreateRating(t *testing.T, r project.Repository, db *pg.DB) {
	test := testmodels.RatingValid

	type args struct {
		ctx   context.Context
		input *models.Rating
	}

	type expect struct {
		output *models.Rating
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"nil", args{ctx, nil}, expect{nil, nilErr("rating")}},
		{"valid", args{ctx, test}, expect{test, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := r.CreateRating(tt.args.ctx, tt.args.input)
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testDeleteRating(t *testing.T, r project.Repository, db *pg.DB) {
	existing := &models.Rating{}
	err := db.WithContext(ctx).Model(existing).First()
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		id  uint64
	}

	type expect struct {
		err error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id existing", args{ctx, existing.ID}, expect{nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := r.DeleteRating(tt.args.ctx, tt.args.id)
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}
