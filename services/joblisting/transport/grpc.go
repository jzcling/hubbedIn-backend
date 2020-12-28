package transport

import (
	"context"
	"in-backend/helpers"
	"in-backend/services/joblisting/endpoints"
	"in-backend/services/joblisting/models"
	"in-backend/services/joblisting/pb"

	"github.com/go-kit/kit/log"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// grpc transport service for joblisting Service.
type grpcServer struct {
	createJobPost     kitgrpc.Handler
	bulkCreateJobPost kitgrpc.Handler
	getAllJobPosts    kitgrpc.Handler
	getJobPostByID    kitgrpc.Handler
	updateJobPost     kitgrpc.Handler
	deleteJobPost     kitgrpc.Handler

	createCompany      kitgrpc.Handler
	localCreateCompany kitgrpc.Handler
	getAllCompanies    kitgrpc.Handler
	updateCompany      kitgrpc.Handler
	localUpdateCompany kitgrpc.Handler
	deleteCompany      kitgrpc.Handler

	createIndustry   kitgrpc.Handler
	getAllIndustries kitgrpc.Handler
	deleteIndustry   kitgrpc.Handler

	createJobFunction  kitgrpc.Handler
	getAllJobFunctions kitgrpc.Handler
	deleteJobFunction  kitgrpc.Handler

	createKeyPerson     kitgrpc.Handler
	bulkCreateKeyPerson kitgrpc.Handler
	getAllKeyPersons    kitgrpc.Handler
	getKeyPersonByID    kitgrpc.Handler
	updateKeyPerson     kitgrpc.Handler
	deleteKeyPerson     kitgrpc.Handler

	createJobPlatform  kitgrpc.Handler
	getAllJobPlatforms kitgrpc.Handler
	deleteJobPlatform  kitgrpc.Handler

	logger log.Logger
}

// NewGRPCServer returns a new gRPC service for the provided Go kit endpoints
func NewGRPCServer(
	endpoints endpoints.Endpoints,
	options []kitgrpc.ServerOption,
	logger log.Logger,
) pb.JoblistingServiceServer {
	errorLogger := kitgrpc.ServerErrorLogger(logger)
	options = append(options, errorLogger)

	return &grpcServer{
		createJobPost: kitgrpc.NewServer(
			endpoints.CreateJobPost,
			decodeCreateJobPostRequest,
			encodeCreateJobPostResponse,
			options...,
		),
		bulkCreateJobPost: kitgrpc.NewServer(
			endpoints.BulkCreateJobPost,
			decodeBulkCreateJobPostRequest,
			encodeBulkCreateJobPostResponse,
			options...,
		),
		getAllJobPosts: kitgrpc.NewServer(
			endpoints.GetAllJobPosts,
			decodeGetAllJobPostsRequest,
			encodeGetAllJobPostsResponse,
			options...,
		),
		getJobPostByID: kitgrpc.NewServer(
			endpoints.GetJobPostByID,
			decodeGetJobPostByIDRequest,
			encodeGetJobPostByIDResponse,
			options...,
		),
		updateJobPost: kitgrpc.NewServer(
			endpoints.UpdateJobPost,
			decodeUpdateJobPostRequest,
			encodeUpdateJobPostResponse,
			options...,
		),
		deleteJobPost: kitgrpc.NewServer(
			endpoints.DeleteJobPost,
			decodeDeleteJobPostRequest,
			encodeDeleteJobPostResponse,
			options...,
		),

		createCompany: kitgrpc.NewServer(
			endpoints.CreateCompany,
			decodeCreateCompanyRequest,
			encodeCreateCompanyResponse,
			options...,
		),
		localCreateCompany: kitgrpc.NewServer(
			endpoints.LocalCreateCompany,
			decodeCreateCompanyRequest,
			encodeCreateCompanyResponse,
			options...,
		),
		getAllCompanies: kitgrpc.NewServer(
			endpoints.GetAllCompanies,
			decodeGetAllCompaniesRequest,
			encodeGetAllCompaniesResponse,
			options...,
		),
		updateCompany: kitgrpc.NewServer(
			endpoints.UpdateCompany,
			decodeUpdateCompanyRequest,
			encodeUpdateCompanyResponse,
			options...,
		),
		localUpdateCompany: kitgrpc.NewServer(
			endpoints.LocalUpdateCompany,
			decodeUpdateCompanyRequest,
			encodeUpdateCompanyResponse,
			options...,
		),
		deleteCompany: kitgrpc.NewServer(
			endpoints.DeleteCompany,
			decodeDeleteCompanyRequest,
			encodeDeleteCompanyResponse,
			options...,
		),

		createIndustry: kitgrpc.NewServer(
			endpoints.CreateIndustry,
			decodeCreateIndustryRequest,
			encodeCreateIndustryResponse,
			options...,
		),
		getAllIndustries: kitgrpc.NewServer(
			endpoints.GetAllIndustries,
			decodeGetAllIndustriesRequest,
			encodeGetAllIndustriesResponse,
			options...,
		),
		deleteIndustry: kitgrpc.NewServer(
			endpoints.DeleteIndustry,
			decodeDeleteIndustryRequest,
			encodeDeleteIndustryResponse,
			options...,
		),

		createJobFunction: kitgrpc.NewServer(
			endpoints.CreateJobFunction,
			decodeCreateJobFunctionRequest,
			encodeCreateJobFunctionResponse,
			options...,
		),
		getAllJobFunctions: kitgrpc.NewServer(
			endpoints.GetAllJobFunctions,
			decodeGetAllJobFunctionsRequest,
			encodeGetAllJobFunctionsResponse,
			options...,
		),
		deleteJobFunction: kitgrpc.NewServer(
			endpoints.DeleteJobFunction,
			decodeDeleteJobFunctionRequest,
			encodeDeleteJobFunctionResponse,
			options...,
		),

		createKeyPerson: kitgrpc.NewServer(
			endpoints.CreateKeyPerson,
			decodeCreateKeyPersonRequest,
			encodeCreateKeyPersonResponse,
			options...,
		),
		bulkCreateKeyPerson: kitgrpc.NewServer(
			endpoints.BulkCreateKeyPerson,
			decodeBulkCreateKeyPersonRequest,
			encodeBulkCreateKeyPersonResponse,
			options...,
		),
		getAllKeyPersons: kitgrpc.NewServer(
			endpoints.GetAllKeyPersons,
			decodeGetAllKeyPersonsRequest,
			encodeGetAllKeyPersonsResponse,
			options...,
		),
		getKeyPersonByID: kitgrpc.NewServer(
			endpoints.GetKeyPersonByID,
			decodeGetKeyPersonByIDRequest,
			encodeGetKeyPersonByIDResponse,
			options...,
		),
		updateKeyPerson: kitgrpc.NewServer(
			endpoints.UpdateKeyPerson,
			decodeUpdateKeyPersonRequest,
			encodeUpdateKeyPersonResponse,
			options...,
		),
		deleteKeyPerson: kitgrpc.NewServer(
			endpoints.DeleteKeyPerson,
			decodeDeleteKeyPersonRequest,
			encodeDeleteKeyPersonResponse,
			options...,
		),

		createJobPlatform: kitgrpc.NewServer(
			endpoints.CreateJobPlatform,
			decodeCreateJobPlatformRequest,
			encodeCreateJobPlatformResponse,
			options...,
		),
		getAllJobPlatforms: kitgrpc.NewServer(
			endpoints.GetAllJobPlatforms,
			decodeGetAllJobPlatformsRequest,
			encodeGetAllJobPlatformsResponse,
			options...,
		),
		deleteJobPlatform: kitgrpc.NewServer(
			endpoints.DeleteJobPlatform,
			decodeDeleteJobPlatformRequest,
			encodeDeleteJobPlatformResponse,
			options...,
		),

		logger: logger,
	}
}

/* --------------- Job Post --------------- */

// CreateJobPost creates a new JobPost
func (s *grpcServer) CreateJobPost(ctx context.Context, req *pb.CreateJobPostRequest) (*pb.JobPost, error) {
	_, rep, err := s.createJobPost.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.JobPost), nil
}

// decodeCreateJobPostRequest decodes the incoming grpc payload to our go kit payload
func decodeCreateJobPostRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateJobPostRequest)
	return endpoints.CreateJobPostRequest{JobPost: models.JobPostToORM(req.JobPost)}, nil
}

// encodeCreateJobPostResponse encodes the outgoing go kit payload to the grpc payload
func encodeCreateJobPostResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.CreateJobPostResponse)
	err := getError(res.Err)
	if err == nil {
		return res.JobPost.ToProto(), nil
	}
	return nil, err
}

// BulkCreateJobPost creates multiple JobPosts
func (s *grpcServer) BulkCreateJobPost(ctx context.Context, req *pb.BulkCreateJobPostRequest) (*pb.BulkCreateJobPostResponse, error) {
	_, rep, err := s.bulkCreateJobPost.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.BulkCreateJobPostResponse), nil
}

// decodeBulkCreateJobPostRequest decodes the incoming grpc payload to our go kit payload
func decodeBulkCreateJobPostRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.BulkCreateJobPostRequest)
	var jobPosts []*models.JobPost
	for _, jobPost := range req.JobPosts {
		jobPosts = append(jobPosts, models.JobPostToORM(jobPost))
	}
	return endpoints.BulkCreateJobPostRequest{JobPosts: jobPosts}, nil
}

// encodeBulkCreateJobPostResponse encodes the outgoing go kit payload to the grpc payload
func encodeBulkCreateJobPostResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.BulkCreateJobPostResponse)
	err := getError(res.Err)
	if err == nil {
		var jobPosts []*pb.JobPost
		for _, jobPost := range res.JobPosts {
			jobPosts = append(jobPosts, jobPost.ToProto())
		}
		return &pb.BulkCreateJobPostResponse{JobPosts: jobPosts}, nil
	}
	return nil, err
}

// GetAllJobPosts returns all JobPosts
func (s *grpcServer) GetAllJobPosts(ctx context.Context, req *pb.GetAllJobPostsRequest) (*pb.GetAllJobPostsResponse, error) {
	_, rep, err := s.getAllJobPosts.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetAllJobPostsResponse), nil
}

// decodeGetAllJobPostsRequest decodes the incoming grpc payload to our go kit payload
func decodeGetAllJobPostsRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetAllJobPostsRequest)
	decoded := endpoints.GetAllJobPostsRequest{
		ID:                 req.Id,
		CompanyID:          req.CompanyId,
		HRContactID:        req.HrContactId,
		HiringManagerID:    req.HiringManagerId,
		JobPlatformID:      req.JobPlatformId,
		SkillID:            req.SkillId,
		Title:              req.Title,
		SeniorityLevel:     req.SeniorityLevel,
		MinYearsExperience: req.MinYearsExperience,
		MaxYearsExperience: req.MaxYearsExperience,
		EmploymentType:     req.EmploymentType,
		FunctionID:         req.FunctionId,
		IndustryID:         req.IndustryId,
		Remote:             req.Remote,
		Salary:             req.Salary,
		UpdatedAt:          helpers.ProtoTimeToTime(req.UpdatedAt),
		ExpireAt:           helpers.ProtoTimeToTime(req.ExpireAt),
	}
	return decoded, nil
}

// encodeGetAllJobPostsResponse encodes the outgoing go kit payload to the grpc payload
func encodeGetAllJobPostsResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.GetAllJobPostsResponse)
	err := getError(res.Err)
	if err == nil {
		var jobPosts []*pb.JobPost
		for _, jobPost := range res.JobPosts {
			jobPosts = append(jobPosts, jobPost.ToProto())
		}
		return &pb.GetAllJobPostsResponse{JobPosts: jobPosts}, nil
	}
	return nil, err
}

// GetJobPostByID returns a JobPost by ID
func (s *grpcServer) GetJobPostByID(ctx context.Context, req *pb.GetJobPostByIDRequest) (*pb.JobPost, error) {
	_, rep, err := s.getJobPostByID.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.JobPost), nil
}

// decodeGetJobPostByIDRequest decodes the incoming grpc payload to our go kit payload
func decodeGetJobPostByIDRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetJobPostByIDRequest)
	return endpoints.GetJobPostByIDRequest{ID: req.Id}, nil
}

// encodeGetJobPostByIDResponse encodes the outgoing go kit payload to the grpc payload
func encodeGetJobPostByIDResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.GetJobPostByIDResponse)
	err := getError(res.Err)
	if err == nil {
		return res.JobPost.ToProto(), nil
	}
	return nil, err
}

// UpdateJobPost updates a JobPost
func (s *grpcServer) UpdateJobPost(ctx context.Context, req *pb.UpdateJobPostRequest) (*pb.JobPost, error) {
	_, rep, err := s.updateJobPost.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.JobPost), nil
}

// decodeUpdateJobPostRequest decodes the incoming grpc payload to our go kit payload
func decodeUpdateJobPostRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateJobPostRequest)
	return endpoints.UpdateJobPostRequest{ID: req.Id, JobPost: models.JobPostToORM(req.JobPost)}, nil
}

// encodeUpdateJobPostResponse encodes the outgoing go kit payload to the grpc payload
func encodeUpdateJobPostResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.UpdateJobPostResponse)
	err := getError(res.Err)
	if err == nil {
		return res.JobPost.ToProto(), nil
	}
	return nil, err
}

// DeleteJobPost deletes a JobPost by ID
func (s *grpcServer) DeleteJobPost(ctx context.Context, req *pb.DeleteJobPostRequest) (*pb.DeleteJobPostResponse, error) {
	_, rep, err := s.deleteJobPost.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteJobPostResponse), nil
}

// decodeDeleteJobPostRequest decodes the incoming grpc payload to our go kit payload
func decodeDeleteJobPostRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.DeleteJobPostRequest)
	return endpoints.DeleteJobPostRequest{ID: req.Id}, nil
}

// encodeDeleteJobPostResponse encodes the outgoing go kit payload to the grpc payload
func encodeDeleteJobPostResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.DeleteJobPostResponse)
	err := getError(res.Err)
	if err == nil {
		return &pb.DeleteJobPostResponse{}, nil
	}
	return nil, err
}

/* --------------- Company --------------- */

// CreateCompany creates a new Company
func (s *grpcServer) CreateCompany(ctx context.Context, req *pb.CreateJobCompanyRequest) (*pb.JobCompany, error) {
	_, rep, err := s.createCompany.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.JobCompany), nil
}

// LocalCreateCompany creates a new Company
func (s *grpcServer) LocalCreateCompany(ctx context.Context, req *pb.CreateJobCompanyRequest) (*pb.JobCompany, error) {
	_, rep, err := s.localCreateCompany.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.JobCompany), nil
}

// decodeCreateCompanyRequest decodes the incoming grpc payload to our go kit payload
func decodeCreateCompanyRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateJobCompanyRequest)
	return endpoints.CreateCompanyRequest{Company: models.JobCompanyToORM(req.Company)}, nil
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

// GetAllCompanies returns all Companies
func (s *grpcServer) GetAllCompanies(ctx context.Context, req *pb.GetAllJobCompaniesRequest) (*pb.GetAllJobCompaniesResponse, error) {
	_, rep, err := s.getAllCompanies.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetAllJobCompaniesResponse), nil
}

// decodeGetAllCompaniesRequest decodes the incoming grpc payload to our go kit payload
func decodeGetAllCompaniesRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetAllJobCompaniesRequest)
	decoded := endpoints.GetAllCompaniesRequest{
		ID:   req.Id,
		Name: req.Name,
	}
	return decoded, nil
}

// encodeGetAllCompaniesResponse encodes the outgoing go kit payload to the grpc payload
func encodeGetAllCompaniesResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.GetAllCompaniesResponse)
	err := getError(res.Err)
	if err == nil {
		var companies []*pb.JobCompany
		for _, company := range res.Companies {
			companies = append(companies, company.ToProto())
		}
		return &pb.GetAllJobCompaniesResponse{Companies: companies}, nil
	}
	return nil, err
}

// UpdateCompany updates a Company
func (s *grpcServer) UpdateCompany(ctx context.Context, req *pb.UpdateJobCompanyRequest) (*pb.JobCompany, error) {
	_, rep, err := s.updateCompany.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.JobCompany), nil
}

// LocalUpdateCompany updates a Company
func (s *grpcServer) LocalUpdateCompany(ctx context.Context, req *pb.UpdateJobCompanyRequest) (*pb.JobCompany, error) {
	_, rep, err := s.localUpdateCompany.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.JobCompany), nil
}

// decodeUpdateCompanyRequest decodes the incoming grpc payload to our go kit payload
func decodeUpdateCompanyRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateJobCompanyRequest)
	return endpoints.UpdateCompanyRequest{ID: req.Id, Company: models.JobCompanyToORM(req.Company)}, nil
}

// encodeUpdateCompanyResponse encodes the outgoing go kit payload to the grpc payload
func encodeUpdateCompanyResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.UpdateCompanyResponse)
	err := getError(res.Err)
	if err == nil {
		return res.Company.ToProto(), nil
	}
	return nil, err
}

// DeleteCompany deletes a Company by ID
func (s *grpcServer) DeleteCompany(ctx context.Context, req *pb.DeleteJobCompanyRequest) (*pb.DeleteJobCompanyResponse, error) {
	_, rep, err := s.deleteCompany.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteJobCompanyResponse), nil
}

// decodeDeleteCompanyRequest decodes the incoming grpc payload to our go kit payload
func decodeDeleteCompanyRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.DeleteJobCompanyRequest)
	return endpoints.DeleteCompanyRequest{ID: req.Id}, nil
}

// encodeDeleteCompanyResponse encodes the outgoing go kit payload to the grpc payload
func encodeDeleteCompanyResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.DeleteCompanyResponse)
	err := getError(res.Err)
	if err == nil {
		return &pb.DeleteJobCompanyResponse{}, nil
	}
	return nil, err
}

/* --------------- Industry --------------- */

// CreateIndustry creates a new Industry
func (s *grpcServer) CreateIndustry(ctx context.Context, req *pb.CreateIndustryRequest) (*pb.Industry, error) {
	_, rep, err := s.createIndustry.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.Industry), nil
}

// decodeCreateIndustryRequest decodes the incoming grpc payload to our go kit payload
func decodeCreateIndustryRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateIndustryRequest)
	return endpoints.CreateIndustryRequest{Industry: models.IndustryToORM(req.Industry)}, nil
}

// encodeCreateIndustryResponse encodes the outgoing go kit payload to the grpc payload
func encodeCreateIndustryResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.CreateIndustryResponse)
	err := getError(res.Err)
	if err == nil {
		return res.Industry.ToProto(), nil
	}
	return nil, err
}

// GetAllIndustries returns all Industries
func (s *grpcServer) GetAllIndustries(ctx context.Context, req *pb.GetAllIndustriesRequest) (*pb.GetAllIndustriesResponse, error) {
	_, rep, err := s.getAllIndustries.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetAllIndustriesResponse), nil
}

// decodeGetAllIndustriesRequest decodes the incoming grpc payload to our go kit payload
func decodeGetAllIndustriesRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetAllIndustriesRequest)
	decoded := endpoints.GetAllIndustriesRequest{
		ID:   req.Id,
		Name: req.Name,
	}
	return decoded, nil
}

// encodeGetAllIndustriesResponse encodes the outgoing go kit payload to the grpc payload
func encodeGetAllIndustriesResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.GetAllIndustriesResponse)
	err := getError(res.Err)
	if err == nil {
		var industries []*pb.Industry
		for _, industry := range res.Industries {
			industries = append(industries, industry.ToProto())
		}
		return &pb.GetAllIndustriesResponse{Industries: industries}, nil
	}
	return nil, err
}

// DeleteIndustry deletes a Industry by ID
func (s *grpcServer) DeleteIndustry(ctx context.Context, req *pb.DeleteIndustryRequest) (*pb.DeleteIndustryResponse, error) {
	_, rep, err := s.deleteIndustry.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteIndustryResponse), nil
}

// decodeDeleteIndustryRequest decodes the incoming grpc payload to our go kit payload
func decodeDeleteIndustryRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.DeleteIndustryRequest)
	return endpoints.DeleteIndustryRequest{ID: req.Id}, nil
}

// encodeDeleteIndustryResponse encodes the outgoing go kit payload to the grpc payload
func encodeDeleteIndustryResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.DeleteIndustryResponse)
	err := getError(res.Err)
	if err == nil {
		return &pb.DeleteIndustryResponse{}, nil
	}
	return nil, err
}

/* --------------- Job Function --------------- */

// CreateJobFunction creates a new JobFunction
func (s *grpcServer) CreateJobFunction(ctx context.Context, req *pb.CreateJobFunctionRequest) (*pb.JobFunction, error) {
	_, rep, err := s.createJobFunction.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.JobFunction), nil
}

// decodeCreateJobFunctionRequest decodes the incoming grpc payload to our go kit payload
func decodeCreateJobFunctionRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateJobFunctionRequest)
	return endpoints.CreateJobFunctionRequest{JobFunction: models.JobFunctionToORM(req.JobFunction)}, nil
}

// encodeCreateJobFunctionResponse encodes the outgoing go kit payload to the grpc payload
func encodeCreateJobFunctionResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.CreateJobFunctionResponse)
	err := getError(res.Err)
	if err == nil {
		return res.JobFunction.ToProto(), nil
	}
	return nil, err
}

// GetAllJobFunctions returns all JobFunctions
func (s *grpcServer) GetAllJobFunctions(ctx context.Context, req *pb.GetAllJobFunctionsRequest) (*pb.GetAllJobFunctionsResponse, error) {
	_, rep, err := s.getAllJobFunctions.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetAllJobFunctionsResponse), nil
}

// decodeGetAllJobFunctionsRequest decodes the incoming grpc payload to our go kit payload
func decodeGetAllJobFunctionsRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetAllJobFunctionsRequest)
	decoded := endpoints.GetAllJobFunctionsRequest{
		ID:   req.Id,
		Name: req.Name,
	}
	return decoded, nil
}

// encodeGetAllJobFunctionsResponse encodes the outgoing go kit payload to the grpc payload
func encodeGetAllJobFunctionsResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.GetAllJobFunctionsResponse)
	err := getError(res.Err)
	if err == nil {
		var jobFunctions []*pb.JobFunction
		for _, jobFunction := range res.JobFunctions {
			jobFunctions = append(jobFunctions, jobFunction.ToProto())
		}
		return &pb.GetAllJobFunctionsResponse{JobFunctions: jobFunctions}, nil
	}
	return nil, err
}

// DeleteJobFunction deletes a JobFunction by ID
func (s *grpcServer) DeleteJobFunction(ctx context.Context, req *pb.DeleteJobFunctionRequest) (*pb.DeleteJobFunctionResponse, error) {
	_, rep, err := s.deleteJobFunction.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteJobFunctionResponse), nil
}

// decodeDeleteJobFunctionRequest decodes the incoming grpc payload to our go kit payload
func decodeDeleteJobFunctionRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.DeleteJobFunctionRequest)
	return endpoints.DeleteJobFunctionRequest{ID: req.Id}, nil
}

// encodeDeleteJobFunctionResponse encodes the outgoing go kit payload to the grpc payload
func encodeDeleteJobFunctionResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.DeleteJobFunctionResponse)
	err := getError(res.Err)
	if err == nil {
		return &pb.DeleteJobFunctionResponse{}, nil
	}
	return nil, err
}

/* --------------- Key Person --------------- */

// CreateKeyPerson creates a new KeyPerson
func (s *grpcServer) CreateKeyPerson(ctx context.Context, req *pb.CreateKeyPersonRequest) (*pb.KeyPerson, error) {
	_, rep, err := s.createKeyPerson.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.KeyPerson), nil
}

// decodeCreateKeyPersonRequest decodes the incoming grpc payload to our go kit payload
func decodeCreateKeyPersonRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateKeyPersonRequest)
	return endpoints.CreateKeyPersonRequest{KeyPerson: models.KeyPersonToORM(req.KeyPerson)}, nil
}

// encodeCreateKeyPersonResponse encodes the outgoing go kit payload to the grpc payload
func encodeCreateKeyPersonResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.CreateKeyPersonResponse)
	err := getError(res.Err)
	if err == nil {
		return res.KeyPerson.ToProto(), nil
	}
	return nil, err
}

// BulkCreateKeyPerson creates multiple KeyPersons
func (s *grpcServer) BulkCreateKeyPerson(ctx context.Context, req *pb.BulkCreateKeyPersonRequest) (*pb.BulkCreateKeyPersonResponse, error) {
	_, rep, err := s.bulkCreateKeyPerson.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.BulkCreateKeyPersonResponse), nil
}

// decodeBulkCreateKeyPersonRequest decodes the incoming grpc payload to our go kit payload
func decodeBulkCreateKeyPersonRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.BulkCreateKeyPersonRequest)
	var keyPersons []*models.KeyPerson
	for _, keyPerson := range req.KeyPersons {
		keyPersons = append(keyPersons, models.KeyPersonToORM(keyPerson))
	}
	return endpoints.BulkCreateKeyPersonRequest{KeyPersons: keyPersons}, nil
}

// encodeBulkCreateKeyPersonResponse encodes the outgoing go kit payload to the grpc payload
func encodeBulkCreateKeyPersonResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.BulkCreateKeyPersonResponse)
	err := getError(res.Err)
	if err == nil {
		var keyPersons []*pb.KeyPerson
		for _, keyPerson := range res.KeyPersons {
			keyPersons = append(keyPersons, keyPerson.ToProto())
		}
		return &pb.BulkCreateKeyPersonResponse{KeyPersons: keyPersons}, nil
	}
	return nil, err
}

// GetAllKeyPersons returns all KeyPersons
func (s *grpcServer) GetAllKeyPersons(ctx context.Context, req *pb.GetAllKeyPersonsRequest) (*pb.GetAllKeyPersonsResponse, error) {
	_, rep, err := s.getAllKeyPersons.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetAllKeyPersonsResponse), nil
}

// decodeGetAllKeyPersonsRequest decodes the incoming grpc payload to our go kit payload
func decodeGetAllKeyPersonsRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetAllKeyPersonsRequest)
	decoded := endpoints.GetAllKeyPersonsRequest{
		ID:            req.Id,
		CompanyID:     req.CompanyId,
		Name:          req.Name,
		ContactNumber: req.ContactNumber,
		Email:         req.Email,
		JobTitle:      req.JobTitle,
	}
	return decoded, nil
}

// encodeGetAllKeyPersonsResponse encodes the outgoing go kit payload to the grpc payload
func encodeGetAllKeyPersonsResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.GetAllKeyPersonsResponse)
	err := getError(res.Err)
	if err == nil {
		var keyPersons []*pb.KeyPerson
		for _, keyPerson := range res.KeyPersons {
			keyPersons = append(keyPersons, keyPerson.ToProto())
		}
		return &pb.GetAllKeyPersonsResponse{KeyPersons: keyPersons}, nil
	}
	return nil, err
}

// GetKeyPersonByID returns a KeyPerson by ID
func (s *grpcServer) GetKeyPersonByID(ctx context.Context, req *pb.GetKeyPersonByIDRequest) (*pb.KeyPerson, error) {
	_, rep, err := s.getKeyPersonByID.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.KeyPerson), nil
}

// decodeGetKeyPersonByIDRequest decodes the incoming grpc payload to our go kit payload
func decodeGetKeyPersonByIDRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetKeyPersonByIDRequest)
	return endpoints.GetKeyPersonByIDRequest{ID: req.Id}, nil
}

// encodeGetKeyPersonByIDResponse encodes the outgoing go kit payload to the grpc payload
func encodeGetKeyPersonByIDResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.GetKeyPersonByIDResponse)
	err := getError(res.Err)
	if err == nil {
		return res.KeyPerson.ToProto(), nil
	}
	return nil, err
}

// UpdateKeyPerson updates a KeyPerson
func (s *grpcServer) UpdateKeyPerson(ctx context.Context, req *pb.UpdateKeyPersonRequest) (*pb.KeyPerson, error) {
	_, rep, err := s.updateKeyPerson.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.KeyPerson), nil
}

// decodeUpdateKeyPersonRequest decodes the incoming grpc payload to our go kit payload
func decodeUpdateKeyPersonRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateKeyPersonRequest)
	return endpoints.UpdateKeyPersonRequest{ID: req.Id, KeyPerson: models.KeyPersonToORM(req.KeyPerson)}, nil
}

// encodeUpdateKeyPersonResponse encodes the outgoing go kit payload to the grpc payload
func encodeUpdateKeyPersonResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.UpdateKeyPersonResponse)
	err := getError(res.Err)
	if err == nil {
		return res.KeyPerson.ToProto(), nil
	}
	return nil, err
}

// DeleteKeyPerson deletes a KeyPerson by ID
func (s *grpcServer) DeleteKeyPerson(ctx context.Context, req *pb.DeleteKeyPersonRequest) (*pb.DeleteKeyPersonResponse, error) {
	_, rep, err := s.deleteKeyPerson.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteKeyPersonResponse), nil
}

// decodeDeleteKeyPersonRequest decodes the incoming grpc payload to our go kit payload
func decodeDeleteKeyPersonRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.DeleteKeyPersonRequest)
	return endpoints.DeleteKeyPersonRequest{ID: req.Id}, nil
}

// encodeDeleteKeyPersonResponse encodes the outgoing go kit payload to the grpc payload
func encodeDeleteKeyPersonResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.DeleteKeyPersonResponse)
	err := getError(res.Err)
	if err == nil {
		return &pb.DeleteKeyPersonResponse{}, nil
	}
	return nil, err
}

/* --------------- Job Platform --------------- */

// CreateJobPlatform creates a new JobPlatform
func (s *grpcServer) CreateJobPlatform(ctx context.Context, req *pb.CreateJobPlatformRequest) (*pb.JobPlatform, error) {
	_, rep, err := s.createJobPlatform.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.JobPlatform), nil
}

// decodeCreateJobPlatformRequest decodes the incoming grpc payload to our go kit payload
func decodeCreateJobPlatformRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateJobPlatformRequest)
	return endpoints.CreateJobPlatformRequest{JobPlatform: models.JobPlatformToORM(req.JobPlatform)}, nil
}

// encodeCreateJobPlatformResponse encodes the outgoing go kit payload to the grpc payload
func encodeCreateJobPlatformResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.CreateJobPlatformResponse)
	err := getError(res.Err)
	if err == nil {
		return res.JobPlatform.ToProto(), nil
	}
	return nil, err
}

// GetAllJobPlatforms returns all JobPlatforms
func (s *grpcServer) GetAllJobPlatforms(ctx context.Context, req *pb.GetAllJobPlatformsRequest) (*pb.GetAllJobPlatformsResponse, error) {
	_, rep, err := s.getAllJobPlatforms.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetAllJobPlatformsResponse), nil
}

// decodeGetAllJobPlatformsRequest decodes the incoming grpc payload to our go kit payload
func decodeGetAllJobPlatformsRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetAllJobPlatformsRequest)
	decoded := endpoints.GetAllJobPlatformsRequest{
		ID:   req.Id,
		Name: req.Name,
	}
	return decoded, nil
}

// encodeGetAllJobPlatformsResponse encodes the outgoing go kit payload to the grpc payload
func encodeGetAllJobPlatformsResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.GetAllJobPlatformsResponse)
	err := getError(res.Err)
	if err == nil {
		var jobPlatforms []*pb.JobPlatform
		for _, jobPlatform := range res.JobPlatforms {
			jobPlatforms = append(jobPlatforms, jobPlatform.ToProto())
		}
		return &pb.GetAllJobPlatformsResponse{JobPlatforms: jobPlatforms}, nil
	}
	return nil, err
}

// DeleteJobPlatform deletes a JobPlatform by ID
func (s *grpcServer) DeleteJobPlatform(ctx context.Context, req *pb.DeleteJobPlatformRequest) (*pb.DeleteJobPlatformResponse, error) {
	_, rep, err := s.deleteJobPlatform.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteJobPlatformResponse), nil
}

// decodeDeleteJobPlatformRequest decodes the incoming grpc payload to our go kit payload
func decodeDeleteJobPlatformRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.DeleteJobPlatformRequest)
	return endpoints.DeleteJobPlatformRequest{ID: req.Id}, nil
}

// encodeDeleteJobPlatformResponse encodes the outgoing go kit payload to the grpc payload
func encodeDeleteJobPlatformResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.DeleteJobPlatformResponse)
	err := getError(res.Err)
	if err == nil {
		return &pb.DeleteJobPlatformResponse{}, nil
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
