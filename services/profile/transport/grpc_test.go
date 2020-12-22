package transport

import (
	"context"
	"errors"
	"net"
	"os"
	"strings"
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	"in-backend/services/profile/endpoints"
	"in-backend/services/profile/models"
	"in-backend/services/profile/pb"
	"in-backend/services/profile/tests/mocks"
)

const bufSize = 1024 * 1024

var (
	lis     *bufconn.Listener
	ctx     context.Context = context.Background()
	svc     *mocks.Service  = &mocks.Service{}
	mockCtx                 = mock.MatchedBy(func(ctx context.Context) bool { return true })
)

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	endpoints := endpoints.MakeEndpoints(svc)
	logger := log.NewLogfmtLogger(os.Stderr)
	server := NewGRPCServer(endpoints, nil, logger)
	pb.RegisterProfileServiceServer(s, server)
	go func() {
		if err := s.Serve(lis); err != nil {
			level.Error(logger).Log("GRPCListener", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func setupClient(t *testing.T) (*grpc.ClientConn, pb.ProfileServiceClient) {
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	require.NoError(t, err)
	return conn, pb.NewProfileServiceClient(conn)
}

/* --------------- Candidate --------------- */

func TestCreateCandidate(t *testing.T) {
	conn, client := setupClient(t)
	defer conn.Close()

	testPbTime := ptypes.TimestampNow()
	test := &pb.Candidate{
		FirstName:     "first",
		LastName:      "last",
		Email:         "email",
		ContactNumber: "contact",
		CreatedAt:     testPbTime,
		UpdatedAt:     testPbTime,
	}
	in := &pb.CreateCandidateRequest{Candidate: test}
	svcIn := models.CandidateToORM(test)
	out := test
	svcOut := models.CandidateToORM(test)

	type args struct {
		ctx      context.Context
		svcInput *models.Candidate
		input    *pb.CreateCandidateRequest
	}

	type expect struct {
		svcOutput *models.Candidate
		output    *pb.Candidate
		err       error
	}

	for _, tt := range []struct {
		name string
		args args
		exp  expect
	}{
		{"error", args{ctx, &models.Candidate{}, &pb.CreateCandidateRequest{Candidate: &pb.Candidate{}}}, expect{nil, nil, errors.New("Marshal called with nil")}},
		{"valid", args{ctx, svcIn, in}, expect{svcOut, out, nil}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("CreateCandidate", mockCtx, tt.args.svcInput).Return(tt.exp.svcOutput, tt.exp.err)

			got, err := client.CreateCandidate(ctx, tt.args.input)
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
	conn, client := setupClient(t)
	defer conn.Close()

	testPbTime := ptypes.TimestampNow()
	test := &pb.Candidate{
		FirstName:     "first",
		LastName:      "last",
		Email:         "email",
		ContactNumber: "contact",
		CreatedAt:     testPbTime,
		UpdatedAt:     testPbTime,
	}
	mockRes := []*pb.Candidate{
		test,
		test,
	}
	svcRes := []*models.Candidate{
		models.CandidateToORM(test),
		models.CandidateToORM(test),
	}
	in := &pb.GetAllCandidatesRequest{}
	out := &pb.GetAllCandidatesResponse{Candidates: mockRes}

	// failIn := &pb.GetAllCandidatesRequest{
	// 	Id: []uint64{uint64(1), uint64(2)},
	// }

	type args struct {
		ctx   context.Context
		input *pb.GetAllCandidatesRequest
	}

	type expect struct {
		svcOutput []*models.Candidate
		output    *pb.GetAllCandidatesResponse
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

			got, err := client.GetAllCandidates(tt.args.ctx, tt.args.input)
			if tt.exp.output != nil {
				assert.Equal(t, len(tt.exp.output.Candidates), len(got.Candidates))
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
	conn, client := setupClient(t)
	defer conn.Close()

	var (
		id     uint64 = 1
		in            = &pb.GetCandidateByIDRequest{Id: id}
		out           = &pb.Candidate{Id: id}
		svcOut        = models.CandidateToORM(out)
	)

	type args struct {
		ctx   context.Context
		id    uint64
		input *pb.GetCandidateByIDRequest
	}

	type expect struct {
		svcOutput *models.Candidate
		output    *pb.Candidate
		err       error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id 1", args{ctx, id, in}, expect{svcOut, out, nil}},
		{"error", args{ctx, 10000, &pb.GetCandidateByIDRequest{Id: 10000}}, expect{nil, nil, errors.New("mock error")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("GetCandidateByID", mockCtx, tt.args.id).Return(tt.exp.svcOutput, tt.exp.err)

			got, err := client.GetCandidateByID(tt.args.ctx, tt.args.input)
			if tt.exp.output != nil && got != nil {
				assert.Equal(t, tt.exp.output.Id, got.Id)
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
	conn, client := setupClient(t)
	defer conn.Close()

	updated := pb.Candidate{
		Id:        1,
		FirstName: "new",
		UpdatedAt: ptypes.TimestampNow(),
	}
	svcInOut := models.CandidateToORM(&updated)
	in := &pb.UpdateCandidateRequest{Id: updated.Id, Candidate: &updated}

	type args struct {
		ctx      context.Context
		svcInput *models.Candidate
		input    *pb.UpdateCandidateRequest
	}

	type expect struct {
		svcOutput *models.Candidate
		output    *pb.Candidate
		err       error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id existing", args{ctx, svcInOut, in}, expect{svcInOut, &updated, nil}},
		{"id 10000", args{ctx, &models.Candidate{ID: 10000}, &pb.UpdateCandidateRequest{Id: 10000, Candidate: &pb.Candidate{Id: 10000}}}, expect{nil, nil, errors.New("Cannot update candidate with id")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("UpdateCandidate", mockCtx, tt.args.svcInput).Return(tt.exp.svcOutput, tt.exp.err)

			got, err := client.UpdateCandidate(tt.args.ctx, tt.args.input)
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
	conn, client := setupClient(t)
	defer conn.Close()

	type args struct {
		ctx   context.Context
		id    uint64
		input *pb.DeleteCandidateRequest
	}

	type expect struct {
		err error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id existing", args{ctx, 1, &pb.DeleteCandidateRequest{Id: 1}}, expect{nil}},
		{"error", args{ctx, 10000, &pb.DeleteCandidateRequest{Id: 10000}}, expect{errors.New("mock error")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("DeleteCandidate", mockCtx, tt.args.id).Return(tt.exp.err)

			_, err := client.DeleteCandidate(tt.args.ctx, tt.args.input)
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
	conn, client := setupClient(t)
	defer conn.Close()

	test := &pb.Skill{
		Name: "skill",
	}
	in := &pb.CreateSkillRequest{Skill: test}
	svcInOut := models.SkillToORM(test)
	out := test

	type args struct {
		ctx      context.Context
		svcInput *models.Skill
		input    *pb.CreateSkillRequest
	}

	type expect struct {
		svcOutput *models.Skill
		output    *pb.Skill
		err       error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"error", args{ctx, &models.Skill{}, &pb.CreateSkillRequest{Skill: &pb.Skill{}}}, expect{nil, nil, errors.New("Failed to insert skill")}},
		{"valid", args{ctx, svcInOut, in}, expect{svcInOut, out, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("CreateSkill", mockCtx, tt.args.svcInput).Return(tt.exp.svcOutput, tt.exp.err)

			got, err := client.CreateSkill(tt.args.ctx, tt.args.input)
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
	conn, client := setupClient(t)
	defer conn.Close()

	mockRes := []*pb.Skill{
		{},
		{},
	}
	svcRes := []*models.Skill{
		{},
		{},
	}
	in := &pb.GetAllSkillsRequest{}
	out := &pb.GetAllSkillsResponse{Skills: mockRes}

	type args struct {
		ctx   context.Context
		input *pb.GetAllSkillsRequest
	}

	type expect struct {
		svcOutput []*models.Skill
		output    *pb.GetAllSkillsResponse
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

			got, err := client.GetAllSkills(tt.args.ctx, tt.args.input)
			if tt.exp.output != nil {
				assert.Equal(t, len(tt.exp.output.Skills), len(got.Skills))
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
	conn, client := setupClient(t)
	defer conn.Close()

	var (
		id     uint64 = 1
		in            = &pb.GetSkillRequest{Id: id}
		out           = &pb.Skill{Id: id}
		svcOut        = models.SkillToORM(out)
	)

	type args struct {
		ctx   context.Context
		id    uint64
		input *pb.GetSkillRequest
	}

	type expect struct {
		svcOutput *models.Skill
		output    *pb.Skill
		err       error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id exists", args{ctx, id, in}, expect{svcOut, out, nil}},
		{"error", args{ctx, 10000, &pb.GetSkillRequest{Id: 10000}}, expect{nil, nil, errors.New("mock error")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("GetSkill", mockCtx, tt.args.id).Return(tt.exp.svcOutput, tt.exp.err)

			got, err := client.GetSkill(tt.args.ctx, tt.args.input)
			if tt.exp.output != nil && got != nil {
				assert.Equal(t, tt.exp.output.Id, got.Id)
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
	conn, client := setupClient(t)
	defer conn.Close()

	test := &pb.Institution{
		Name:    "institution",
		Country: "singapore",
	}
	in := &pb.CreateInstitutionRequest{Institution: test}
	svcInOut := models.InstitutionToORM(test)
	out := test

	type args struct {
		ctx      context.Context
		svcInput *models.Institution
		input    *pb.CreateInstitutionRequest
	}

	type expect struct {
		svcOutput *models.Institution
		output    *pb.Institution
		err       error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"error", args{ctx, &models.Institution{}, &pb.CreateInstitutionRequest{Institution: &pb.Institution{}}}, expect{nil, nil, errors.New("Failed to insert institution")}},
		{"valid", args{ctx, svcInOut, in}, expect{svcInOut, out, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("CreateInstitution", mockCtx, tt.args.svcInput).Return(tt.exp.svcOutput, tt.exp.err)

			got, err := client.CreateInstitution(tt.args.ctx, tt.args.input)
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
	conn, client := setupClient(t)
	defer conn.Close()

	mockRes := []*pb.Institution{
		{},
		{},
	}
	svcRes := []*models.Institution{
		{},
		{},
	}
	in := &pb.GetAllInstitutionsRequest{}
	out := &pb.GetAllInstitutionsResponse{Institutions: mockRes}

	type args struct {
		ctx   context.Context
		input *pb.GetAllInstitutionsRequest
	}

	type expect struct {
		svcOutput []*models.Institution
		output    *pb.GetAllInstitutionsResponse
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

			got, err := client.GetAllInstitutions(tt.args.ctx, tt.args.input)
			if tt.exp.output != nil {
				assert.Equal(t, len(tt.exp.output.Institutions), len(got.Institutions))
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
	conn, client := setupClient(t)
	defer conn.Close()

	var (
		id     uint64 = 1
		in            = &pb.GetInstitutionRequest{Id: id}
		out           = &pb.Institution{Id: id}
		svcOut        = models.InstitutionToORM(out)
	)

	type args struct {
		ctx   context.Context
		id    uint64
		input *pb.GetInstitutionRequest
	}

	type expect struct {
		svcOutput *models.Institution
		output    *pb.Institution
		err       error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id exists", args{ctx, id, in}, expect{svcOut, out, nil}},
		{"error", args{ctx, 10000, &pb.GetInstitutionRequest{Id: 10000}}, expect{nil, nil, errors.New("mock error")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("GetInstitution", mockCtx, tt.args.id).Return(tt.exp.svcOutput, tt.exp.err)

			got, err := client.GetInstitution(tt.args.ctx, tt.args.input)
			if tt.exp.output != nil && got != nil {
				assert.Equal(t, tt.exp.output.Id, got.Id)
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
	conn, client := setupClient(t)
	defer conn.Close()

	test := &pb.Course{
		Name:  "course",
		Level: "bachelor",
	}
	in := &pb.CreateCourseRequest{Course: test}
	svcInOut := models.CourseToORM(test)
	out := test

	type args struct {
		ctx      context.Context
		svcInput *models.Course
		input    *pb.CreateCourseRequest
	}

	type expect struct {
		svcOutput *models.Course
		output    *pb.Course
		err       error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"error", args{ctx, &models.Course{}, &pb.CreateCourseRequest{Course: &pb.Course{}}}, expect{nil, nil, errors.New("Failed to insert course")}},
		{"valid", args{ctx, svcInOut, in}, expect{svcInOut, out, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("CreateCourse", mockCtx, tt.args.svcInput).Return(tt.exp.svcOutput, tt.exp.err)

			got, err := client.CreateCourse(tt.args.ctx, tt.args.input)
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
	conn, client := setupClient(t)
	defer conn.Close()

	mockRes := []*pb.Course{
		{},
		{},
	}
	svcRes := []*models.Course{
		{},
		{},
	}
	in := &pb.GetAllCoursesRequest{}
	out := &pb.GetAllCoursesResponse{Courses: mockRes}

	type args struct {
		ctx   context.Context
		input *pb.GetAllCoursesRequest
	}

	type expect struct {
		svcOutput []*models.Course
		output    *pb.GetAllCoursesResponse
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

			got, err := client.GetAllCourses(tt.args.ctx, tt.args.input)
			if tt.exp.output != nil {
				assert.Equal(t, len(tt.exp.output.Courses), len(got.Courses))
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
	conn, client := setupClient(t)
	defer conn.Close()

	var (
		id     uint64 = 1
		in            = &pb.GetCourseRequest{Id: id}
		out           = &pb.Course{Id: id}
		svcOut        = models.CourseToORM(out)
	)

	type args struct {
		ctx   context.Context
		id    uint64
		input *pb.GetCourseRequest
	}

	type expect struct {
		svcOutput *models.Course
		output    *pb.Course
		err       error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id exists", args{ctx, id, in}, expect{svcOut, out, nil}},
		{"error", args{ctx, 10000, &pb.GetCourseRequest{Id: 10000}}, expect{nil, nil, errors.New("mock error")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("GetCourse", mockCtx, tt.args.id).Return(tt.exp.svcOutput, tt.exp.err)

			got, err := client.GetCourse(tt.args.ctx, tt.args.input)
			if tt.exp.output != nil && got != nil {
				assert.Equal(t, tt.exp.output.Id, got.Id)
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
	conn, client := setupClient(t)
	defer conn.Close()

	testPbTime := ptypes.TimestampNow()
	test := &pb.AcademicHistory{
		CandidateId:   1,
		InstitutionId: 1,
		CourseId:      1,
		CreatedAt:     testPbTime,
		UpdatedAt:     testPbTime,
	}
	in := &pb.CreateAcademicHistoryRequest{AcademicHistory: test}
	svcIn := models.AcademicHistoryToORM(test)
	out := test
	svcOut := models.AcademicHistoryToORM(test)

	type args struct {
		ctx      context.Context
		svcInput *models.AcademicHistory
		input    *pb.CreateAcademicHistoryRequest
	}

	type expect struct {
		svcOutput *models.AcademicHistory
		output    *pb.AcademicHistory
		err       error
	}

	for _, tt := range []struct {
		name string
		args args
		exp  expect
	}{
		{"error", args{ctx, &models.AcademicHistory{}, &pb.CreateAcademicHistoryRequest{AcademicHistory: &pb.AcademicHistory{}}}, expect{nil, nil, errors.New("Marshal called with nil")}},
		{"valid", args{ctx, svcIn, in}, expect{svcOut, out, nil}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("CreateAcademicHistory", mockCtx, tt.args.svcInput).Return(tt.exp.svcOutput, tt.exp.err)

			got, err := client.CreateAcademicHistory(ctx, tt.args.input)
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
	conn, client := setupClient(t)
	defer conn.Close()

	var (
		id     uint64 = 1
		in            = &pb.GetAcademicHistoryRequest{Id: id}
		out           = &pb.AcademicHistory{Id: id}
		svcOut        = models.AcademicHistoryToORM(out)
	)

	type args struct {
		ctx   context.Context
		id    uint64
		input *pb.GetAcademicHistoryRequest
	}

	type expect struct {
		svcOutput *models.AcademicHistory
		output    *pb.AcademicHistory
		err       error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id 1", args{ctx, id, in}, expect{svcOut, out, nil}},
		{"error", args{ctx, 10000, &pb.GetAcademicHistoryRequest{Id: 10000}}, expect{nil, nil, errors.New("mock error")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("GetAcademicHistory", mockCtx, tt.args.id).Return(tt.exp.svcOutput, tt.exp.err)

			got, err := client.GetAcademicHistory(tt.args.ctx, tt.args.input)
			if tt.exp.output != nil && got != nil {
				assert.Equal(t, tt.exp.output.Id, got.Id)
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
	conn, client := setupClient(t)
	defer conn.Close()

	updated := pb.AcademicHistory{
		Id:           1,
		YearObtained: 2020,
		UpdatedAt:    ptypes.TimestampNow(),
	}
	svcInOut := models.AcademicHistoryToORM(&updated)
	in := &pb.UpdateAcademicHistoryRequest{AcademicHistory: &updated}

	type args struct {
		ctx      context.Context
		svcInput *models.AcademicHistory
		input    *pb.UpdateAcademicHistoryRequest
	}

	type expect struct {
		svcOutput *models.AcademicHistory
		output    *pb.AcademicHistory
		err       error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id existing", args{ctx, svcInOut, in}, expect{svcInOut, &updated, nil}},
		{"id 10000", args{ctx, &models.AcademicHistory{ID: 10000}, &pb.UpdateAcademicHistoryRequest{AcademicHistory: &pb.AcademicHistory{Id: 10000}}}, expect{nil, nil, errors.New("Cannot update academic history with id")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("UpdateAcademicHistory", mockCtx, tt.args.svcInput).Return(tt.exp.svcOutput, tt.exp.err)

			got, err := client.UpdateAcademicHistory(tt.args.ctx, tt.args.input)
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
	conn, client := setupClient(t)
	defer conn.Close()

	type args struct {
		ctx   context.Context
		id    uint64
		input *pb.DeleteAcademicHistoryRequest
	}

	type expect struct {
		err error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id existing", args{ctx, 1, &pb.DeleteAcademicHistoryRequest{CandidateId: 1, Id: 1}}, expect{nil}},
		{"error", args{ctx, 10000, &pb.DeleteAcademicHistoryRequest{CandidateId: 10000, Id: 10000}}, expect{errors.New("mock error")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("DeleteAcademicHistory", mockCtx, tt.args.id).Return(tt.exp.err)

			_, err := client.DeleteAcademicHistory(tt.args.ctx, tt.args.input)
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
	conn, client := setupClient(t)
	defer conn.Close()

	test := &pb.Company{
		Name: "company",
	}
	in := &pb.CreateCompanyRequest{Company: test}
	svcInOut := models.CompanyToORM(test)
	out := test

	type args struct {
		ctx      context.Context
		svcInput *models.Company
		input    *pb.CreateCompanyRequest
	}

	type expect struct {
		svcOutput *models.Company
		output    *pb.Company
		err       error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"error", args{ctx, &models.Company{}, &pb.CreateCompanyRequest{Company: &pb.Company{}}}, expect{nil, nil, errors.New("Failed to insert company")}},
		{"valid", args{ctx, svcInOut, in}, expect{svcInOut, out, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("CreateCompany", mockCtx, tt.args.svcInput).Return(tt.exp.svcOutput, tt.exp.err)

			got, err := client.CreateCompany(tt.args.ctx, tt.args.input)
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
	conn, client := setupClient(t)
	defer conn.Close()

	mockRes := []*pb.Company{
		{},
		{},
	}
	svcRes := []*models.Company{
		{},
		{},
	}
	in := &pb.GetAllCompaniesRequest{}
	out := &pb.GetAllCompaniesResponse{Companies: mockRes}

	type args struct {
		ctx   context.Context
		input *pb.GetAllCompaniesRequest
	}

	type expect struct {
		svcOutput []*models.Company
		output    *pb.GetAllCompaniesResponse
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

			got, err := client.GetAllCompanies(tt.args.ctx, tt.args.input)
			if tt.exp.output != nil {
				assert.Equal(t, len(tt.exp.output.Companies), len(got.Companies))
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
	conn, client := setupClient(t)
	defer conn.Close()

	var (
		id     uint64 = 1
		in            = &pb.GetCompanyRequest{Id: id}
		out           = &pb.Company{Id: id}
		svcOut        = models.CompanyToORM(out)
	)

	type args struct {
		ctx   context.Context
		id    uint64
		input *pb.GetCompanyRequest
	}

	type expect struct {
		svcOutput *models.Company
		output    *pb.Company
		err       error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id exists", args{ctx, id, in}, expect{svcOut, out, nil}},
		{"error", args{ctx, 10000, &pb.GetCompanyRequest{Id: 10000}}, expect{nil, nil, errors.New("mock error")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("GetCompany", mockCtx, tt.args.id).Return(tt.exp.svcOutput, tt.exp.err)

			got, err := client.GetCompany(tt.args.ctx, tt.args.input)
			if tt.exp.output != nil && got != nil {
				assert.Equal(t, tt.exp.output.Id, got.Id)
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
	conn, client := setupClient(t)
	defer conn.Close()

	test := &pb.Department{
		Name: "department",
	}
	in := &pb.CreateDepartmentRequest{Department: test}
	svcInOut := models.DepartmentToORM(test)
	out := test

	type args struct {
		ctx      context.Context
		svcInput *models.Department
		input    *pb.CreateDepartmentRequest
	}

	type expect struct {
		svcOutput *models.Department
		output    *pb.Department
		err       error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"error", args{ctx, &models.Department{}, &pb.CreateDepartmentRequest{Department: &pb.Department{}}}, expect{nil, nil, errors.New("Failed to insert department")}},
		{"valid", args{ctx, svcInOut, in}, expect{svcInOut, out, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("CreateDepartment", mockCtx, tt.args.svcInput).Return(tt.exp.svcOutput, tt.exp.err)

			got, err := client.CreateDepartment(tt.args.ctx, tt.args.input)
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
	conn, client := setupClient(t)
	defer conn.Close()

	mockRes := []*pb.Department{
		{},
		{},
	}
	svcRes := []*models.Department{
		{},
		{},
	}
	in := &pb.GetAllDepartmentsRequest{}
	out := &pb.GetAllDepartmentsResponse{Departments: mockRes}

	type args struct {
		ctx   context.Context
		input *pb.GetAllDepartmentsRequest
	}

	type expect struct {
		svcOutput []*models.Department
		output    *pb.GetAllDepartmentsResponse
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

			got, err := client.GetAllDepartments(tt.args.ctx, tt.args.input)
			if tt.exp.output != nil {
				assert.Equal(t, len(tt.exp.output.Departments), len(got.Departments))
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
	conn, client := setupClient(t)
	defer conn.Close()

	var (
		id     uint64 = 1
		in            = &pb.GetDepartmentRequest{Id: id}
		out           = &pb.Department{Id: id}
		svcOut        = models.DepartmentToORM(out)
	)

	type args struct {
		ctx   context.Context
		id    uint64
		input *pb.GetDepartmentRequest
	}

	type expect struct {
		svcOutput *models.Department
		output    *pb.Department
		err       error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id exists", args{ctx, id, in}, expect{svcOut, out, nil}},
		{"error", args{ctx, 10000, &pb.GetDepartmentRequest{Id: 10000}}, expect{nil, nil, errors.New("mock error")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("GetDepartment", mockCtx, tt.args.id).Return(tt.exp.svcOutput, tt.exp.err)

			got, err := client.GetDepartment(tt.args.ctx, tt.args.input)
			if tt.exp.output != nil && got != nil {
				assert.Equal(t, tt.exp.output.Id, got.Id)
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
	conn, client := setupClient(t)
	defer conn.Close()

	testPbTime := ptypes.TimestampNow()
	test := &pb.JobHistory{
		CandidateId:  1,
		CompanyId:    1,
		DepartmentId: 1,
		Country:      "singapore",
		Title:        "software engineer",
		StartDate:    testPbTime,
		CreatedAt:    testPbTime,
		UpdatedAt:    testPbTime,
	}
	in := &pb.CreateJobHistoryRequest{JobHistory: test}
	svcIn := models.JobHistoryToORM(test)
	out := test
	svcOut := models.JobHistoryToORM(test)

	type args struct {
		ctx      context.Context
		svcInput *models.JobHistory
		input    *pb.CreateJobHistoryRequest
	}

	type expect struct {
		svcOutput *models.JobHistory
		output    *pb.JobHistory
		err       error
	}

	for _, tt := range []struct {
		name string
		args args
		exp  expect
	}{
		{"error", args{ctx, &models.JobHistory{}, &pb.CreateJobHistoryRequest{JobHistory: &pb.JobHistory{}}}, expect{nil, nil, errors.New("Marshal called with nil")}},
		{"valid", args{ctx, svcIn, in}, expect{svcOut, out, nil}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("CreateJobHistory", mockCtx, tt.args.svcInput).Return(tt.exp.svcOutput, tt.exp.err)

			got, err := client.CreateJobHistory(ctx, tt.args.input)
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
	conn, client := setupClient(t)
	defer conn.Close()

	var (
		id     uint64 = 1
		in            = &pb.GetJobHistoryRequest{Id: id}
		out           = &pb.JobHistory{Id: id}
		svcOut        = models.JobHistoryToORM(out)
	)

	type args struct {
		ctx   context.Context
		id    uint64
		input *pb.GetJobHistoryRequest
	}

	type expect struct {
		svcOutput *models.JobHistory
		output    *pb.JobHistory
		err       error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id 1", args{ctx, id, in}, expect{svcOut, out, nil}},
		{"error", args{ctx, 10000, &pb.GetJobHistoryRequest{Id: 10000}}, expect{nil, nil, errors.New("mock error")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("GetJobHistory", mockCtx, tt.args.id).Return(tt.exp.svcOutput, tt.exp.err)

			got, err := client.GetJobHistory(tt.args.ctx, tt.args.input)
			if tt.exp.output != nil && got != nil {
				assert.Equal(t, tt.exp.output.Id, got.Id)
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
	conn, client := setupClient(t)
	defer conn.Close()

	updated := pb.JobHistory{
		Id:        1,
		Country:   "indonesia",
		UpdatedAt: ptypes.TimestampNow(),
	}
	svcInOut := models.JobHistoryToORM(&updated)
	in := &pb.UpdateJobHistoryRequest{JobHistory: &updated}

	type args struct {
		ctx      context.Context
		svcInput *models.JobHistory
		input    *pb.UpdateJobHistoryRequest
	}

	type expect struct {
		svcOutput *models.JobHistory
		output    *pb.JobHistory
		err       error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id existing", args{ctx, svcInOut, in}, expect{svcInOut, &updated, nil}},
		{"id 10000", args{ctx, &models.JobHistory{ID: 10000}, &pb.UpdateJobHistoryRequest{JobHistory: &pb.JobHistory{Id: 10000}}}, expect{nil, nil, errors.New("Cannot update job history with id")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("UpdateJobHistory", mockCtx, tt.args.svcInput).Return(tt.exp.svcOutput, tt.exp.err)

			got, err := client.UpdateJobHistory(tt.args.ctx, tt.args.input)
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
	conn, client := setupClient(t)
	defer conn.Close()

	type args struct {
		ctx   context.Context
		id    uint64
		input *pb.DeleteJobHistoryRequest
	}

	type expect struct {
		err error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id existing", args{ctx, 1, &pb.DeleteJobHistoryRequest{CandidateId: 1, Id: 1}}, expect{nil}},
		{"error", args{ctx, 10000, &pb.DeleteJobHistoryRequest{CandidateId: 10000, Id: 10000}}, expect{errors.New("mock error")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc.On("DeleteJobHistory", mockCtx, tt.args.id).Return(tt.exp.err)

			_, err := client.DeleteJobHistory(tt.args.ctx, tt.args.input)
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
	svc.AssertExpectations(t)
}
