package endpoints

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"in-backend/services/profile/models"
	"in-backend/services/profile/tests/mocks"
)

var (
	svc     *mocks.Service = &mocks.Service{}
	mockCtx                = mock.MatchedBy(func(ctx context.Context) bool { return true })

	ctx context.Context = context.Background()
	now time.Time       = time.Now()
)

func TestMakeEndpoints(t *testing.T) {
	endpoints := MakeEndpoints(svc)
	assert.IsType(t, Endpoints{}, endpoints)
}

/* --------------- Candidate --------------- */

func TestCreateCandidate(t *testing.T) {
	test := &models.Candidate{
		FirstName:     "first",
		LastName:      "last",
		Email:         "first@last.com",
		ContactNumber: "+6591234567",
		CreatedAt:     &now,
		UpdatedAt:     &now,
	}
	in := &CreateCandidateRequest{Candidate: test}
	svcIn := test
	out := test
	svcOut := test

	type args struct {
		ctx      context.Context
		svcInput *models.Candidate
		input    *CreateCandidateRequest
	}

	type expect struct {
		svcOutput *models.Candidate
		output    *models.Candidate
		err       error
	}

	for _, tt := range []struct {
		name string
		args args
		exp  expect
	}{
		{"error", args{ctx, &models.Candidate{}, &CreateCandidateRequest{Candidate: &models.Candidate{}}}, expect{nil, nil, errors.New("Marshal called with nil")}},
		{"valid", args{ctx, svcIn, in}, expect{svcOut, out, nil}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("CreateCandidate", mockCtx, tt.args.svcInput).Return(tt.exp.svcOutput, tt.exp.err)

			res, _ := makeCreateCandidateEndpoint(svc)(tt.args.ctx, *tt.args.input)
			got := res.(CreateCandidateResponse).Candidate
			err := res.(CreateCandidateResponse).Err
			if tt.exp.output != nil {
				assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
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
	svc.AssertExpectations(t)
}

func TestGetAllCandidates(t *testing.T) {
	test := &models.Candidate{
		FirstName:     "first",
		LastName:      "last",
		Email:         "email",
		ContactNumber: "contact",
		CreatedAt:     &now,
		UpdatedAt:     &now,
	}
	mockRes := []*models.Candidate{
		test,
		test,
	}
	svcRes := []*models.Candidate{
		test,
		test,
	}
	in := &GetAllCandidatesRequest{}
	out := mockRes

	// failIn := &GetAllCandidatesRequest{
	// 	ID: []uint64{uint64(1), uint64(2)},
	// }

	type args struct {
		ctx   context.Context
		input *GetAllCandidatesRequest
	}

	type expect struct {
		svcOutput []*models.Candidate
		output    []*models.Candidate
		err       error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		// {"error", args{ctx, failIn}, expect{nil, nil, errors.New("Marshal called with nil")}},
		{"no filter", args{ctx, in}, expect{svcRes, out, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("GetAllCandidates", mockCtx).Return(tt.exp.svcOutput, tt.exp.err)

			res, _ := makeGetAllCandidatesEndpoint(svc)(tt.args.ctx, tt.args.input)
			got := res.(GetAllCandidatesResponse).Candidates
			err := res.(GetAllCandidatesResponse).Err
			if tt.exp.output != nil {
				assert.Equal(t, len(tt.exp.output), len(got))
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
	svc.AssertExpectations(t)
}

func TestGetCandidateByID(t *testing.T) {
	var (
		id     uint64 = 1
		in            = &GetCandidateByIDRequest{ID: id}
		out           = &models.Candidate{ID: id}
		svcOut        = out
	)

	type args struct {
		ctx   context.Context
		id    uint64
		input *GetCandidateByIDRequest
	}

	type expect struct {
		svcOutput *models.Candidate
		output    *models.Candidate
		err       error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id 1", args{ctx, id, in}, expect{svcOut, out, nil}},
		{"error", args{ctx, 10000, &GetCandidateByIDRequest{ID: 10000}}, expect{nil, nil, errors.New("mock error")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("GetCandidateByID", mockCtx, tt.args.id).Return(tt.exp.svcOutput, tt.exp.err)

			res, _ := makeGetCandidateByIDEndpoint(svc)(tt.args.ctx, *tt.args.input)
			got := res.(GetCandidateByIDResponse).Candidate
			err := res.(GetCandidateByIDResponse).Err
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
	svc.AssertExpectations(t)
}

func TestUpdateCandidate(t *testing.T) {
	updated := models.Candidate{
		ID:        1,
		FirstName: "new",
		UpdatedAt: &now,
	}
	svcInOut := &updated
	in := &UpdateCandidateRequest{
		ID:        updated.ID,
		Candidate: &updated,
	}

	type args struct {
		ctx      context.Context
		svcInput *models.Candidate
		input    *UpdateCandidateRequest
	}

	type expect struct {
		svcOutput *models.Candidate
		output    *models.Candidate
		err       error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id existing", args{ctx, svcInOut, in}, expect{svcInOut, &updated, nil}},
		{"id 10000", args{ctx, &models.Candidate{ID: 10000}, &UpdateCandidateRequest{ID: 10000, Candidate: &models.Candidate{ID: 10000}}}, expect{nil, nil, errors.New("Cannot update candidate with id")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("UpdateCandidate", mockCtx, tt.args.svcInput).Return(tt.exp.svcOutput, tt.exp.err)

			res, _ := makeUpdateCandidateEndpoint(svc)(tt.args.ctx, *tt.args.input)
			got := res.(UpdateCandidateResponse).Candidate
			err := res.(UpdateCandidateResponse).Err
			if tt.exp.output != nil {
				assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
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
	svc.AssertExpectations(t)
}

func TestDeleteCandidate(t *testing.T) {
	type args struct {
		ctx   context.Context
		id    uint64
		input *DeleteCandidateRequest
	}

	type expect struct {
		err error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id existing", args{ctx, 1, &DeleteCandidateRequest{ID: 1}}, expect{nil}},
		{"error", args{ctx, 10000, &DeleteCandidateRequest{ID: 10000}}, expect{errors.New("mock error")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("DeleteCandidate", mockCtx, tt.args.id).Return(tt.exp.err)

			res, _ := makeDeleteCandidateEndpoint(svc)(tt.args.ctx, *tt.args.input)
			err := res.(DeleteCandidateResponse).Err
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
	svc.AssertExpectations(t)
}

/* --------------- Skill --------------- */

func TestCreateSkill(t *testing.T) {
	test := &models.Skill{
		Name: "skill",
	}
	in := &CreateSkillRequest{Skill: test}
	svcInOut := test
	out := test

	type args struct {
		ctx      context.Context
		svcInput *models.Skill
		input    *CreateSkillRequest
	}

	type expect struct {
		svcOutput *models.Skill
		output    *models.Skill
		err       error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"error", args{ctx, &models.Skill{}, &CreateSkillRequest{Skill: &models.Skill{}}}, expect{nil, nil, errors.New("Failed to insert skill")}},
		{"valid", args{ctx, svcInOut, in}, expect{svcInOut, out, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("CreateSkill", mockCtx, tt.args.svcInput).Return(tt.exp.svcOutput, tt.exp.err)

			res, _ := makeCreateSkillEndpoint(svc)(tt.args.ctx, *tt.args.input)
			got := res.(CreateSkillResponse).Skill
			err := res.(CreateSkillResponse).Err
			if tt.exp.output != nil {
				assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
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
	svc.AssertExpectations(t)
}

func TestGetAllSkills(t *testing.T) {
	mockRes := []*models.Skill{
		{},
		{},
	}
	svcRes := []*models.Skill{
		{},
		{},
	}
	in := &GetAllSkillsRequest{}
	out := mockRes

	type args struct {
		ctx   context.Context
		input *GetAllSkillsRequest
	}

	type expect struct {
		svcOutput []*models.Skill
		output    []*models.Skill
		err       error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"no filter", args{ctx, in}, expect{svcRes, out, nil}},
		// {"error", args{nil, nil}, expect{nil, nil, errors.New("Context cannot be nil")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("GetAllSkills", mockCtx).Return(tt.exp.svcOutput, tt.exp.err)

			res, _ := makeGetAllSkillsEndpoint(svc)(tt.args.ctx, *tt.args.input)
			got := res.(GetAllSkillsResponse).Skills
			err := res.(GetAllSkillsResponse).Err
			if tt.exp.output != nil {
				assert.Equal(t, len(tt.exp.output), len(got))
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
	svc.AssertExpectations(t)
}

func TestGetSkill(t *testing.T) {
	var (
		id     uint64 = 1
		in            = &GetSkillRequest{ID: id}
		out           = &models.Skill{ID: id}
		svcOut        = out
	)

	type args struct {
		ctx   context.Context
		id    uint64
		input *GetSkillRequest
	}

	type expect struct {
		svcOutput *models.Skill
		output    *models.Skill
		err       error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id exists", args{ctx, id, in}, expect{svcOut, out, nil}},
		{"error", args{ctx, 10000, &GetSkillRequest{ID: 10000}}, expect{nil, nil, errors.New("mock error")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("GetSkill", mockCtx, tt.args.id).Return(tt.exp.svcOutput, tt.exp.err)

			res, _ := makeGetSkillEndpoint(svc)(tt.args.ctx, *tt.args.input)
			got := res.(GetSkillResponse).Skill
			err := res.(GetSkillResponse).Err
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
	svc.AssertExpectations(t)
}

/* --------------- Institution --------------- */

func TestCreateInstitution(t *testing.T) {
	test := &models.Institution{
		Name:    "institution",
		Country: "singapore",
	}
	in := &CreateInstitutionRequest{Institution: test}
	svcInOut := test
	out := test

	type args struct {
		ctx      context.Context
		svcInput *models.Institution
		input    *CreateInstitutionRequest
	}

	type expect struct {
		svcOutput *models.Institution
		output    *models.Institution
		err       error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"error", args{ctx, &models.Institution{}, &CreateInstitutionRequest{Institution: &models.Institution{}}}, expect{nil, nil, errors.New("Failed to insert institution")}},
		{"valid", args{ctx, svcInOut, in}, expect{svcInOut, out, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("CreateInstitution", mockCtx, tt.args.svcInput).Return(tt.exp.svcOutput, tt.exp.err)

			res, _ := makeCreateInstitutionEndpoint(svc)(tt.args.ctx, *tt.args.input)
			got := res.(CreateInstitutionResponse).Institution
			err := res.(CreateInstitutionResponse).Err
			if tt.exp.output != nil {
				assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
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
	svc.AssertExpectations(t)
}

func TestGetAllInstitutions(t *testing.T) {
	mockRes := []*models.Institution{
		{},
		{},
	}
	svcRes := []*models.Institution{
		{},
		{},
	}
	in := &GetAllInstitutionsRequest{}
	out := mockRes

	type args struct {
		ctx   context.Context
		input *GetAllInstitutionsRequest
	}

	type expect struct {
		svcOutput []*models.Institution
		output    []*models.Institution
		err       error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"no filter", args{ctx, in}, expect{svcRes, out, nil}},
		// {"error", args{nil, nil}, expect{nil, nil, errors.New("Context cannot be nil")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("GetAllInstitutions", mockCtx).Return(tt.exp.svcOutput, tt.exp.err)

			res, _ := makeGetAllInstitutionsEndpoint(svc)(tt.args.ctx, *tt.args.input)
			got := res.(GetAllInstitutionsResponse).Institutions
			err := res.(GetAllInstitutionsResponse).Err
			if tt.exp.output != nil {
				assert.Equal(t, len(tt.exp.output), len(got))
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
	svc.AssertExpectations(t)
}

func TestGetInstitution(t *testing.T) {
	var (
		id     uint64 = 1
		in            = &GetInstitutionRequest{ID: id}
		out           = &models.Institution{ID: id}
		svcOut        = out
	)

	type args struct {
		ctx   context.Context
		id    uint64
		input *GetInstitutionRequest
	}

	type expect struct {
		svcOutput *models.Institution
		output    *models.Institution
		err       error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id exists", args{ctx, id, in}, expect{svcOut, out, nil}},
		{"error", args{ctx, 10000, &GetInstitutionRequest{ID: 10000}}, expect{nil, nil, errors.New("mock error")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("GetInstitution", mockCtx, tt.args.id).Return(tt.exp.svcOutput, tt.exp.err)

			res, _ := makeGetInstitutionEndpoint(svc)(tt.args.ctx, *tt.args.input)
			got := res.(GetInstitutionResponse).Institution
			err := res.(GetInstitutionResponse).Err
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
	svc.AssertExpectations(t)
}

/* --------------- Course --------------- */

func TestCreateCourse(t *testing.T) {
	test := &models.Course{
		Name:  "course",
		Level: "bachelor",
	}
	in := &CreateCourseRequest{Course: test}
	svcInOut := test
	out := test

	type args struct {
		ctx      context.Context
		svcInput *models.Course
		input    *CreateCourseRequest
	}

	type expect struct {
		svcOutput *models.Course
		output    *models.Course
		err       error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"error", args{ctx, &models.Course{}, &CreateCourseRequest{Course: &models.Course{}}}, expect{nil, nil, errors.New("Failed to insert course")}},
		{"valid", args{ctx, svcInOut, in}, expect{svcInOut, out, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("CreateCourse", mockCtx, tt.args.svcInput).Return(tt.exp.svcOutput, tt.exp.err)

			res, _ := makeCreateCourseEndpoint(svc)(tt.args.ctx, *tt.args.input)
			got := res.(CreateCourseResponse).Course
			err := res.(CreateCourseResponse).Err
			if tt.exp.output != nil {
				assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
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
	svc.AssertExpectations(t)
}

func TestGetAllCourses(t *testing.T) {
	mockRes := []*models.Course{
		{},
		{},
	}
	svcRes := []*models.Course{
		{},
		{},
	}
	in := &GetAllCoursesRequest{}
	out := mockRes

	type args struct {
		ctx   context.Context
		input *GetAllCoursesRequest
	}

	type expect struct {
		svcOutput []*models.Course
		output    []*models.Course
		err       error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"no filter", args{ctx, in}, expect{svcRes, out, nil}},
		// {"error", args{nil, nil}, expect{nil, nil, errors.New("Context cannot be nil")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("GetAllCourses", mockCtx).Return(tt.exp.svcOutput, tt.exp.err)

			res, _ := makeGetAllCoursesEndpoint(svc)(tt.args.ctx, *tt.args.input)
			got := res.(GetAllCoursesResponse).Courses
			err := res.(GetAllCoursesResponse).Err
			if tt.exp.output != nil {
				assert.Equal(t, len(tt.exp.output), len(got))
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
	svc.AssertExpectations(t)
}

func TestGetCourse(t *testing.T) {
	var (
		id     uint64 = 1
		in            = &GetCourseRequest{ID: id}
		out           = &models.Course{ID: id}
		svcOut        = out
	)

	type args struct {
		ctx   context.Context
		id    uint64
		input *GetCourseRequest
	}

	type expect struct {
		svcOutput *models.Course
		output    *models.Course
		err       error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id exists", args{ctx, id, in}, expect{svcOut, out, nil}},
		{"error", args{ctx, 10000, &GetCourseRequest{ID: 10000}}, expect{nil, nil, errors.New("mock error")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("GetCourse", mockCtx, tt.args.id).Return(tt.exp.svcOutput, tt.exp.err)

			res, _ := makeGetCourseEndpoint(svc)(tt.args.ctx, *tt.args.input)
			got := res.(GetCourseResponse).Course
			err := res.(GetCourseResponse).Err
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
	svc.AssertExpectations(t)
}

/* --------------- AcademicHistory --------------- */

func TestCreateAcademicHistory(t *testing.T) {
	test := &models.AcademicHistory{
		CandidateID:   1,
		InstitutionID: 1,
		CourseID:      1,
		CreatedAt:     &now,
		UpdatedAt:     &now,
	}
	in := &CreateAcademicHistoryRequest{AcademicHistory: test}
	svcIn := test
	out := test
	svcOut := test

	type args struct {
		ctx      context.Context
		svcInput *models.AcademicHistory
		input    *CreateAcademicHistoryRequest
	}

	type expect struct {
		svcOutput *models.AcademicHistory
		output    *models.AcademicHistory
		err       error
	}

	for _, tt := range []struct {
		name string
		args args
		exp  expect
	}{
		{"error", args{ctx, &models.AcademicHistory{}, &CreateAcademicHistoryRequest{AcademicHistory: &models.AcademicHistory{}}}, expect{nil, nil, errors.New("Marshal called with nil")}},
		{"valid", args{ctx, svcIn, in}, expect{svcOut, out, nil}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("CreateAcademicHistory", mockCtx, tt.args.svcInput).Return(tt.exp.svcOutput, tt.exp.err)

			res, _ := makeCreateAcademicHistoryEndpoint(svc)(tt.args.ctx, *tt.args.input)
			got := res.(CreateAcademicHistoryResponse).AcademicHistory
			err := res.(CreateAcademicHistoryResponse).Err
			if tt.exp.output != nil {
				assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
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
	svc.AssertExpectations(t)
}

func TestGetAcademicHistory(t *testing.T) {
	var (
		id     uint64 = 1
		in            = &GetAcademicHistoryRequest{ID: id}
		out           = &models.AcademicHistory{ID: id}
		svcOut        = out
	)

	type args struct {
		ctx   context.Context
		id    uint64
		input *GetAcademicHistoryRequest
	}

	type expect struct {
		svcOutput *models.AcademicHistory
		output    *models.AcademicHistory
		err       error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id 1", args{ctx, id, in}, expect{svcOut, out, nil}},
		{"error", args{ctx, 10000, &GetAcademicHistoryRequest{ID: 10000}}, expect{nil, nil, errors.New("mock error")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("GetAcademicHistory", mockCtx, tt.args.id).Return(tt.exp.svcOutput, tt.exp.err)

			res, _ := makeGetAcademicHistoryEndpoint(svc)(tt.args.ctx, *tt.args.input)
			got := res.(GetAcademicHistoryResponse).AcademicHistory
			err := res.(GetAcademicHistoryResponse).Err
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
	svc.AssertExpectations(t)
}

func TestUpdateAcademicHistory(t *testing.T) {
	updated := models.AcademicHistory{
		ID:           1,
		YearObtained: 2020,
		UpdatedAt:    &now,
	}
	svcInOut := &updated
	in := &UpdateAcademicHistoryRequest{AcademicHistory: &updated}

	type args struct {
		ctx      context.Context
		svcInput *models.AcademicHistory
		input    *UpdateAcademicHistoryRequest
	}

	type expect struct {
		svcOutput *models.AcademicHistory
		output    *models.AcademicHistory
		err       error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id existing", args{ctx, svcInOut, in}, expect{svcInOut, &updated, nil}},
		{"id 10000", args{ctx, &models.AcademicHistory{ID: 10000}, &UpdateAcademicHistoryRequest{AcademicHistory: &models.AcademicHistory{ID: 10000}}}, expect{nil, nil, errors.New("Cannot update academic history with id")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("UpdateAcademicHistory", mockCtx, tt.args.svcInput).Return(tt.exp.svcOutput, tt.exp.err)

			res, _ := makeUpdateAcademicHistoryEndpoint(svc)(tt.args.ctx, *tt.args.input)
			got := res.(UpdateAcademicHistoryResponse).AcademicHistory
			err := res.(UpdateAcademicHistoryResponse).Err
			if tt.exp.output != nil {
				assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
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
	svc.AssertExpectations(t)
}

func TestDeleteAcademicHistory(t *testing.T) {
	type args struct {
		ctx   context.Context
		id    uint64
		input *DeleteAcademicHistoryRequest
	}

	type expect struct {
		err error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id existing", args{ctx, 1, &DeleteAcademicHistoryRequest{CandidateID: 1, AcademicHistoryID: 1}}, expect{nil}},
		{"error", args{ctx, 10000, &DeleteAcademicHistoryRequest{CandidateID: 10000, AcademicHistoryID: 10000}}, expect{errors.New("mock error")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("DeleteAcademicHistory", mockCtx, tt.args.id).Return(tt.exp.err)

			res, _ := makeDeleteAcademicHistoryEndpoint(svc)(tt.args.ctx, *tt.args.input)
			err := res.(DeleteAcademicHistoryResponse).Err
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
	svc.AssertExpectations(t)
}

/* --------------- Company --------------- */

func TestCreateCompany(t *testing.T) {
	test := &models.Company{
		Name: "company",
	}
	in := &CreateCompanyRequest{Company: test}
	svcInOut := test
	out := test

	type args struct {
		ctx      context.Context
		svcInput *models.Company
		input    *CreateCompanyRequest
	}

	type expect struct {
		svcOutput *models.Company
		output    *models.Company
		err       error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"error", args{ctx, &models.Company{}, &CreateCompanyRequest{Company: &models.Company{}}}, expect{nil, nil, errors.New("Failed to insert company")}},
		{"valid", args{ctx, svcInOut, in}, expect{svcInOut, out, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("CreateCompany", mockCtx, tt.args.svcInput).Return(tt.exp.svcOutput, tt.exp.err)

			res, _ := makeCreateCompanyEndpoint(svc)(tt.args.ctx, *tt.args.input)
			got := res.(CreateCompanyResponse).Company
			err := res.(CreateCompanyResponse).Err
			if tt.exp.output != nil {
				assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
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
	svc.AssertExpectations(t)
}

func TestGetAllCompanies(t *testing.T) {
	mockRes := []*models.Company{
		{},
		{},
	}
	svcRes := []*models.Company{
		{},
		{},
	}
	in := &GetAllCompaniesRequest{}
	out := mockRes

	type args struct {
		ctx   context.Context
		input *GetAllCompaniesRequest
	}

	type expect struct {
		svcOutput []*models.Company
		output    []*models.Company
		err       error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"no filter", args{ctx, in}, expect{svcRes, out, nil}},
		// {"error", args{nil, nil}, expect{nil, nil, errors.New("Context cannot be nil")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("GetAllCompanies", mockCtx).Return(tt.exp.svcOutput, tt.exp.err)

			res, _ := makeGetAllCompaniesEndpoint(svc)(tt.args.ctx, *tt.args.input)
			got := res.(GetAllCompaniesResponse).Companies
			err := res.(GetAllCompaniesResponse).Err
			if tt.exp.output != nil {
				assert.Equal(t, len(tt.exp.output), len(got))
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
	svc.AssertExpectations(t)
}

func TestGetCompany(t *testing.T) {
	var (
		id     uint64 = 1
		in            = &GetCompanyRequest{ID: id}
		out           = &models.Company{ID: id}
		svcOut        = out
	)

	type args struct {
		ctx   context.Context
		id    uint64
		input *GetCompanyRequest
	}

	type expect struct {
		svcOutput *models.Company
		output    *models.Company
		err       error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id exists", args{ctx, id, in}, expect{svcOut, out, nil}},
		{"error", args{ctx, 10000, &GetCompanyRequest{ID: 10000}}, expect{nil, nil, errors.New("mock error")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("GetCompany", mockCtx, tt.args.id).Return(tt.exp.svcOutput, tt.exp.err)

			res, _ := makeGetCompanyEndpoint(svc)(tt.args.ctx, *tt.args.input)
			got := res.(GetCompanyResponse).Company
			err := res.(GetCompanyResponse).Err
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
	svc.AssertExpectations(t)
}

/* --------------- Department --------------- */

func TestCreateDepartment(t *testing.T) {
	test := &models.Department{
		Name: "department",
	}
	in := &CreateDepartmentRequest{Department: test}
	svcInOut := test
	out := test

	type args struct {
		ctx      context.Context
		svcInput *models.Department
		input    *CreateDepartmentRequest
	}

	type expect struct {
		svcOutput *models.Department
		output    *models.Department
		err       error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"error", args{ctx, &models.Department{}, &CreateDepartmentRequest{Department: &models.Department{}}}, expect{nil, nil, errors.New("Failed to insert department")}},
		{"valid", args{ctx, svcInOut, in}, expect{svcInOut, out, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("CreateDepartment", mockCtx, tt.args.svcInput).Return(tt.exp.svcOutput, tt.exp.err)

			res, _ := makeCreateDepartmentEndpoint(svc)(tt.args.ctx, *tt.args.input)
			got := res.(CreateDepartmentResponse).Department
			err := res.(CreateDepartmentResponse).Err
			if tt.exp.output != nil {
				assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
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
	svc.AssertExpectations(t)
}

func TestGetAllDepartments(t *testing.T) {
	mockRes := []*models.Department{
		{},
		{},
	}
	svcRes := []*models.Department{
		{},
		{},
	}
	in := &GetAllDepartmentsRequest{}
	out := mockRes

	type args struct {
		ctx   context.Context
		input *GetAllDepartmentsRequest
	}

	type expect struct {
		svcOutput []*models.Department
		output    []*models.Department
		err       error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"no filter", args{ctx, in}, expect{svcRes, out, nil}},
		// {"error", args{nil, nil}, expect{nil, nil, errors.New("Context cannot be nil")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("GetAllDepartments", mockCtx).Return(tt.exp.svcOutput, tt.exp.err)

			res, _ := makeGetAllDepartmentsEndpoint(svc)(tt.args.ctx, *tt.args.input)
			got := res.(GetAllDepartmentsResponse).Departments
			err := res.(GetAllDepartmentsResponse).Err
			if tt.exp.output != nil {
				assert.Equal(t, len(tt.exp.output), len(got))
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
	svc.AssertExpectations(t)
}

func TestGetDepartment(t *testing.T) {
	var (
		id     uint64 = 1
		in            = &GetDepartmentRequest{ID: id}
		out           = &models.Department{ID: id}
		svcOut        = out
	)

	type args struct {
		ctx   context.Context
		id    uint64
		input *GetDepartmentRequest
	}

	type expect struct {
		svcOutput *models.Department
		output    *models.Department
		err       error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id exists", args{ctx, id, in}, expect{svcOut, out, nil}},
		{"error", args{ctx, 10000, &GetDepartmentRequest{ID: 10000}}, expect{nil, nil, errors.New("mock error")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("GetDepartment", mockCtx, tt.args.id).Return(tt.exp.svcOutput, tt.exp.err)

			res, _ := makeGetDepartmentEndpoint(svc)(tt.args.ctx, *tt.args.input)
			got := res.(GetDepartmentResponse).Department
			err := res.(GetDepartmentResponse).Err
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
	svc.AssertExpectations(t)
}

/* --------------- AcademicHistory --------------- */

func TestCreateJobHistory(t *testing.T) {
	test := &models.JobHistory{
		CandidateID:  1,
		CompanyID:    1,
		DepartmentID: 1,
		Country:      "singapore",
		Title:        "software engineer",
		StartDate:    &now,
		CreatedAt:    &now,
		UpdatedAt:    &now,
	}
	in := &CreateJobHistoryRequest{JobHistory: test}
	svcIn := test
	out := test
	svcOut := test

	type args struct {
		ctx      context.Context
		svcInput *models.JobHistory
		input    *CreateJobHistoryRequest
	}

	type expect struct {
		svcOutput *models.JobHistory
		output    *models.JobHistory
		err       error
	}

	for _, tt := range []struct {
		name string
		args args
		exp  expect
	}{
		{"error", args{ctx, &models.JobHistory{}, &CreateJobHistoryRequest{JobHistory: &models.JobHistory{}}}, expect{nil, nil, errors.New("Marshal called with nil")}},
		{"valid", args{ctx, svcIn, in}, expect{svcOut, out, nil}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("CreateJobHistory", mockCtx, tt.args.svcInput).Return(tt.exp.svcOutput, tt.exp.err)

			res, _ := makeCreateJobHistoryEndpoint(svc)(tt.args.ctx, *tt.args.input)
			got := res.(CreateJobHistoryResponse).JobHistory
			err := res.(CreateJobHistoryResponse).Err
			if tt.exp.output != nil {
				assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
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
	svc.AssertExpectations(t)
}

func TestGetJobHistory(t *testing.T) {
	var (
		id     uint64 = 1
		in            = &GetJobHistoryRequest{ID: id}
		out           = &models.JobHistory{ID: id}
		svcOut        = out
	)

	type args struct {
		ctx   context.Context
		id    uint64
		input *GetJobHistoryRequest
	}

	type expect struct {
		svcOutput *models.JobHistory
		output    *models.JobHistory
		err       error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id 1", args{ctx, id, in}, expect{svcOut, out, nil}},
		{"error", args{ctx, 10000, &GetJobHistoryRequest{ID: 10000}}, expect{nil, nil, errors.New("mock error")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("GetJobHistory", mockCtx, tt.args.id).Return(tt.exp.svcOutput, tt.exp.err)

			res, _ := makeGetJobHistoryEndpoint(svc)(tt.args.ctx, *tt.args.input)
			got := res.(GetJobHistoryResponse).JobHistory
			err := res.(GetJobHistoryResponse).Err
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
	svc.AssertExpectations(t)
}

func TestUpdateJobHistory(t *testing.T) {
	updated := models.JobHistory{
		ID:        1,
		Country:   "indonesia",
		UpdatedAt: &now,
	}
	svcInOut := &updated
	in := &UpdateJobHistoryRequest{JobHistory: &updated}

	type args struct {
		ctx      context.Context
		svcInput *models.JobHistory
		input    *UpdateJobHistoryRequest
	}

	type expect struct {
		svcOutput *models.JobHistory
		output    *models.JobHistory
		err       error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id existing", args{ctx, svcInOut, in}, expect{svcInOut, &updated, nil}},
		{"id 10000", args{ctx, &models.JobHistory{ID: 10000}, &UpdateJobHistoryRequest{JobHistory: &models.JobHistory{ID: 10000}}}, expect{nil, nil, errors.New("Cannot update job history with id")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("UpdateJobHistory", mockCtx, tt.args.svcInput).Return(tt.exp.svcOutput, tt.exp.err)

			res, _ := makeUpdateJobHistoryEndpoint(svc)(tt.args.ctx, *tt.args.input)
			got := res.(UpdateJobHistoryResponse).JobHistory
			err := res.(UpdateJobHistoryResponse).Err
			if tt.exp.output != nil {
				assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
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
	svc.AssertExpectations(t)
}

func TestDeleteJobHistory(t *testing.T) {
	type args struct {
		ctx   context.Context
		id    uint64
		input *DeleteJobHistoryRequest
	}

	type expect struct {
		err error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id existing", args{ctx, 1, &DeleteJobHistoryRequest{CandidateID: 1, JobHistoryID: 1}}, expect{nil}},
		{"error", args{ctx, 10000, &DeleteJobHistoryRequest{CandidateID: 10000, JobHistoryID: 10000}}, expect{errors.New("mock error")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("DeleteJobHistory", mockCtx, tt.args.id).Return(tt.exp.err)

			res, _ := makeDeleteJobHistoryEndpoint(svc)(tt.args.ctx, *tt.args.input)
			err := res.(DeleteJobHistoryResponse).Err
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
	svc.AssertExpectations(t)
}
