package service

import (
	"context"
	"errors"
	"in-backend/services/assessment/interfaces"
	"in-backend/services/assessment/models"
	"in-backend/services/assessment/tests/mocks"
	testmodels "in-backend/services/assessment/tests/models"
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

	testCreateAssessment(t, s)
	testGetAllAssessments(t, s)
	testGetAssessmentByID(t, s)
	testUpdateAssessment(t, s)
	testDeleteAssessment(t, s)

	// testCreateAssessmentStatus(t, s)
	// testUpdateAssessmentStatus(t, s)
	// testDeleteAssessmentStatus(t, s)

	// testCreateQuestion(t, s)
	// testGetAllQuestions(t, s)
	// testGetQuestionByID(t, s)
	// testUpdateQuestion(t, s)
	// testDeleteQuestion(t, s)

	// testCreateTag(t, s)
	// testDeleteTag(t, s)

	// testCreateResponse(t, s)
	// testDeleteResponse(t, s)

	r.AssertExpectations(t)
}

/* --------------- Assessment --------------- */

func testCreateAssessment(t *testing.T, s interfaces.Service) {
	testNoName := &testmodels.AssessmentNoName

	test := *testNoName
	test.Name = "javascript"

	type args struct {
		ctx   context.Context
		input *models.Assessment
	}

	type expect struct {
		output *models.Assessment
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"failed not null", args{ctx, testNoName}, expect{nil, errors.New("Failed to insert assessment")}},
		{"valid", args{ctx, &test}, expect{&test, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.On("CreateAssessment", tt.args.ctx, tt.args.input).Return(tt.exp.output, tt.exp.err)

			got, err := s.CreateAssessment(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.NotNil(t, err)
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetAllAssessments(t *testing.T, s interfaces.Service) {
	mockRes := []*models.Assessment{
		{},
		{},
	}

	type args struct {
		ctx   context.Context
		f     *models.AssessmentFilters
		admin bool
	}

	type expect struct {
		output []*models.Assessment
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"nil", args{nil, nil, true}, expect{nil, errors.New("Context cannot be nil")}},
		{"no filter", args{ctx, nil, true}, expect{mockRes, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.On("GetAllAssessments", tt.args.ctx).Return(tt.exp.output, tt.exp.err)

			got, err := s.GetAllAssessments(tt.args.ctx, *tt.args.f, &tt.args.admin)
			assert.Equal(t, len(tt.exp.output), len(got))
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetAssessmentByID(t *testing.T, s interfaces.Service) {
	type args struct {
		ctx   context.Context
		id    uint64
		admin bool
	}

	type expect struct {
		output *models.Assessment
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id 1", args{ctx, 1, true}, expect{&models.Assessment{ID: 1}, nil}},
		{"error", args{ctx, 10000, true}, expect{nil, errors.New("mock error")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.On("GetAssessmentByID", tt.args.ctx, tt.args.id).Return(tt.exp.output, tt.exp.err)

			got, err := s.GetAssessmentByID(tt.args.ctx, tt.args.id, &tt.args.admin)
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

func testUpdateAssessment(t *testing.T, s interfaces.Service) {
	updated := testmodels.AssessmentNoName
	updated.Name = "php"

	type args struct {
		ctx   context.Context
		input *models.Assessment
	}

	type expect struct {
		output *models.Assessment
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"nil", args{ctx, nil}, expect{nil, errors.New("Assessment is nil")}},
		{"id existing", args{ctx, &updated}, expect{&updated, nil}},
		{"id 10000", args{ctx, &models.Assessment{ID: 10000}}, expect{nil, errors.New("Cannot update assessment with id")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.On("UpdateAssessment", tt.args.ctx, tt.args.input).Return(tt.exp.output, tt.exp.err)

			got, err := s.UpdateAssessment(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testDeleteAssessment(t *testing.T, s interfaces.Service) {
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
			r.On("DeleteAssessment", tt.args.ctx, tt.args.id).Return(tt.exp.err)

			err := s.DeleteAssessment(tt.args.ctx, tt.args.id)
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}
