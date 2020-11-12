package transport

import (
	"context"
	"in-backend/services/profile/endpoints"
	"in-backend/services/profile/models"
	"in-backend/services/profile/pb"

	"github.com/go-kit/kit/log"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// grpc transport service for Profile Service.
type grpcServer struct {
	createCandidate  kitgrpc.Handler
	getAllCandidates kitgrpc.Handler
	getCandidateByID kitgrpc.Handler
	updateCandidate  kitgrpc.Handler
	deleteCandidate  kitgrpc.Handler

	createSkill  kitgrpc.Handler
	getSkill     kitgrpc.Handler
	getAllSkills kitgrpc.Handler

	createInstitution  kitgrpc.Handler
	getInstitution     kitgrpc.Handler
	getAllInstitutions kitgrpc.Handler

	createCourse  kitgrpc.Handler
	getCourse     kitgrpc.Handler
	getAllCourses kitgrpc.Handler

	createAcademicHistory kitgrpc.Handler
	getAcademicHistory    kitgrpc.Handler
	updateAcademicHistory kitgrpc.Handler
	deleteAcademicHistory kitgrpc.Handler

	createCompany   kitgrpc.Handler
	getCompany      kitgrpc.Handler
	getAllCompanies kitgrpc.Handler

	createDepartment  kitgrpc.Handler
	getDepartment     kitgrpc.Handler
	getAllDepartments kitgrpc.Handler

	createJobHistory kitgrpc.Handler
	getJobHistory    kitgrpc.Handler
	updateJobHistory kitgrpc.Handler
	deleteJobHistory kitgrpc.Handler

	logger log.Logger
}

// NewGRPCServer returns a new gRPC service for the provided Go kit endpoints
func NewGRPCServer(
	endpoints endpoints.Endpoints,
	options []kitgrpc.ServerOption,
	logger log.Logger,
) pb.ProfileServiceServer {
	errorLogger := kitgrpc.ServerErrorLogger(logger)
	options = append(options, errorLogger)

	return &grpcServer{
		createCandidate: kitgrpc.NewServer(
			endpoints.CreateCandidate,
			decodeCreateCandidateRequest,
			encodeCreateCandidateResponse,
			options...,
		),
		getAllCandidates: kitgrpc.NewServer(
			endpoints.GetAllCandidates,
			decodeGetAllCandidatesRequest,
			encodeGetAllCandidatesResponse,
			options...,
		),
		getCandidateByID: kitgrpc.NewServer(
			endpoints.GetCandidateByID,
			decodeGetCandidateByIDRequest,
			encodeGetCandidateByIDResponse,
			options...,
		),
		updateCandidate: kitgrpc.NewServer(
			endpoints.UpdateCandidate,
			decodeUpdateCandidateRequest,
			encodeUpdateCandidateResponse,
			options...,
		),
		deleteCandidate: kitgrpc.NewServer(
			endpoints.DeleteCandidate,
			decodeDeleteCandidateRequest,
			encodeDeleteCandidateResponse,
			options...,
		),

		createSkill: kitgrpc.NewServer(
			endpoints.CreateSkill,
			decodeCreateSkillRequest,
			encodeCreateSkillResponse,
			options...,
		),
		getSkill: kitgrpc.NewServer(
			endpoints.GetSkill,
			decodeGetSkillRequest,
			encodeGetSkillResponse,
			options...,
		),
		getAllSkills: kitgrpc.NewServer(
			endpoints.GetAllSkills,
			decodeGetAllSkillsRequest,
			encodeGetAllSkillsResponse,
			options...,
		),

		createInstitution: kitgrpc.NewServer(
			endpoints.CreateInstitution,
			decodeCreateInstitutionRequest,
			encodeCreateInstitutionResponse,
			options...,
		),
		getInstitution: kitgrpc.NewServer(
			endpoints.GetInstitution,
			decodeGetInstitutionRequest,
			encodeGetInstitutionResponse,
			options...,
		),
		getAllInstitutions: kitgrpc.NewServer(
			endpoints.GetAllInstitutions,
			decodeGetAllInstitutionsRequest,
			encodeGetAllInstitutionsResponse,
			options...,
		),

		createCourse: kitgrpc.NewServer(
			endpoints.CreateCourse,
			decodeCreateCourseRequest,
			encodeCreateCourseResponse,
			options...,
		),
		getCourse: kitgrpc.NewServer(
			endpoints.GetCourse,
			decodeGetCourseRequest,
			encodeGetCourseResponse,
			options...,
		),
		getAllCourses: kitgrpc.NewServer(
			endpoints.GetAllCourses,
			decodeGetAllCoursesRequest,
			encodeGetAllCoursesResponse,
			options...,
		),

		createAcademicHistory: kitgrpc.NewServer(
			endpoints.CreateAcademicHistory,
			decodeCreateAcademicHistoryRequest,
			encodeCreateAcademicHistoryResponse,
			options...,
		),
		getAcademicHistory: kitgrpc.NewServer(
			endpoints.GetAcademicHistory,
			decodeGetAcademicHistoryRequest,
			encodeGetAcademicHistoryResponse,
			options...,
		),
		updateAcademicHistory: kitgrpc.NewServer(
			endpoints.UpdateAcademicHistory,
			decodeUpdateAcademicHistoryRequest,
			encodeUpdateAcademicHistoryResponse,
			options...,
		),
		deleteAcademicHistory: kitgrpc.NewServer(
			endpoints.DeleteAcademicHistory,
			decodeDeleteAcademicHistoryRequest,
			encodeDeleteAcademicHistoryResponse,
			options...,
		),

		createCompany: kitgrpc.NewServer(
			endpoints.CreateCompany,
			decodeCreateCompanyRequest,
			encodeCreateCompanyResponse,
			options...,
		),
		getCompany: kitgrpc.NewServer(
			endpoints.GetCompany,
			decodeGetCompanyRequest,
			encodeGetCompanyResponse,
			options...,
		),
		getAllCompanies: kitgrpc.NewServer(
			endpoints.GetAllCompanies,
			decodeGetAllCompaniesRequest,
			encodeGetAllCompaniesResponse,
			options...,
		),

		createDepartment: kitgrpc.NewServer(
			endpoints.CreateDepartment,
			decodeCreateDepartmentRequest,
			encodeCreateDepartmentResponse,
			options...,
		),
		getDepartment: kitgrpc.NewServer(
			endpoints.GetDepartment,
			decodeGetDepartmentRequest,
			encodeGetDepartmentResponse,
			options...,
		),
		getAllDepartments: kitgrpc.NewServer(
			endpoints.GetAllDepartments,
			decodeGetAllDepartmentsRequest,
			encodeGetAllDepartmentsResponse,
			options...,
		),

		createJobHistory: kitgrpc.NewServer(
			endpoints.CreateJobHistory,
			decodeCreateJobHistoryRequest,
			encodeCreateJobHistoryResponse,
			options...,
		),
		getJobHistory: kitgrpc.NewServer(
			endpoints.GetJobHistory,
			decodeGetJobHistoryRequest,
			encodeGetJobHistoryResponse,
			options...,
		),
		updateJobHistory: kitgrpc.NewServer(
			endpoints.UpdateJobHistory,
			decodeUpdateJobHistoryRequest,
			encodeUpdateJobHistoryResponse,
			options...,
		),
		deleteJobHistory: kitgrpc.NewServer(
			endpoints.DeleteJobHistory,
			decodeDeleteJobHistoryRequest,
			encodeDeleteJobHistoryResponse,
			options...,
		),

		logger: logger,
	}
}

/* --------------- Candidate --------------- */

// CreateCandidate creates a new Candidate
func (s *grpcServer) CreateCandidate(ctx context.Context, req *pb.CreateCandidateRequest) (*pb.Candidate, error) {
	_, rep, err := s.createCandidate.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.Candidate), nil
}

// decodeCreateCandidateRequest decodes the incoming grpc payload to our go kit payload
func decodeCreateCandidateRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateCandidateRequest)
	return endpoints.CreateCandidateRequest{Candidate: models.CandidateToORM(req.Candidate)}, nil
}

// encodeCreateCandidateResponse encodes the outgoing go kit payload to the grpc payload
func encodeCreateCandidateResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.CreateCandidateResponse)
	err := getError(res.Err)
	if err == nil {
		return res.Candidate.ToProto(), nil
	}
	return nil, err
}

// GetAllCandidates returns all Candidates
func (s *grpcServer) GetAllCandidates(ctx context.Context, req *pb.GetAllCandidatesRequest) (*pb.GetAllCandidatesResponse, error) {
	_, rep, err := s.getAllCandidates.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetAllCandidatesResponse), nil
}

// decodeGetAllCandidatesRequest decodes the incoming grpc payload to our go kit payload
func decodeGetAllCandidatesRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetAllCandidatesRequest)
	return endpoints.GetAllCandidatesRequest{ID: req.Id}, nil
}

// encodeGetAllCandidatesResponse encodes the outgoing go kit payload to the grpc payload
func encodeGetAllCandidatesResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.GetAllCandidatesResponse)
	err := getError(res.Err)
	if err == nil {
		var candidates []*pb.Candidate
		for _, candidate := range res.Candidates {
			candidates = append(candidates, candidate.ToProto())
		}
		return &pb.GetAllCandidatesResponse{Candidates: candidates}, nil
	}
	return nil, err
}

// GetCandidateByID returns a Candidate by ID
func (s *grpcServer) GetCandidateByID(ctx context.Context, req *pb.GetCandidateByIDRequest) (*pb.Candidate, error) {
	_, rep, err := s.getCandidateByID.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.Candidate), nil
}

// decodeGetCandidateByIDRequest decodes the incoming grpc payload to our go kit payload
func decodeGetCandidateByIDRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetCandidateByIDRequest)
	return endpoints.GetCandidateByIDRequest{ID: req.Id}, nil
}

// encodeGetCandidateByIDResponse encodes the outgoing go kit payload to the grpc payload
func encodeGetCandidateByIDResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.GetCandidateByIDResponse)
	err := getError(res.Err)
	if err == nil {
		return res.Candidate.ToProto(), nil
	}
	return nil, err
}

// UpdateCandidate updates a Candidate
func (s *grpcServer) UpdateCandidate(ctx context.Context, req *pb.UpdateCandidateRequest) (*pb.Candidate, error) {
	_, rep, err := s.updateCandidate.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.Candidate), nil
}

// decodeUpdateCandidateRequest decodes the incoming grpc payload to our go kit payload
func decodeUpdateCandidateRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateCandidateRequest)
	return endpoints.UpdateCandidateRequest{Candidate: models.CandidateToORM(req.Candidate)}, nil
}

// encodeUpdateCandidateResponse encodes the outgoing go kit payload to the grpc payload
func encodeUpdateCandidateResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.UpdateCandidateResponse)
	err := getError(res.Err)
	if err == nil {
		return res.Candidate.ToProto(), nil
	}
	return nil, err
}

// DeleteCandidate deletes a Candidate by ID
func (s *grpcServer) DeleteCandidate(ctx context.Context, req *pb.DeleteCandidateRequest) (*pb.DeleteCandidateResponse, error) {
	_, rep, err := s.deleteCandidate.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteCandidateResponse), nil
}

// decodeDeleteCandidateRequest decodes the incoming grpc payload to our go kit payload
func decodeDeleteCandidateRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.DeleteCandidateRequest)
	return endpoints.DeleteCandidateRequest{ID: req.Id}, nil
}

// encodeDeleteCandidateResponse encodes the outgoing go kit payload to the grpc payload
func encodeDeleteCandidateResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.DeleteCandidateResponse)
	err := getError(res.Err)
	if err == nil {
		return &pb.DeleteCandidateResponse{}, nil
	}
	return nil, err
}

/* --------------- Skill --------------- */

// CreateSkill creates a new Skill
func (s *grpcServer) CreateSkill(ctx context.Context, req *pb.CreateSkillRequest) (*pb.Skill, error) {
	_, rep, err := s.createSkill.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.Skill), nil
}

// decodeCreateSkillRequest decodes the incoming grpc payload to our go kit payload
func decodeCreateSkillRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateSkillRequest)
	return endpoints.CreateSkillRequest{Skill: models.SkillToORM(req.Skill)}, nil
}

// encodeCreateSkillResponse encodes the outgoing go kit payload to the grpc payload
func encodeCreateSkillResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.CreateSkillResponse)
	err := getError(res.Err)
	if err == nil {
		return res.Skill.ToProto(), nil
	}
	return nil, err
}

// GetSkill returns a Skill by ID
func (s *grpcServer) GetSkill(ctx context.Context, req *pb.GetSkillRequest) (*pb.Skill, error) {
	_, rep, err := s.getSkill.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.Skill), nil
}

// decodeGetSkillRequest decodes the incoming grpc payload to our go kit payload
func decodeGetSkillRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetSkillRequest)
	return endpoints.GetSkillRequest{ID: req.Id}, nil
}

// encodeGetSkillResponse encodes the outgoing go kit payload to the grpc payload
func encodeGetSkillResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.GetSkillResponse)
	err := getError(res.Err)
	if err == nil {
		return res.Skill.ToProto(), nil
	}
	return nil, err
}

// GetAllSkills returns all Skills
func (s *grpcServer) GetAllSkills(ctx context.Context, req *pb.GetAllSkillsRequest) (*pb.GetAllSkillsResponse, error) {
	_, rep, err := s.getAllSkills.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetAllSkillsResponse), nil
}

// decodeGetAllSkillsRequest decodes the incoming grpc payload to our go kit payload
func decodeGetAllSkillsRequest(_ context.Context, request interface{}) (interface{}, error) {
	_ = request.(*pb.GetAllSkillsRequest)
	return endpoints.GetAllSkillsRequest{}, nil
}

// encodeGetAllSkillsResponse encodes the outgoing go kit payload to the grpc payload
func encodeGetAllSkillsResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.GetAllSkillsResponse)
	err := getError(res.Err)
	if err == nil {
		var skills []*pb.Skill
		for _, skill := range res.Skills {
			skills = append(skills, skill.ToProto())
		}
		return &pb.GetAllSkillsResponse{Skills: skills}, nil
	}
	return nil, err
}

/* --------------- Institution --------------- */

// CreateInstitution creates a new Institution
func (s *grpcServer) CreateInstitution(ctx context.Context, req *pb.CreateInstitutionRequest) (*pb.Institution, error) {
	_, rep, err := s.createInstitution.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.Institution), nil
}

// decodeCreateInstitutionRequest decodes the incoming grpc payload to our go kit payload
func decodeCreateInstitutionRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateInstitutionRequest)
	return endpoints.CreateInstitutionRequest{Institution: models.InstitutionToORM(req.Institution)}, nil
}

// encodeCreateInstitutionResponse encodes the outgoing go kit payload to the grpc payload
func encodeCreateInstitutionResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.CreateInstitutionResponse)
	err := getError(res.Err)
	if err == nil {
		return res.Institution.ToProto(), nil
	}
	return nil, err
}

// GetInstitution returns a Institution by ID
func (s *grpcServer) GetInstitution(ctx context.Context, req *pb.GetInstitutionRequest) (*pb.Institution, error) {
	_, rep, err := s.getInstitution.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.Institution), nil
}

// decodeGetInstitutionRequest decodes the incoming grpc payload to our go kit payload
func decodeGetInstitutionRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetInstitutionRequest)
	return endpoints.GetInstitutionRequest{ID: req.Id}, nil
}

// encodeGetInstitutionResponse encodes the outgoing go kit payload to the grpc payload
func encodeGetInstitutionResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.GetInstitutionResponse)
	err := getError(res.Err)
	if err == nil {
		return res.Institution.ToProto(), nil
	}
	return nil, err
}

// GetAllInstitutions returns all Institutions
func (s *grpcServer) GetAllInstitutions(ctx context.Context, req *pb.GetAllInstitutionsRequest) (*pb.GetAllInstitutionsResponse, error) {
	_, rep, err := s.getAllInstitutions.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetAllInstitutionsResponse), nil
}

// decodeGetAllInstitutionsRequest decodes the incoming grpc payload to our go kit payload
func decodeGetAllInstitutionsRequest(_ context.Context, request interface{}) (interface{}, error) {
	_ = request.(*pb.GetAllInstitutionsRequest)
	return endpoints.GetAllInstitutionsRequest{}, nil
}

// encodeGetAllInstitutionsResponse encodes the outgoing go kit payload to the grpc payload
func encodeGetAllInstitutionsResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.GetAllInstitutionsResponse)
	err := getError(res.Err)
	if err == nil {
		var institutions []*pb.Institution
		for _, institution := range res.Institutions {
			institutions = append(institutions, institution.ToProto())
		}
		return &pb.GetAllInstitutionsResponse{Institutions: institutions}, nil
	}
	return nil, err
}

/* --------------- Course --------------- */

// CreateCourse creates a new Course
func (s *grpcServer) CreateCourse(ctx context.Context, req *pb.CreateCourseRequest) (*pb.Course, error) {
	_, rep, err := s.createCourse.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.Course), nil
}

// decodeCreateCourseRequest decodes the incoming grpc payload to our go kit payload
func decodeCreateCourseRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateCourseRequest)
	return endpoints.CreateCourseRequest{Course: models.CourseToORM(req.Course)}, nil
}

// encodeCreateCourseResponse encodes the outgoing go kit payload to the grpc payload
func encodeCreateCourseResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.CreateCourseResponse)
	err := getError(res.Err)
	if err == nil {
		return res.Course.ToProto(), nil
	}
	return nil, err
}

// GetCourse returns a Course by ID
func (s *grpcServer) GetCourse(ctx context.Context, req *pb.GetCourseRequest) (*pb.Course, error) {
	_, rep, err := s.getCourse.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.Course), nil
}

// decodeGetCourseRequest decodes the incoming grpc payload to our go kit payload
func decodeGetCourseRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetCourseRequest)
	return endpoints.GetCourseRequest{ID: req.Id}, nil
}

// encodeGetCourseResponse encodes the outgoing go kit payload to the grpc payload
func encodeGetCourseResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.GetCourseResponse)
	err := getError(res.Err)
	if err == nil {
		return res.Course.ToProto(), nil
	}
	return nil, err
}

// GetAllCourses returns all Courses
func (s *grpcServer) GetAllCourses(ctx context.Context, req *pb.GetAllCoursesRequest) (*pb.GetAllCoursesResponse, error) {
	_, rep, err := s.getAllCourses.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetAllCoursesResponse), nil
}

// decodeGetAllCoursesRequest decodes the incoming grpc payload to our go kit payload
func decodeGetAllCoursesRequest(_ context.Context, request interface{}) (interface{}, error) {
	_ = request.(*pb.GetAllCoursesRequest)
	return endpoints.GetAllCoursesRequest{}, nil
}

// encodeGetAllCoursesResponse encodes the outgoing go kit payload to the grpc payload
func encodeGetAllCoursesResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.GetAllCoursesResponse)
	err := getError(res.Err)
	if err == nil {
		var courses []*pb.Course
		for _, course := range res.Courses {
			courses = append(courses, course.ToProto())
		}
		return &pb.GetAllCoursesResponse{Courses: courses}, nil
	}
	return nil, err
}

/* --------------- Academic History --------------- */

// CreateAcademicHistory creates a new AcademicHistory
func (s *grpcServer) CreateAcademicHistory(ctx context.Context, req *pb.CreateAcademicHistoryRequest) (*pb.AcademicHistory, error) {
	_, rep, err := s.createAcademicHistory.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AcademicHistory), nil
}

// decodeCreateAcademicHistoryRequest decodes the incoming grpc payload to our go kit payload
func decodeCreateAcademicHistoryRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateAcademicHistoryRequest)
	return endpoints.CreateAcademicHistoryRequest{AcademicHistory: models.AcademicHistoryToORM(req.AcademicHistory)}, nil
}

// encodeCreateAcademicHistoryResponse encodes the outgoing go kit payload to the grpc payload
func encodeCreateAcademicHistoryResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.CreateAcademicHistoryResponse)
	err := getError(res.Err)
	if err == nil {
		return res.AcademicHistory.ToProto(), nil
	}
	return nil, err
}

// GetAcademicHistory returns a AcademicHistory by ID
func (s *grpcServer) GetAcademicHistory(ctx context.Context, req *pb.GetAcademicHistoryRequest) (*pb.AcademicHistory, error) {
	_, rep, err := s.getAcademicHistory.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AcademicHistory), nil
}

// decodeGetAcademicHistoryRequest decodes the incoming grpc payload to our go kit payload
func decodeGetAcademicHistoryRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetAcademicHistoryRequest)
	return endpoints.GetAcademicHistoryRequest{ID: req.Id}, nil
}

// encodeGetAcademicHistoryResponse encodes the outgoing go kit payload to the grpc payload
func encodeGetAcademicHistoryResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.GetAcademicHistoryResponse)
	err := getError(res.Err)
	if err == nil {
		return res.AcademicHistory.ToProto(), nil
	}
	return nil, err
}

// UpdateAcademicHistory updates a AcademicHistory
func (s *grpcServer) UpdateAcademicHistory(ctx context.Context, req *pb.UpdateAcademicHistoryRequest) (*pb.AcademicHistory, error) {
	_, rep, err := s.updateAcademicHistory.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AcademicHistory), nil
}

// decodeUpdateAcademicHistoryRequest decodes the incoming grpc payload to our go kit payload
func decodeUpdateAcademicHistoryRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateAcademicHistoryRequest)
	return endpoints.UpdateAcademicHistoryRequest{AcademicHistory: models.AcademicHistoryToORM(req.AcademicHistory)}, nil
}

// encodeUpdateAcademicHistoryResponse encodes the outgoing go kit payload to the grpc payload
func encodeUpdateAcademicHistoryResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.UpdateAcademicHistoryResponse)
	err := getError(res.Err)
	if err == nil {
		return res.AcademicHistory.ToProto(), nil
	}
	return nil, err
}

// DeleteAcademicHistory deletes a AcademicHistory by ID
func (s *grpcServer) DeleteAcademicHistory(ctx context.Context, req *pb.DeleteAcademicHistoryRequest) (*pb.DeleteAcademicHistoryResponse, error) {
	_, rep, err := s.deleteAcademicHistory.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteAcademicHistoryResponse), nil
}

// decodeDeleteAcademicHistoryRequest decodes the incoming grpc payload to our go kit payload
func decodeDeleteAcademicHistoryRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.DeleteAcademicHistoryRequest)
	return endpoints.DeleteAcademicHistoryRequest{ID: req.Id}, nil
}

// encodeDeleteAcademicHistoryResponse encodes the outgoing go kit payload to the grpc payload
func encodeDeleteAcademicHistoryResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.DeleteAcademicHistoryResponse)
	err := getError(res.Err)
	if err == nil {
		return &pb.DeleteAcademicHistoryResponse{}, nil
	}
	return nil, err
}

/* --------------- Course --------------- */

// CreateCompany creates a new Company
func (s *grpcServer) CreateCompany(ctx context.Context, req *pb.CreateCompanyRequest) (*pb.Company, error) {
	_, rep, err := s.createCompany.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.Company), nil
}

// decodeCreateCompanyRequest decodes the incoming grpc payload to our go kit payload
func decodeCreateCompanyRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateCompanyRequest)
	return endpoints.CreateCompanyRequest{Company: models.CompanyToORM(req.Company)}, nil
}

// encodeCreateCompanyResponse encodes the outgoing go kit payload to the grpc payload
func encodeCreateCompanyResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.CreateCompanyResponse)
	err := getError(res.Err)
	if err == nil {
		return res.Company.ToProto(), nil
	}
	return nil, err
}

// GetCompany returns a Company by ID
func (s *grpcServer) GetCompany(ctx context.Context, req *pb.GetCompanyRequest) (*pb.Company, error) {
	_, rep, err := s.getCompany.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.Company), nil
}

// decodeGetCompanyRequest decodes the incoming grpc payload to our go kit payload
func decodeGetCompanyRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetCompanyRequest)
	return endpoints.GetCompanyRequest{ID: req.Id}, nil
}

// encodeGetCompanyResponse encodes the outgoing go kit payload to the grpc payload
func encodeGetCompanyResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.GetCompanyResponse)
	err := getError(res.Err)
	if err == nil {
		return res.Company.ToProto(), nil
	}
	return nil, err
}

// GetAllCompanys returns all Companys
func (s *grpcServer) GetAllCompanies(ctx context.Context, req *pb.GetAllCompaniesRequest) (*pb.GetAllCompaniesResponse, error) {
	_, rep, err := s.getAllCompanies.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetAllCompaniesResponse), nil
}

// decodeGetAllCompaniesRequest decodes the incoming grpc payload to our go kit payload
func decodeGetAllCompaniesRequest(_ context.Context, request interface{}) (interface{}, error) {
	_ = request.(*pb.GetAllCompaniesRequest)
	return endpoints.GetAllCompaniesRequest{}, nil
}

// encodeGetAllCompaniesResponse encodes the outgoing go kit payload to the grpc payload
func encodeGetAllCompaniesResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.GetAllCompaniesResponse)
	err := getError(res.Err)
	if err == nil {
		var companies []*pb.Company
		for _, company := range res.Companies {
			companies = append(companies, company.ToProto())
		}
		return &pb.GetAllCompaniesResponse{Companies: companies}, nil
	}
	return nil, err
}

/* --------------- Department --------------- */

// CreateDepartment creates a new Department
func (s *grpcServer) CreateDepartment(ctx context.Context, req *pb.CreateDepartmentRequest) (*pb.Department, error) {
	_, rep, err := s.createDepartment.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.Department), nil
}

// decodeCreateDepartmentRequest decodes the incoming grpc payload to our go kit payload
func decodeCreateDepartmentRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateDepartmentRequest)
	return endpoints.CreateDepartmentRequest{Department: models.DepartmentToORM(req.Department)}, nil
}

// encodeCreateDepartmentResponse encodes the outgoing go kit payload to the grpc payload
func encodeCreateDepartmentResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.CreateDepartmentResponse)
	err := getError(res.Err)
	if err == nil {
		return res.Department.ToProto(), nil
	}
	return nil, err
}

// GetDepartment returns a Department by ID
func (s *grpcServer) GetDepartment(ctx context.Context, req *pb.GetDepartmentRequest) (*pb.Department, error) {
	_, rep, err := s.getDepartment.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.Department), nil
}

// decodeGetDepartmentRequest decodes the incoming grpc payload to our go kit payload
func decodeGetDepartmentRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetDepartmentRequest)
	return endpoints.GetDepartmentRequest{ID: req.Id}, nil
}

// encodeGetDepartmentResponse encodes the outgoing go kit payload to the grpc payload
func encodeGetDepartmentResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.GetDepartmentResponse)
	err := getError(res.Err)
	if err == nil {
		return res.Department.ToProto(), nil
	}
	return nil, err
}

// GetAllDepartments returns all Departments
func (s *grpcServer) GetAllDepartments(ctx context.Context, req *pb.GetAllDepartmentsRequest) (*pb.GetAllDepartmentsResponse, error) {
	_, rep, err := s.getAllDepartments.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetAllDepartmentsResponse), nil
}

// decodeGetAllDepartmentsRequest decodes the incoming grpc payload to our go kit payload
func decodeGetAllDepartmentsRequest(_ context.Context, request interface{}) (interface{}, error) {
	_ = request.(*pb.GetAllDepartmentsRequest)
	return endpoints.GetAllDepartmentsRequest{}, nil
}

// encodeGetAllDepartmentsResponse encodes the outgoing go kit payload to the grpc payload
func encodeGetAllDepartmentsResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.GetAllDepartmentsResponse)
	err := getError(res.Err)
	if err == nil {
		var departments []*pb.Department
		for _, department := range res.Departments {
			departments = append(departments, department.ToProto())
		}
		return &pb.GetAllDepartmentsResponse{Departments: departments}, nil
	}
	return nil, err
}

/* --------------- Job History --------------- */

// CreateJobHistory creates a new JobHistory
func (s *grpcServer) CreateJobHistory(ctx context.Context, req *pb.CreateJobHistoryRequest) (*pb.JobHistory, error) {
	_, rep, err := s.createJobHistory.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.JobHistory), nil
}

// decodeCreateJobHistoryRequest decodes the incoming grpc payload to our go kit payload
func decodeCreateJobHistoryRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateJobHistoryRequest)
	return endpoints.CreateJobHistoryRequest{JobHistory: models.JobHistoryToORM(req.JobHistory)}, nil
}

// encodeCreateJobHistoryResponse encodes the outgoing go kit payload to the grpc payload
func encodeCreateJobHistoryResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.CreateJobHistoryResponse)
	err := getError(res.Err)
	if err == nil {
		return res.JobHistory.ToProto(), nil
	}
	return nil, err
}

// GetJobHistory returns a JobHistory by ID
func (s *grpcServer) GetJobHistory(ctx context.Context, req *pb.GetJobHistoryRequest) (*pb.JobHistory, error) {
	_, rep, err := s.getJobHistory.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.JobHistory), nil
}

// decodeGetJobHistoryRequest decodes the incoming grpc payload to our go kit payload
func decodeGetJobHistoryRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetJobHistoryRequest)
	return endpoints.GetJobHistoryRequest{ID: req.Id}, nil
}

// encodeGetJobHistoryResponse encodes the outgoing go kit payload to the grpc payload
func encodeGetJobHistoryResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.GetJobHistoryResponse)
	err := getError(res.Err)
	if err == nil {
		return res.JobHistory.ToProto(), nil
	}
	return nil, err
}

// UpdateJobHistory updates a JobHistory
func (s *grpcServer) UpdateJobHistory(ctx context.Context, req *pb.UpdateJobHistoryRequest) (*pb.JobHistory, error) {
	_, rep, err := s.updateJobHistory.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.JobHistory), nil
}

// decodeUpdateJobHistoryRequest decodes the incoming grpc payload to our go kit payload
func decodeUpdateJobHistoryRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateJobHistoryRequest)
	return endpoints.UpdateJobHistoryRequest{JobHistory: models.JobHistoryToORM(req.JobHistory)}, nil
}

// encodeUpdateJobHistoryResponse encodes the outgoing go kit payload to the grpc payload
func encodeUpdateJobHistoryResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.UpdateJobHistoryResponse)
	err := getError(res.Err)
	if err == nil {
		return res.JobHistory.ToProto(), nil
	}
	return nil, err
}

// DeleteJobHistory deletes a JobHistory by ID
func (s *grpcServer) DeleteJobHistory(ctx context.Context, req *pb.DeleteJobHistoryRequest) (*pb.DeleteJobHistoryResponse, error) {
	_, rep, err := s.deleteJobHistory.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteJobHistoryResponse), nil
}

// decodeDeleteJobHistoryRequest decodes the incoming grpc payload to our go kit payload
func decodeDeleteJobHistoryRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.DeleteJobHistoryRequest)
	return endpoints.DeleteJobHistoryRequest{ID: req.Id}, nil
}

// encodeDeleteJobHistoryResponse encodes the outgoing go kit payload to the grpc payload
func encodeDeleteJobHistoryResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.DeleteJobHistoryResponse)
	err := getError(res.Err)
	if err == nil {
		return &pb.DeleteJobHistoryResponse{}, nil
	}
	return nil, err
}

func getError(err error) error {
	switch err {
	case nil:
		return nil
	default:
		return status.Error(codes.Unknown, err.Error())
	}
}
