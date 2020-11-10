package database

import (
	"context"
	"in-backend/services/profile/configs"
	"in-backend/services/profile/models"
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
	want := &Repository{
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

	testCreateCandidate(t, r)
	testGetAllCandidates(t, r)
	testGetCandidateByID(t, r)
	testUpdateCandidate(t, r)
	testDeleteCandidate(t, r)

	testCreateSkill(t, r)
	testGetAllSkills(t, r)
	testGetSkill(t, r)

	testCreateInstitution(t, r)
	testGetAllInstitutions(t, r)
	testGetInstitution(t, r)

	testCreateCourse(t, r)
	testGetAllCourses(t, r)
	testGetCourse(t, r)

	testCreateAcademicHistory(t, r)
	testGetAcademicHistory(t, r)
	testUpdateAcademicHistory(t, r)
	testDeleteAcademicHistory(t, r)

	testCreateCompany(t, r)
	testGetAllCompanies(t, r)
	testGetCompany(t, r)

	testCreateDepartment(t, r)
	testGetAllDepartments(t, r)
	testGetDepartment(t, r)

	testCreateJobHistory(t, r)
	testGetJobHistory(t, r)
	testUpdateJobHistory(t, r)
	testDeleteJobHistory(t, r)

	cleanDb(db)
	cleanContainer(c)
}

/* --------------- Candidate --------------- */

func testCreateCandidate(t *testing.T, r *Repository) {
	testNoFirstName := &models.Candidate{
		LastName:      "last",
		Email:         "first@last.com",
		ContactNumber: "+6591234567",
		CreatedAt:     &now,
		UpdatedAt:     &now,
	}

	test := *testNoFirstName
	test.FirstName = "first"

	testDupEmail := test

	// this is required to insert 2 candidates so that one can be used
	// for other tests after the first gets deleted
	test2 := test
	test2.Email = "test@test.com"
	test2.ContactNumber = "+6587654321"

	type args struct {
		ctx   context.Context
		input *models.Candidate
	}

	type expect struct {
		output *models.Candidate
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"nil", args{ctx, nil}, expect{nil, errors.New("Input parameter candidate is nil")}},
		{"failed not null", args{ctx, testNoFirstName}, expect{nil, errors.New("Failed to insert candidate")}},
		{"valid", args{ctx, &test}, expect{&test, nil}},
		{"failed unique", args{ctx, &testDupEmail}, expect{nil, errors.New("Failed to insert candidate")}},
		{"valid2", args{ctx, &test2}, expect{&test2, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.CreateCandidate(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetAllCandidates(t *testing.T, r *Repository) {
	count, err := r.DB.WithContext(ctx).Model((*models.Candidate)(nil)).Count()
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		f   *models.CandidateFilters
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetAllCandidates(tt.args.ctx)
			assert.Equal(t, tt.exp.cnt, len(got))
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetCandidateByID(t *testing.T, r *Repository) {
	existing := &models.Candidate{}
	err := r.DB.WithContext(ctx).Model(existing).First()
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		id  uint64
	}

	type expect struct {
		output *models.Candidate
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id exists", args{ctx, existing.ID}, expect{&models.Candidate{ID: existing.ID}, nil}},
		{"id 10000", args{ctx, 10000}, expect{nil, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetCandidateByID(tt.args.ctx, tt.args.id)
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

func testUpdateCandidate(t *testing.T, r *Repository) {
	existing := &models.Candidate{}
	err := r.DB.WithContext(ctx).Model(existing).First()
	require.NoError(t, err)

	updated := *existing
	updated.FirstName = "new"

	type args struct {
		ctx   context.Context
		input *models.Candidate
	}

	type expect struct {
		output *models.Candidate
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"nil", args{ctx, nil}, expect{nil, errors.New("Candidate is nil")}},
		{"id existing", args{ctx, &updated}, expect{&updated, nil}},
		{"id 10000", args{ctx, &models.Candidate{ID: 10000}}, expect{nil, errors.New("Cannot update candidate with id")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.UpdateCandidate(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testDeleteCandidate(t *testing.T, r *Repository) {
	existing := &models.Candidate{}
	err := r.DB.WithContext(ctx).Model(existing).First()
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
			err := r.DeleteCandidate(tt.args.ctx, tt.args.id)
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

/* --------------- Skill --------------- */

func testCreateSkill(t *testing.T, r *Repository) {
	testNoName := &models.Skill{
		ID: 1,
	}

	test := &models.Skill{
		Name: "skill",
	}

	type args struct {
		ctx   context.Context
		input *models.Skill
	}

	type expect struct {
		output *models.Skill
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"nil", args{ctx, nil}, expect{nil, errors.New("Input parameter skill is nil")}},
		{"failed not null", args{ctx, testNoName}, expect{nil, errors.New("Failed to insert skill")}},
		{"valid", args{ctx, test}, expect{test, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.CreateSkill(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetAllSkills(t *testing.T, r *Repository) {
	count, err := r.DB.WithContext(ctx).Model((*models.Skill)(nil)).Count()
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		f   *models.SkillFilters
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetAllSkills(tt.args.ctx)
			assert.Equal(t, tt.exp.cnt, len(got))
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetSkill(t *testing.T, r *Repository) {
	existing := &models.Skill{}
	err := r.DB.WithContext(ctx).Model(existing).First()
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		id  uint64
	}

	type expect struct {
		output *models.Skill
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id exists", args{ctx, existing.ID}, expect{&models.Skill{ID: existing.ID}, nil}},
		{"id 10000", args{ctx, 10000}, expect{nil, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetSkill(tt.args.ctx, tt.args.id)
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

/* --------------- Institution --------------- */

func testCreateInstitution(t *testing.T, r *Repository) {
	testNoName := &models.Institution{
		ID: 1,
	}

	test := &models.Institution{
		Name:    "institution",
		Country: "singapore",
	}

	type args struct {
		ctx   context.Context
		input *models.Institution
	}

	type expect struct {
		output *models.Institution
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"nil", args{ctx, nil}, expect{nil, errors.New("Input parameter institution is nil")}},
		{"failed not null", args{ctx, testNoName}, expect{nil, errors.New("Failed to insert institution")}},
		{"valid", args{ctx, test}, expect{test, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.CreateInstitution(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetAllInstitutions(t *testing.T, r *Repository) {
	count, err := r.DB.WithContext(ctx).Model((*models.Institution)(nil)).Count()
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		f   *models.InstitutionFilters
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetAllInstitutions(tt.args.ctx)
			assert.Equal(t, tt.exp.cnt, len(got))
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetInstitution(t *testing.T, r *Repository) {
	existing := &models.Institution{}
	err := r.DB.WithContext(ctx).Model(existing).First()
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		id  uint64
	}

	type expect struct {
		output *models.Institution
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id exists", args{ctx, existing.ID}, expect{&models.Institution{ID: existing.ID}, nil}},
		{"id 10000", args{ctx, 10000}, expect{nil, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetInstitution(tt.args.ctx, tt.args.id)
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

/* --------------- Course --------------- */

func testCreateCourse(t *testing.T, r *Repository) {
	testNoName := &models.Course{
		ID: 1,
	}

	test := &models.Course{
		Name:  "course",
		Level: "bachelor",
	}

	type args struct {
		ctx   context.Context
		input *models.Course
	}

	type expect struct {
		output *models.Course
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"nil", args{ctx, nil}, expect{nil, errors.New("Input parameter course is nil")}},
		{"failed not null", args{ctx, testNoName}, expect{nil, errors.New("Failed to insert course")}},
		{"valid", args{ctx, test}, expect{test, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.CreateCourse(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetAllCourses(t *testing.T, r *Repository) {
	count, err := r.DB.WithContext(ctx).Model((*models.Course)(nil)).Count()
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		f   *models.CourseFilters
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetAllCourses(tt.args.ctx)
			assert.Equal(t, tt.exp.cnt, len(got))
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetCourse(t *testing.T, r *Repository) {
	existing := &models.Course{}
	err := r.DB.WithContext(ctx).Model(existing).First()
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		id  uint64
	}

	type expect struct {
		output *models.Course
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id exists", args{ctx, existing.ID}, expect{&models.Course{ID: existing.ID}, nil}},
		{"id 10000", args{ctx, 10000}, expect{nil, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetCourse(tt.args.ctx, tt.args.id)
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

/* --------------- AcademicHistory --------------- */

func testCreateAcademicHistory(t *testing.T, r *Repository) {
	existingC := &models.Candidate{}
	err := r.DB.WithContext(ctx).Model(existingC).First()
	require.NoError(t, err)

	existingI := &models.Institution{}
	err = r.DB.WithContext(ctx).Model(existingI).First()
	require.NoError(t, err)

	existingCr := &models.Course{}
	err = r.DB.WithContext(ctx).Model(existingCr).First()
	require.NoError(t, err)

	testNoCID := &models.AcademicHistory{
		InstitutionID: existingI.ID,
		CourseID:      existingCr.ID,
		CreatedAt:     &now,
		UpdatedAt:     &now,
	}

	testMissingFK := &models.AcademicHistory{
		CandidateID:   10000,
		InstitutionID: 10000,
		CourseID:      10000,
		CreatedAt:     &now,
		UpdatedAt:     &now,
	}

	test := &models.AcademicHistory{
		CandidateID:   existingC.ID,
		InstitutionID: existingI.ID,
		CourseID:      existingCr.ID,
		CreatedAt:     &now,
		UpdatedAt:     &now,
	}

	type args struct {
		ctx   context.Context
		input *models.AcademicHistory
	}

	type expect struct {
		output *models.AcademicHistory
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"nil", args{ctx, nil}, expect{nil, errors.New("Input parameter academic history is nil")}},
		{"failed not null", args{ctx, testNoCID}, expect{nil, errors.New("Failed to insert academic history")}},
		{"failed missing fk", args{ctx, testMissingFK}, expect{nil, errors.New("Failed to insert academic history")}},
		{"valid", args{ctx, test}, expect{test, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.CreateAcademicHistory(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetAcademicHistory(t *testing.T, r *Repository) {
	existing := &models.AcademicHistory{}
	err := r.DB.WithContext(ctx).Model(existing).First()
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		id  uint64
	}

	type expect struct {
		output *models.AcademicHistory
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id exists", args{ctx, existing.ID}, expect{&models.AcademicHistory{ID: existing.ID}, nil}},
		{"id 10000", args{ctx, 10000}, expect{nil, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetAcademicHistory(tt.args.ctx, tt.args.id)
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

func testUpdateAcademicHistory(t *testing.T, r *Repository) {
	existing := &models.AcademicHistory{}
	err := r.DB.WithContext(ctx).Model(existing).First()
	require.NoError(t, err)

	updated := *existing
	updated.YearObtained = 2020

	type args struct {
		ctx   context.Context
		input *models.AcademicHistory
	}

	type expect struct {
		output *models.AcademicHistory
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"nil", args{ctx, nil}, expect{nil, errors.New("AcademicHistory is nil")}},
		{"id existing", args{ctx, &updated}, expect{&updated, nil}},
		{"id 10000", args{ctx, &models.AcademicHistory{ID: 10000}}, expect{nil, errors.New("Cannot update academic history with id")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.UpdateAcademicHistory(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testDeleteAcademicHistory(t *testing.T, r *Repository) {
	existing := &models.AcademicHistory{}
	err := r.DB.WithContext(ctx).Model(existing).First()
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
			err := r.DeleteAcademicHistory(tt.args.ctx, tt.args.id)
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

/* --------------- Company --------------- */

func testCreateCompany(t *testing.T, r *Repository) {
	testNoName := &models.Company{
		ID: 1,
	}

	test := &models.Company{
		Name: "company",
	}

	type args struct {
		ctx   context.Context
		input *models.Company
	}

	type expect struct {
		output *models.Company
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"nil", args{ctx, nil}, expect{nil, errors.New("Input parameter company is nil")}},
		{"failed not null", args{ctx, testNoName}, expect{nil, errors.New("Failed to insert company")}},
		{"valid", args{ctx, test}, expect{test, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.CreateCompany(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetAllCompanies(t *testing.T, r *Repository) {
	count, err := r.DB.WithContext(ctx).Model((*models.Company)(nil)).Count()
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		f   *models.CompanyFilters
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetAllCompanies(tt.args.ctx)
			assert.Equal(t, tt.exp.cnt, len(got))
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetCompany(t *testing.T, r *Repository) {
	existing := &models.Company{}
	err := r.DB.WithContext(ctx).Model(existing).First()
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		id  uint64
	}

	type expect struct {
		output *models.Company
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id exists", args{ctx, existing.ID}, expect{&models.Company{ID: existing.ID}, nil}},
		{"id 10000", args{ctx, 10000}, expect{nil, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetCompany(tt.args.ctx, tt.args.id)
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

/* --------------- Department --------------- */

func testCreateDepartment(t *testing.T, r *Repository) {
	testNoName := &models.Department{
		ID: 1,
	}

	test := &models.Department{
		Name: "department",
	}

	type args struct {
		ctx   context.Context
		input *models.Department
	}

	type expect struct {
		output *models.Department
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"nil", args{ctx, nil}, expect{nil, errors.New("Input parameter department is nil")}},
		{"failed not null", args{ctx, testNoName}, expect{nil, errors.New("Failed to insert department")}},
		{"valid", args{ctx, test}, expect{test, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.CreateDepartment(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetAllDepartments(t *testing.T, r *Repository) {
	count, err := r.DB.WithContext(ctx).Model((*models.Department)(nil)).Count()
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		f   *models.DepartmentFilters
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetAllDepartments(tt.args.ctx)
			assert.Equal(t, tt.exp.cnt, len(got))
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetDepartment(t *testing.T, r *Repository) {
	existing := &models.Department{}
	err := r.DB.WithContext(ctx).Model(existing).First()
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		id  uint64
	}

	type expect struct {
		output *models.Department
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id exists", args{ctx, existing.ID}, expect{&models.Department{ID: existing.ID}, nil}},
		{"id 10000", args{ctx, 10000}, expect{nil, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetDepartment(tt.args.ctx, tt.args.id)
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

/* --------------- JobHistory --------------- */

func testCreateJobHistory(t *testing.T, r *Repository) {
	existingC := &models.Candidate{}
	err := r.DB.WithContext(ctx).Model(existingC).First()
	require.NoError(t, err)

	existingCo := &models.Company{}
	err = r.DB.WithContext(ctx).Model(existingCo).First()
	require.NoError(t, err)

	existingD := &models.Department{}
	err = r.DB.WithContext(ctx).Model(existingD).First()
	require.NoError(t, err)

	start := time.Date(2020, 11, 10, 13, 0, 0, 0, time.Local)

	testNoCID := &models.JobHistory{
		CompanyID:    existingCo.ID,
		DepartmentID: existingD.ID,
		Country:      "singapore",
		Title:        "software engineer",
		StartDate:    &start,
		CreatedAt:    &now,
		UpdatedAt:    &now,
	}

	test := *testNoCID
	test.CandidateID = existingC.ID

	testMissingFK := test
	testMissingFK.CandidateID = 10000
	testMissingFK.CompanyID = 10000
	testMissingFK.DepartmentID = 10000

	type args struct {
		ctx   context.Context
		input *models.JobHistory
	}

	type expect struct {
		output *models.JobHistory
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"nil", args{ctx, nil}, expect{nil, errors.New("Input parameter job history is nil")}},
		{"failed not null", args{ctx, testNoCID}, expect{nil, errors.New("Failed to insert job history")}},
		{"failed missing fk", args{ctx, &testMissingFK}, expect{nil, errors.New("Failed to insert job history")}},
		{"valid", args{ctx, &test}, expect{&test, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.CreateJobHistory(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetJobHistory(t *testing.T, r *Repository) {
	existing := &models.JobHistory{}
	err := r.DB.WithContext(ctx).Model(existing).First()
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		id  uint64
	}

	type expect struct {
		output *models.JobHistory
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id exists", args{ctx, existing.ID}, expect{&models.JobHistory{ID: existing.ID}, nil}},
		{"id 10000", args{ctx, 10000}, expect{nil, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetJobHistory(tt.args.ctx, tt.args.id)
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

func testUpdateJobHistory(t *testing.T, r *Repository) {
	existing := &models.JobHistory{}
	err := r.DB.WithContext(ctx).Model(existing).First()
	require.NoError(t, err)

	updated := *existing
	updated.Country = "indonesia"

	type args struct {
		ctx   context.Context
		input *models.JobHistory
	}

	type expect struct {
		output *models.JobHistory
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"nil", args{ctx, nil}, expect{nil, errors.New("JobHistory is nil")}},
		{"id existing", args{ctx, &updated}, expect{&updated, nil}},
		{"id 10000", args{ctx, &models.JobHistory{ID: 10000}}, expect{nil, errors.New("Cannot update job history with id")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.UpdateJobHistory(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testDeleteJobHistory(t *testing.T, r *Repository) {
	existing := &models.JobHistory{}
	err := r.DB.WithContext(ctx).Model(existing).First()
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
			err := r.DeleteJobHistory(tt.args.ctx, tt.args.id)
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}
