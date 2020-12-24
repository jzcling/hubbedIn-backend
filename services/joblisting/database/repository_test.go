package database

import (
	"context"
	"in-backend/services/joblisting/configs"
	"in-backend/services/joblisting/interfaces"
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
	defer cleanDb(db)
	defer cleanContainer(c)
	require.NoError(t, err)

	r := NewRepository(db)

	testCreateJobPost(t, r, db)
	testGetAllJobPosts(t, r, db)
	testGetJobPostByID(t, r, db)
	testUpdateJobPost(t, r, db)
	testDeleteJobPost(t, r, db)
}

/* --------------- Job Post --------------- */

func testCreateJobPost(t *testing.T, r interfaces.Repository, db *pg.DB) {
	testNotNull := &testmodels.JobPostNoTitle

	test := *testNotNull
	test.Title = "software engineer"

	type args struct {
		ctx   context.Context
		input *models.JobPost
	}

	type expect struct {
		output *models.JobPost
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"nil", args{ctx, nil}, expect{nil, errors.New("Input parameter job post is nil")}},
		{"failed not null", args{ctx, testNotNull}, expect{nil, errors.New("Failed to insert job post")}},
		{"valid", args{ctx, &test}, expect{&test, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.CreateJobPost(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetAllJobPosts(t *testing.T, r interfaces.Repository, db *pg.DB) {
	count, err := db.WithContext(ctx).Model((*models.JobPost)(nil)).Count()
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		f   *models.JobPostFilters
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
		{"no filter", args{ctx, &models.JobPostFilters{}}, expect{count, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetAllJobPosts(tt.args.ctx, *tt.args.f)
			assert.Equal(t, tt.exp.cnt, len(got))
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetJobPostByID(t *testing.T, r interfaces.Repository, db *pg.DB) {
	existing := &models.JobPost{}
	err := db.WithContext(ctx).Model(existing).First()
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		id  uint64
	}

	type expect struct {
		output *models.JobPost
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id exists", args{ctx, existing.ID}, expect{&models.JobPost{ID: existing.ID}, nil}},
		{"id 10000", args{ctx, 10000}, expect{nil, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetJobPostByID(tt.args.ctx, tt.args.id)
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

func testUpdateJobPost(t *testing.T, r interfaces.Repository, db *pg.DB) {
	existing := &models.JobPost{}
	err := db.WithContext(ctx).Model(existing).First()
	require.NoError(t, err)

	updated := *existing
	updated.Description = "new"

	type args struct {
		ctx   context.Context
		input *models.JobPost
	}

	type expect struct {
		output *models.JobPost
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"nil", args{ctx, nil}, expect{nil, errors.New("JobPost is nil")}},
		{"id existing", args{ctx, &updated}, expect{&updated, nil}},
		{"id 10000", args{ctx, &models.JobPost{ID: 10000}}, expect{nil, errors.New("Cannot update joblisting with id")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.UpdateJobPost(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testDeleteJobPost(t *testing.T, r interfaces.Repository, db *pg.DB) {
	existing := &models.JobPost{}
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
			err := r.DeleteJobPost(tt.args.ctx, tt.args.id)
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}
