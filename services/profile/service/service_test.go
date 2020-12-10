package service

import (
	"context"
	"errors"
	"in-backend/services/profile/interfaces"
	"in-backend/services/profile/models"
	"in-backend/services/profile/tests/mocks"
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

	testCreateCandidate(t, s)
	testGetAllCandidates(t, s)
	testGetCandidateByID(t, s)
	testUpdateCandidate(t, s)
	testDeleteCandidate(t, s)

	testCreateSkill(t, s)
	testGetAllSkills(t, s)
	testGetSkill(t, s)

	testCreateUserSkill(t, s)
	testDeleteUserSkill(t, s)

	testCreateInstitution(t, s)
	testGetAllInstitutions(t, s)
	testGetInstitution(t, s)

	testCreateCourse(t, s)
	testGetAllCourses(t, s)
	testGetCourse(t, s)

	testCreateAcademicHistory(t, s)
	testGetAcademicHistory(t, s)
	testUpdateAcademicHistory(t, s)
	testDeleteAcademicHistory(t, s)

	testCreateCompany(t, s)
	testGetAllCompanies(t, s)
	testGetCompany(t, s)

	testCreateDepartment(t, s)
	testGetAllDepartments(t, s)
	testGetDepartment(t, s)

	testCreateJobHistory(t, s)
	testGetJobHistory(t, s)
	testUpdateJobHistory(t, s)
	testDeleteJobHistory(t, s)

	r.AssertExpectations(t)
}

/* --------------- Candidate --------------- */

func testCreateCandidate(t *testing.T, s interfaces.Service) {
	testNoFirstName := &models.Candidate{
		LastName:      "last",
		Email:         "first@last.com",
		ContactNumber: "+6591234567",
		CreatedAt:     &now,
		UpdatedAt:     &now,
	}

	test := *testNoFirstName
	test.FirstName = "first"

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
		{"failed not null", args{ctx, testNoFirstName}, expect{nil, errors.New("Failed to insert candidate")}},
		{"valid", args{ctx, &test}, expect{&test, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.On("CreateCandidate", tt.args.ctx, tt.args.input).Return(tt.exp.output, tt.exp.err)

			got, err := s.CreateCandidate(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.NotNil(t, err)
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetAllCandidates(t *testing.T, s interfaces.Service) {
	mockRes := []*models.Candidate{
		{},
		{},
	}

	type args struct {
		ctx context.Context
		f   *models.CandidateFilters
	}

	type expect struct {
		output []*models.Candidate
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
			r.On("GetAllCandidates", tt.args.ctx).Return(tt.exp.output, tt.exp.err)

			got, err := s.GetAllCandidates(tt.args.ctx, *tt.args.f)
			assert.Equal(t, len(tt.exp.output), len(got))
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetCandidateByID(t *testing.T, s interfaces.Service) {
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
		{"id 1", args{ctx, 1}, expect{&models.Candidate{ID: 1}, nil}},
		{"error", args{ctx, 10000}, expect{nil, errors.New("mock error")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.On("GetCandidateByID", tt.args.ctx, tt.args.id).Return(tt.exp.output, tt.exp.err)

			got, err := s.GetCandidateByID(tt.args.ctx, tt.args.id)
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

func testUpdateCandidate(t *testing.T, s interfaces.Service) {
	updated := models.Candidate{
		ID:        1,
		FirstName: "new",
		UpdatedAt: &now,
	}

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
			r.On("UpdateCandidate", tt.args.ctx, tt.args.input).Return(tt.exp.output, tt.exp.err)

			got, err := s.UpdateCandidate(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testDeleteCandidate(t *testing.T, s interfaces.Service) {
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
			r.On("DeleteCandidate", tt.args.ctx, tt.args.id).Return(tt.exp.err)

			err := s.DeleteCandidate(tt.args.ctx, tt.args.id)
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

/* --------------- Skill --------------- */

func testCreateSkill(t *testing.T, s interfaces.Service) {
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
		{"failed not null", args{ctx, testNoName}, expect{nil, errors.New("Failed to insert skill")}},
		{"valid", args{ctx, test}, expect{test, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.On("CreateSkill", tt.args.ctx, tt.args.input).Return(tt.exp.output, tt.exp.err)

			got, err := s.CreateSkill(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetAllSkills(t *testing.T, s interfaces.Service) {
	mockRes := []*models.Skill{
		{},
		{},
	}

	type args struct {
		ctx context.Context
		f   *models.SkillFilters
	}

	type expect struct {
		output []*models.Skill
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"no filter", args{ctx, nil}, expect{mockRes, nil}},
		{"error", args{nil, nil}, expect{nil, errors.New("Context cannot be nil")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.On("GetAllSkills", tt.args.ctx).Return(tt.exp.output, tt.exp.err)

			got, err := s.GetAllSkills(tt.args.ctx, *tt.args.f)
			assert.Equal(t, len(tt.exp.output), len(got))
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetSkill(t *testing.T, s interfaces.Service) {
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
		{"id exists", args{ctx, 1}, expect{&models.Skill{ID: 1}, nil}},
		{"error", args{ctx, 10000}, expect{nil, errors.New("mock error")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.On("GetSkill", tt.args.ctx, tt.args.id).Return(tt.exp.output, tt.exp.err)

			got, err := s.GetSkill(tt.args.ctx, tt.args.id)
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

/* --------------- User Skill --------------- */

func testCreateUserSkill(t *testing.T, s interfaces.Service) {
	testNoCID := &models.UserSkill{
		CandidateID: 1,
	}

	test := &models.UserSkill{
		CandidateID: 1,
		SkillID:     1,
	}

	type args struct {
		ctx   context.Context
		input *models.UserSkill
	}

	type expect struct {
		output *models.UserSkill
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"failed not null", args{ctx, testNoCID}, expect{nil, errors.New("Failed to insert user skill")}},
		{"valid", args{ctx, test}, expect{test, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.On("CreateUserSkill", tt.args.ctx, tt.args.input).Return(tt.exp.output, tt.exp.err)

			got, err := s.CreateUserSkill(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testDeleteUserSkill(t *testing.T, s interfaces.Service) {
	type args struct {
		ctx context.Context
		cid uint64
		sid uint64
	}

	type expect struct {
		err error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id existing", args{ctx, 1, 1}, expect{nil}},
		{"error", args{ctx, 10000, 10000}, expect{errors.New("mock error")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.On("DeleteUserSkill", tt.args.ctx, tt.args.cid, tt.args.sid).Return(tt.exp.err)

			err := s.DeleteUserSkill(tt.args.ctx, tt.args.cid, tt.args.sid)
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

/* --------------- Institution --------------- */

func testCreateInstitution(t *testing.T, s interfaces.Service) {
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
		{"failed not null", args{ctx, testNoName}, expect{nil, errors.New("Failed to insert institution")}},
		{"valid", args{ctx, test}, expect{test, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.On("CreateInstitution", tt.args.ctx, tt.args.input).Return(tt.exp.output, tt.exp.err)

			got, err := s.CreateInstitution(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetAllInstitutions(t *testing.T, s interfaces.Service) {
	mockRes := []*models.Institution{
		{},
		{},
	}

	type args struct {
		ctx context.Context
		f   *models.InstitutionFilters
	}

	type expect struct {
		output []*models.Institution
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"no filter", args{ctx, nil}, expect{mockRes, nil}},
		{"nil", args{nil, nil}, expect{nil, errors.New("Context cannot be nil")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.On("GetAllInstitutions", tt.args.ctx).Return(tt.exp.output, tt.exp.err)

			got, err := s.GetAllInstitutions(tt.args.ctx, *tt.args.f)
			assert.Equal(t, len(tt.exp.output), len(got))
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetInstitution(t *testing.T, s interfaces.Service) {
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
		{"id exists", args{ctx, 1}, expect{&models.Institution{ID: 1}, nil}},
		{"error", args{ctx, 10000}, expect{nil, errors.New("mock error")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.On("GetInstitution", tt.args.ctx, tt.args.id).Return(tt.exp.output, tt.exp.err)

			got, err := s.GetInstitution(tt.args.ctx, tt.args.id)
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

func testCreateCourse(t *testing.T, s interfaces.Service) {
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
		{"failed not null", args{ctx, testNoName}, expect{nil, errors.New("Failed to insert course")}},
		{"valid", args{ctx, test}, expect{test, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.On("CreateCourse", tt.args.ctx, tt.args.input).Return(tt.exp.output, tt.exp.err)

			got, err := s.CreateCourse(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetAllCourses(t *testing.T, s interfaces.Service) {
	mockRes := []*models.Course{
		{},
		{},
	}

	type args struct {
		ctx context.Context
		f   *models.CourseFilters
	}

	type expect struct {
		output []*models.Course
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"no filter", args{ctx, nil}, expect{mockRes, nil}},
		{"nil", args{nil, nil}, expect{nil, errors.New("Context cannot be nil")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.On("GetAllCourses", tt.args.ctx).Return(tt.exp.output, tt.exp.err)

			got, err := s.GetAllCourses(tt.args.ctx, *tt.args.f)
			assert.Equal(t, len(tt.exp.output), len(got))
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetCourse(t *testing.T, s interfaces.Service) {
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
		{"id exists", args{ctx, 1}, expect{&models.Course{ID: 1}, nil}},
		{"error", args{ctx, 10000}, expect{nil, errors.New("mock error")}},
	}

	for _, tt := range tests {
		testUpdateCandidate(t, s)
		t.Run(tt.name, func(t *testing.T) {
			r.On("GetCourse", tt.args.ctx, tt.args.id).Return(tt.exp.output, tt.exp.err)

			got, err := s.GetCourse(tt.args.ctx, tt.args.id)
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

func testCreateAcademicHistory(t *testing.T, s interfaces.Service) {
	testNoCID := &models.AcademicHistory{
		InstitutionID: 1,
		CourseID:      1,
		CreatedAt:     &now,
		UpdatedAt:     &now,
	}

	test := *testNoCID
	test.CandidateID = 1

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
		{"failed not null", args{ctx, testNoCID}, expect{nil, errors.New("Failed to insert academic history")}},
		{"valid", args{ctx, &test}, expect{&test, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.On("CreateAcademicHistory", tt.args.ctx, tt.args.input).Return(tt.exp.output, tt.exp.err)

			got, err := s.CreateAcademicHistory(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetAcademicHistory(t *testing.T, s interfaces.Service) {
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
		{"id exists", args{ctx, 1}, expect{&models.AcademicHistory{ID: 1}, nil}},
		{"error", args{ctx, 10000}, expect{nil, errors.New("mock error")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.On("GetAcademicHistory", tt.args.ctx, tt.args.id).Return(tt.exp.output, tt.exp.err)

			got, err := s.GetAcademicHistory(tt.args.ctx, tt.args.id)
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

func testUpdateAcademicHistory(t *testing.T, s interfaces.Service) {
	updated := models.AcademicHistory{
		ID:           1,
		YearObtained: 2020,
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
		{"id existing", args{ctx, &updated}, expect{&updated, nil}},
		{"id 10000", args{ctx, &models.AcademicHistory{ID: 10000}}, expect{nil, errors.New("Cannot update academic history with id")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.On("UpdateAcademicHistory", tt.args.ctx, tt.args.input).Return(tt.exp.output, tt.exp.err)

			got, err := s.UpdateAcademicHistory(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testDeleteAcademicHistory(t *testing.T, s interfaces.Service) {
	type args struct {
		ctx  context.Context
		cid  uint64
		ahid uint64
	}

	type expect struct {
		err error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id existing", args{ctx, 1, 1}, expect{nil}},
		{"error", args{ctx, 10000, 10000}, expect{errors.New("mock error")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.On("DeleteAcademicHistory", tt.args.ctx, tt.args.cid, tt.args.ahid).Return(tt.exp.err)

			err := s.DeleteAcademicHistory(tt.args.ctx, tt.args.cid, tt.args.ahid)
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

/* --------------- Company --------------- */

func testCreateCompany(t *testing.T, s interfaces.Service) {
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
		{"failed not null", args{ctx, testNoName}, expect{nil, errors.New("Failed to insert company")}},
		{"valid", args{ctx, test}, expect{test, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.On("CreateCompany", tt.args.ctx, tt.args.input).Return(tt.exp.output, tt.exp.err)

			got, err := s.CreateCompany(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetAllCompanies(t *testing.T, s interfaces.Service) {
	mockRes := []*models.Company{
		{},
		{},
	}

	type args struct {
		ctx context.Context
		f   *models.CompanyFilters
	}

	type expect struct {
		output []*models.Company
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"no filter", args{ctx, nil}, expect{mockRes, nil}},
		{"nil", args{nil, nil}, expect{nil, errors.New("Context cannot be nil")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.On("GetAllCompanies", tt.args.ctx).Return(tt.exp.output, tt.exp.err)

			got, err := s.GetAllCompanies(tt.args.ctx, *tt.args.f)
			assert.Equal(t, len(tt.exp.output), len(got))
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetCompany(t *testing.T, s interfaces.Service) {
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
		{"id exists", args{ctx, 1}, expect{&models.Company{ID: 1}, nil}},
		{"error", args{ctx, 10000}, expect{nil, errors.New("mock error")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.On("GetCompany", tt.args.ctx, tt.args.id).Return(tt.exp.output, tt.exp.err)

			got, err := s.GetCompany(tt.args.ctx, tt.args.id)
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

func testCreateDepartment(t *testing.T, s interfaces.Service) {
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
		{"failed not null", args{ctx, testNoName}, expect{nil, errors.New("Failed to insert department")}},
		{"valid", args{ctx, test}, expect{test, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.On("CreateDepartment", tt.args.ctx, tt.args.input).Return(tt.exp.output, tt.exp.err)

			got, err := s.CreateDepartment(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetAllDepartments(t *testing.T, s interfaces.Service) {
	mockRes := []*models.Department{
		{},
		{},
	}

	type args struct {
		ctx context.Context
		f   *models.DepartmentFilters
	}

	type expect struct {
		output []*models.Department
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"no filter", args{ctx, nil}, expect{mockRes, nil}},
		{"nil", args{nil, nil}, expect{nil, errors.New("Context cannot be nil")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.On("GetAllDepartments", tt.args.ctx).Return(tt.exp.output, tt.exp.err)

			got, err := s.GetAllDepartments(tt.args.ctx, *tt.args.f)
			assert.Equal(t, len(tt.exp.output), len(got))
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetDepartment(t *testing.T, s interfaces.Service) {
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
		{"id exists", args{ctx, 1}, expect{&models.Department{ID: 1}, nil}},
		{"error", args{ctx, 10000}, expect{nil, errors.New("mock error")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.On("GetDepartment", tt.args.ctx, tt.args.id).Return(tt.exp.output, tt.exp.err)

			got, err := s.GetDepartment(tt.args.ctx, tt.args.id)
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

func testCreateJobHistory(t *testing.T, s interfaces.Service) {
	start := time.Date(2020, 11, 10, 13, 0, 0, 0, time.Local)

	testNoCID := &models.JobHistory{
		CompanyID:    1,
		DepartmentID: 1,
		Country:      "singapore",
		Title:        "software engineer",
		StartDate:    &start,
		CreatedAt:    &now,
		UpdatedAt:    &now,
	}

	test := *testNoCID
	test.CandidateID = 1

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
		{"failed not null", args{ctx, testNoCID}, expect{nil, errors.New("Failed to insert job history")}},
		{"valid", args{ctx, &test}, expect{&test, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.On("CreateJobHistory", tt.args.ctx, tt.args.input).Return(tt.exp.output, tt.exp.err)

			got, err := s.CreateJobHistory(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetJobHistory(t *testing.T, s interfaces.Service) {
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
		{"id exists", args{ctx, 1}, expect{&models.JobHistory{ID: 1}, nil}},
		{"error", args{ctx, 10000}, expect{nil, errors.New("mock error")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.On("GetJobHistory", tt.args.ctx, tt.args.id).Return(tt.exp.output, tt.exp.err)

			got, err := s.GetJobHistory(tt.args.ctx, tt.args.id)
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

func testUpdateJobHistory(t *testing.T, s interfaces.Service) {
	updated := models.JobHistory{
		ID:      1,
		Country: "indonesia",
	}

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
		{"id existing", args{ctx, &updated}, expect{&updated, nil}},
		{"id 10000", args{ctx, &models.JobHistory{ID: 10000}}, expect{nil, errors.New("Cannot update job history with id")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.On("UpdateJobHistory", tt.args.ctx, tt.args.input).Return(tt.exp.output, tt.exp.err)

			got, err := s.UpdateJobHistory(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testDeleteJobHistory(t *testing.T, s interfaces.Service) {
	type args struct {
		ctx  context.Context
		cid  uint64
		jhid uint64
	}

	type expect struct {
		err error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id existing", args{ctx, 1, 1}, expect{nil}},
		{"error", args{ctx, 10000, 10000}, expect{errors.New("mock error")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.On("DeleteJobHistory", tt.args.ctx, tt.args.cid, tt.args.jhid).Return(tt.exp.err)

			err := s.DeleteJobHistory(tt.args.ctx, tt.args.cid, tt.args.jhid)
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}
