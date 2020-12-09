package database

import (
	"context"
	"in-backend/services/assessment/configs"
	"in-backend/services/assessment/interfaces"
	"in-backend/services/assessment/models"
	testmodels "in-backend/services/assessment/tests/models"
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

	testCreateAssessment(t, r, db)
	testGetAllAssessments(t, r, db)
	testGetAssessmentByID(t, r, db)
	testUpdateAssessment(t, r, db)
	testDeleteAssessment(t, r, db)

	testCreateAssessmentStatus(t, r, db)
	testUpdateAssessmentStatus(t, r, db)
	testDeleteAssessmentStatus(t, r, db)

	testCreateQuestion(t, r, db)
	testGetAllQuestions(t, r, db)
	testGetQuestionByID(t, r, db)
	testUpdateQuestion(t, r, db)
	testDeleteQuestion(t, r, db)

	testCreateTag(t, r, db)
	testDeleteTag(t, r, db)

	testCreateResponse(t, r, db)
	testDeleteResponse(t, r, db)
}

/* --------------- Assessment --------------- */

func testCreateAssessment(t *testing.T, r interfaces.Repository, db *pg.DB) {
	testNoName := &testmodels.AssessmentNoName

	test := *testNoName
	test.Name = "javascript"

	testDupName := test

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
		{"nil", args{ctx, nil}, expect{nil, errors.New("Input parameter assessment is nil")}},
		{"failed not null", args{ctx, testNoName}, expect{nil, errors.New("Failed to insert assessment")}},
		{"valid", args{ctx, &test}, expect{&test, nil}},
		{"failed unique", args{ctx, &testDupName}, expect{nil, errors.New("Failed to insert assessment")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.CreateAssessment(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetAllAssessments(t *testing.T, r interfaces.Repository, db *pg.DB) {
	count, err := db.WithContext(ctx).Model((*models.Assessment)(nil)).Count()
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		f   *models.AssessmentFilters
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
		{"no filter", args{ctx, &models.AssessmentFilters{}}, expect{count, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetAllAssessments(tt.args.ctx, *tt.args.f)
			assert.Equal(t, tt.exp.cnt, len(got))
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetAssessmentByID(t *testing.T, r interfaces.Repository, db *pg.DB) {
	existing := &models.Assessment{}
	err := db.WithContext(ctx).Model(existing).First()
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		id  uint64
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
		{"id exists", args{ctx, existing.ID}, expect{&models.Assessment{ID: existing.ID}, nil}},
		{"id 10000", args{ctx, 10000}, expect{nil, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetAssessmentByID(tt.args.ctx, tt.args.id)
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

func testUpdateAssessment(t *testing.T, r interfaces.Repository, db *pg.DB) {
	existing := &models.Assessment{}
	err := db.WithContext(ctx).Model(existing).First()
	require.NoError(t, err)

	updated := *existing
	updated.Description = "new"

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
			got, err := r.UpdateAssessment(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testDeleteAssessment(t *testing.T, r interfaces.Repository, db *pg.DB) {
	existing := &models.Assessment{}
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
			err := r.DeleteAssessment(tt.args.ctx, tt.args.id)
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

/* --------------- Assessment Status --------------- */

func testCreateAssessmentStatus(t *testing.T, r interfaces.Repository, db *pg.DB) {
	testNoAssessmentID := &testmodels.AssessmentStatusNoAssessmentID

	test := *testNoAssessmentID
	test.AssessmentID = 1

	type args struct {
		ctx   context.Context
		input *models.AssessmentStatus
	}

	type expect struct {
		output *models.AssessmentStatus
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"nil", args{ctx, nil}, expect{nil, errors.New("Input parameter assessment status is nil")}},
		{"failed not null", args{ctx, testNoAssessmentID}, expect{nil, errors.New("Failed to insert assessment status")}},
		{"valid", args{ctx, &test}, expect{&test, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.CreateAssessmentStatus(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testUpdateAssessmentStatus(t *testing.T, r interfaces.Repository, db *pg.DB) {
	existing := &models.AssessmentStatus{}
	err := db.WithContext(ctx).Model(existing).First()
	require.NoError(t, err)

	updated := *existing
	updated.Status = "Attempted"

	type args struct {
		ctx   context.Context
		input *models.AssessmentStatus
	}

	type expect struct {
		output *models.AssessmentStatus
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"nil", args{ctx, nil}, expect{nil, errors.New("AssessmentStatus is nil")}},
		{"id existing", args{ctx, &updated}, expect{&updated, nil}},
		{"id 10000", args{ctx, &models.AssessmentStatus{ID: 10000}}, expect{nil, errors.New("Cannot update assessment status with id")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.UpdateAssessmentStatus(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testDeleteAssessmentStatus(t *testing.T, r interfaces.Repository, db *pg.DB) {
	existing := &models.AssessmentStatus{}
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
			err := r.DeleteAssessmentStatus(tt.args.ctx, tt.args.id)
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

/* --------------- Question --------------- */

func testCreateQuestion(t *testing.T, r interfaces.Repository, db *pg.DB) {
	testNoType := &testmodels.QuestionNoType

	test := *testNoType
	test.Type = "Multiple Choice"

	type args struct {
		ctx   context.Context
		input *models.Question
	}

	type expect struct {
		output *models.Question
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"nil", args{ctx, nil}, expect{nil, errors.New("Input parameter question is nil")}},
		{"failed not null", args{ctx, testNoType}, expect{nil, errors.New("Failed to insert question")}},
		{"valid", args{ctx, &test}, expect{&test, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.CreateQuestion(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetAllQuestions(t *testing.T, r interfaces.Repository, db *pg.DB) {
	count, err := db.WithContext(ctx).Model((*models.Question)(nil)).Count()
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		f   *models.QuestionFilters
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
		{"no filter", args{ctx, &models.QuestionFilters{}}, expect{count, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetAllQuestions(tt.args.ctx, *tt.args.f)
			assert.Equal(t, tt.exp.cnt, len(got))
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetQuestionByID(t *testing.T, r interfaces.Repository, db *pg.DB) {
	existing := &models.Question{}
	err := db.WithContext(ctx).Model(existing).First()
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		id  uint64
	}

	type expect struct {
		output *models.Question
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id exists", args{ctx, existing.ID}, expect{&models.Question{ID: existing.ID}, nil}},
		{"id 10000", args{ctx, 10000}, expect{nil, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetQuestionByID(tt.args.ctx, tt.args.id)
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

func testUpdateQuestion(t *testing.T, r interfaces.Repository, db *pg.DB) {
	existing := &models.Question{}
	err := db.WithContext(ctx).Model(existing).First()
	require.NoError(t, err)

	updated := *existing
	updated.CreatedBy = 1

	type args struct {
		ctx   context.Context
		input *models.Question
	}

	type expect struct {
		output *models.Question
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"nil", args{ctx, nil}, expect{nil, errors.New("Question is nil")}},
		{"id existing", args{ctx, &updated}, expect{&updated, nil}},
		{"id 10000", args{ctx, &models.Question{ID: 10000}}, expect{nil, errors.New("Cannot update question with id")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.UpdateQuestion(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testDeleteQuestion(t *testing.T, r interfaces.Repository, db *pg.DB) {
	existing := &models.Question{}
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
			err := r.DeleteQuestion(tt.args.ctx, tt.args.id)
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

/* --------------- Tag --------------- */

func testCreateTag(t *testing.T, r interfaces.Repository, db *pg.DB) {
	testNoName := &testmodels.TagNoName

	test := *testNoName
	test.Name = "javascript"

	testDupName := test

	type args struct {
		ctx   context.Context
		input *models.Tag
	}

	type expect struct {
		output *models.Tag
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"nil", args{ctx, nil}, expect{nil, errors.New("Input parameter tag is nil")}},
		{"failed not null", args{ctx, testNoName}, expect{nil, errors.New("Failed to insert tag")}},
		{"valid", args{ctx, &test}, expect{&test, nil}},
		{"failed unique", args{ctx, &testDupName}, expect{nil, errors.New("Failed to insert tag")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.CreateTag(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testDeleteTag(t *testing.T, r interfaces.Repository, db *pg.DB) {
	existing := &models.Tag{}
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
			err := r.DeleteTag(tt.args.ctx, tt.args.id)
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

/* --------------- Response --------------- */

func testCreateResponse(t *testing.T, r interfaces.Repository, db *pg.DB) {
	testNoQuestionID := &testmodels.ResponseNoQuestionID

	test := *testNoQuestionID
	test.QuestionID = 1

	type args struct {
		ctx   context.Context
		input *models.Response
	}

	type expect struct {
		output *models.Response
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"nil", args{ctx, nil}, expect{nil, errors.New("Input parameter response is nil")}},
		{"failed not null", args{ctx, testNoQuestionID}, expect{nil, errors.New("Failed to insert response")}},
		{"valid", args{ctx, &test}, expect{&test, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.CreateResponse(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testDeleteResponse(t *testing.T, r interfaces.Repository, db *pg.DB) {
	existing := &models.Response{}
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
			err := r.DeleteResponse(tt.args.ctx, tt.args.id)
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}
