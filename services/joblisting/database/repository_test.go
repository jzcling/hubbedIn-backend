package database

import (
	"context"
	"in-backend/services/joblisting"
	"in-backend/services/joblisting/configs"
	"in-backend/services/joblisting/models"
	testmodels "in-backend/services/joblisting/tests/models"
	"strings"
	"testing"
	"time"

	pg "github.com/go-pg/pg/v10"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type test func(t *testing.T, r joblisting.Repository, db *pg.DB)

var (
	ctx         context.Context = context.Background()
	now         time.Time       = time.Now()
	testRoutine []test          = []test{
		testCreateJoblisting,
		testGetAllJoblistings,
		testGetJoblistingByID,
		testUpdateJoblisting,
		testDeleteJoblisting,

		testCreateCandidateJoblisting,
		testDeleteCandidateJoblisting,

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

/* --------------- Joblisting --------------- */

func testCreateJoblisting(t *testing.T, r joblisting.Repository, db *pg.DB) {
	testNoName := testmodels.JoblistingNoName
	test := testmodels.JoblistingValid
	testDupRepoURL := test

	type args struct {
		ctx   context.Context
		input *models.Joblisting
	}

	type expect struct {
		output *models.Joblisting
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"nil", args{ctx, nil}, expect{nil, nilErr("joblisting")}},
		{"failed not null", args{ctx, testNoName}, expect{nil, failedToInsertErr(nil, "joblisting", testNoName)}},
		{"valid", args{ctx, test}, expect{test, nil}},
		{"failed unique", args{ctx, testDupRepoURL}, expect{nil, errors.New("Failed to insert joblisting")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.CreateJoblisting(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetAllJoblistings(t *testing.T, r joblisting.Repository, db *pg.DB) {
	count, err := db.WithContext(ctx).Model((*models.Joblisting)(nil)).Count()
	require.NoError(t, err)

	f := &models.JoblistingFilters{
		ID:          []uint64{1, 2},
		CandidateID: 1,
		Name:        "test",
		RepoURL:     "repo",
	}

	type args struct {
		ctx context.Context
		f   *models.JoblistingFilters
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
			got, err := r.GetAllJoblistings(tt.args.ctx, *tt.args.f)
			assert.Equal(t, tt.exp.cnt, len(got))
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetJoblistingByID(t *testing.T, r joblisting.Repository, db *pg.DB) {
	existing := &models.Joblisting{}
	err := db.WithContext(ctx).Model(existing).First()
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		id  uint64
	}

	type expect struct {
		output *models.Joblisting
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id exists", args{ctx, existing.ID}, expect{&models.Joblisting{ID: existing.ID}, nil}},
		{"id 10000", args{ctx, 10000}, expect{nil, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetJoblistingByID(tt.args.ctx, tt.args.id)
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

func testUpdateJoblisting(t *testing.T, r joblisting.Repository, db *pg.DB) {
	existing := &models.Joblisting{}
	err := db.WithContext(ctx).Model(existing).First()
	require.NoError(t, err)

	updated := *existing
	updated.Name = "newname"

	type args struct {
		ctx   context.Context
		input *models.Joblisting
	}

	type expect struct {
		output *models.Joblisting
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"nil", args{ctx, nil}, expect{nil, nilErr("joblisting")}},
		{"id existing", args{ctx, &updated}, expect{&updated, nil}},
		{"id 10000", args{ctx, &models.Joblisting{ID: 10000}}, expect{nil, updateErr(nil, "joblisting", 10000)}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.UpdateJoblisting(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testDeleteJoblisting(t *testing.T, r joblisting.Repository, db *pg.DB) {
	existing := &models.Joblisting{}
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
			err := r.DeleteJoblisting(tt.args.ctx, tt.args.id)
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

/* --------------- Candidate Joblisting --------------- */

func testCreateCandidateJoblisting(t *testing.T, r joblisting.Repository, db *pg.DB) {
	test := testmodels.CandidateJoblistingValid

	type args struct {
		ctx   context.Context
		input *models.CandidateJoblisting
	}

	type expect struct {
		output *models.CandidateJoblisting
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"nil", args{ctx, nil}, expect{nil, nilErr("candidate joblisting")}},
		{"valid", args{ctx, test}, expect{test, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := r.CreateCandidateJoblisting(tt.args.ctx, tt.args.input)
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testDeleteCandidateJoblisting(t *testing.T, r joblisting.Repository, db *pg.DB) {
	existing := &models.CandidateJoblisting{}
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
			err := r.DeleteCandidateJoblisting(tt.args.ctx, tt.args.id)
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

/* --------------- Rating --------------- */

func testCreateRating(t *testing.T, r joblisting.Repository, db *pg.DB) {
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

func testDeleteRating(t *testing.T, r joblisting.Repository, db *pg.DB) {
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
