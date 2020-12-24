package service

import (
	"context"
	"errors"
	"in-backend/services/joblisting/interfaces"
	"in-backend/services/joblisting/models"
	"in-backend/services/joblisting/tests/mocks"
	testmodels "in-backend/services/joblisting/tests/models"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	r *mocks.Repository = &mocks.Repository{}
	w io.Writer         = log.NewSyncWriter(os.Stderr)

	ctx context.Context = context.Background()
	now time.Time       = time.Now()
)

func TestNew(t *testing.T) {
	expect := &service{
		repository: r,
	}

	got := New(r)

	require.Equal(t, expect, got)
}

func TestAllCRUD(t *testing.T) {
	s := New(r)

	testCreateJobPost(t, s)
	testGetAllJobPosts(t, s)
	testGetJobPostByID(t, s)
	testUpdateJobPost(t, s)
	testDeleteJobPost(t, s)

	r.AssertExpectations(t)
}

/* --------------- Job Post --------------- */

func testCreateJobPost(t *testing.T, s interfaces.Service) {
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
		{"failed not null", args{ctx, testNotNull}, expect{nil, errors.New("Failed to insert joblisting")}},
		{"valid", args{ctx, &test}, expect{&test, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.On("CreateJobPost", tt.args.ctx, tt.args.input).Return(tt.exp.output, tt.exp.err)

			got, err := s.CreateJobPost(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.NotNil(t, err)
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetAllJobPosts(t *testing.T, s interfaces.Service) {
	mockRes := []*models.JobPost{
		{},
		{},
	}

	type args struct {
		ctx context.Context
		f   *models.JobPostFilters
	}

	type expect struct {
		output []*models.JobPost
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"nil", args{nil, nil}, expect{nil, errors.New("Context cannot be nil")}},
		{"no filter", args{ctx, nil}, expect{mockRes, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.On("GetAllJobPosts", tt.args.ctx).Return(tt.exp.output, tt.exp.err)

			got, err := s.GetAllJobPosts(tt.args.ctx, *tt.args.f)
			assert.Equal(t, len(tt.exp.output), len(got))
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetJobPostByID(t *testing.T, s interfaces.Service) {
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
		{"id 1", args{ctx, 1}, expect{&models.JobPost{ID: 1}, nil}},
		{"error", args{ctx, 10000}, expect{nil, errors.New("mock error")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.On("GetJobPostByID", tt.args.ctx, tt.args.id).Return(tt.exp.output, tt.exp.err)

			got, err := s.GetJobPostByID(tt.args.ctx, tt.args.id)
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

func testUpdateJobPost(t *testing.T, s interfaces.Service) {
	updated := testmodels.JobPostNoTitle
	updated.Title = "software engineer"

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
			r.On("UpdateJobPost", tt.args.ctx, tt.args.input).Return(tt.exp.output, tt.exp.err)

			got, err := s.UpdateJobPost(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testDeleteJobPost(t *testing.T, s interfaces.Service) {
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
		{"id existing", args{ctx, 1}, expect{nil}},
		{"error", args{ctx, 10000}, expect{errors.New("mock error")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.On("DeleteJobPost", tt.args.ctx, tt.args.id).Return(tt.exp.err)

			err := s.DeleteJobPost(tt.args.ctx, tt.args.id)
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}
