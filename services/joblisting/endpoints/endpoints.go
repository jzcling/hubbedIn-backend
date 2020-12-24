package endpoints

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"

	"in-backend/services/joblisting/interfaces"
	"in-backend/services/joblisting/models"
)

// Endpoints holds all Go kit endpoints for the joblisting Service.
type Endpoints struct {
	CreateJobPost     endpoint.Endpoint
	BulkCreateJobPost endpoint.Endpoint
	GetAllJobPosts    endpoint.Endpoint
	GetJobPostByID    endpoint.Endpoint
	UpdateJobPost     endpoint.Endpoint
	DeleteJobPost     endpoint.Endpoint

	CreateCompany   endpoint.Endpoint
	GetAllCompanies endpoint.Endpoint
	UpdateCompany   endpoint.Endpoint
	DeleteCompany   endpoint.Endpoint

	CreateIndustry   endpoint.Endpoint
	GetAllIndustries endpoint.Endpoint
	DeleteIndustry   endpoint.Endpoint

	CreateJobFunction  endpoint.Endpoint
	GetAllJobFunctions endpoint.Endpoint
	DeleteJobFunction  endpoint.Endpoint

	CreateKeyPerson     endpoint.Endpoint
	BulkCreateKeyPerson endpoint.Endpoint
	GetAllKeyPersons    endpoint.Endpoint
	GetKeyPersonByID    endpoint.Endpoint
	UpdateKeyPerson     endpoint.Endpoint
	DeleteKeyPerson     endpoint.Endpoint

	CreateJobPlatform  endpoint.Endpoint
	GetAllJobPlatforms endpoint.Endpoint
	DeleteJobPlatform  endpoint.Endpoint
}

// MakeEndpoints initializes all Go kit endpoints for the joblisting service.
func MakeEndpoints(s interfaces.Service) Endpoints {
	return Endpoints{
		CreateJobPost:     makeCreateJobPostEndpoint(s),
		BulkCreateJobPost: makeBulkCreateJobPostEndpoint(s),
		GetAllJobPosts:    makeGetAllJobPostsEndpoint(s),
		GetJobPostByID:    makeGetJobPostByIDEndpoint(s),
		UpdateJobPost:     makeUpdateJobPostEndpoint(s),
		DeleteJobPost:     makeDeleteJobPostEndpoint(s),

		CreateCompany:   makeCreateCompanyEndpoint(s),
		GetAllCompanies: makeGetAllCompaniesEndpoint(s),
		UpdateCompany:   makeUpdateCompanyEndpoint(s),
		DeleteCompany:   makeDeleteCompanyEndpoint(s),

		CreateIndustry:   makeCreateIndustryEndpoint(s),
		GetAllIndustries: makeGetAllIndustriesEndpoint(s),
		DeleteIndustry:   makeDeleteIndustryEndpoint(s),

		CreateJobFunction:  makeCreateJobFunctionEndpoint(s),
		GetAllJobFunctions: makeGetAllJobFunctionsEndpoint(s),
		DeleteJobFunction:  makeDeleteJobFunctionEndpoint(s),

		CreateKeyPerson:     makeCreateKeyPersonEndpoint(s),
		BulkCreateKeyPerson: makeBulkCreateKeyPersonEndpoint(s),
		GetAllKeyPersons:    makeGetAllKeyPersonsEndpoint(s),
		GetKeyPersonByID:    makeGetKeyPersonByIDEndpoint(s),
		UpdateKeyPerson:     makeUpdateKeyPersonEndpoint(s),
		DeleteKeyPerson:     makeDeleteKeyPersonEndpoint(s),

		CreateJobPlatform:  makeCreateJobPlatformEndpoint(s),
		GetAllJobPlatforms: makeGetAllJobPlatformsEndpoint(s),
		DeleteJobPlatform:  makeDeleteJobPlatformEndpoint(s),
	}
}

/* -------------- Job Post -------------- */

func makeCreateJobPostEndpoint(s interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateJobPostRequest)
		m, err := s.CreateJobPost(ctx, req.JobPost)
		return CreateJobPostResponse{JobPost: m, Err: err}, nil
	}
}

// CreateJobPostRequest declares the inputs required for creating a joblisting
type CreateJobPostRequest struct {
	JobPost *models.JobPost
}

// CreateJobPostResponse declares the outputs after attempting to create a joblisting
type CreateJobPostResponse struct {
	JobPost *models.JobPost
	Err     error
}

func makeBulkCreateJobPostEndpoint(s interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(BulkCreateJobPostRequest)
		m, err := s.BulkCreateJobPost(ctx, req.JobPosts)
		return BulkCreateJobPostResponse{JobPosts: m, Err: err}, nil
	}
}

// BulkCreateJobPostRequest declares the inputs required for creating a JobPost
type BulkCreateJobPostRequest struct {
	JobPosts []*models.JobPost
}

// BulkCreateJobPostResponse declares the outputs after attempting to create a JobPost
type BulkCreateJobPostResponse struct {
	JobPosts []*models.JobPost
	Err      error
}

func makeGetAllJobPostsEndpoint(s interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAllJobPostsRequest)
		f := models.JobPostFilters(req)
		m, err := s.GetAllJobPosts(ctx, f)
		return GetAllJobPostsResponse{JobPosts: m, Err: err}, nil
	}
}

// GetAllJobPostsRequest declares the inputs required for getting all joblistings
type GetAllJobPostsRequest struct {
	ID                 []uint64
	CompanyID          []uint64
	HRContactID        []uint64
	HiringManagerID    []uint64
	JobPlatformID      []uint64
	SkillID            []uint64
	Title              string
	SeniorityLevel     []string
	MinYearsExperience uint64
	MaxYearsExperience uint64
	EmploymentType     []string
	FunctionID         []uint64
	IndustryID         []uint64
	Remote             bool
	Salary             uint64
	UpdatedAt          *time.Time
	ExpireAt           *time.Time
}

// GetAllJobPostsResponse declares the outputs after attempting to get all joblistings
type GetAllJobPostsResponse struct {
	JobPosts []*models.JobPost
	Err      error
}

func makeGetJobPostByIDEndpoint(s interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetJobPostByIDRequest)
		m, err := s.GetJobPostByID(ctx, req.ID)
		return GetJobPostByIDResponse{JobPost: m, Err: err}, nil
	}
}

// GetJobPostByIDRequest declares the inputs required for getting a single joblisting by ID
type GetJobPostByIDRequest struct {
	ID uint64
}

// GetJobPostByIDResponse declares the outputs after attempting to get a single joblisting by ID
type GetJobPostByIDResponse struct {
	JobPost *models.JobPost
	Err     error
}

func makeUpdateJobPostEndpoint(s interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateJobPostRequest)
		m, err := s.UpdateJobPost(ctx, req.JobPost)
		return UpdateJobPostResponse{JobPost: m, Err: err}, nil
	}
}

// UpdateJobPostRequest declares the inputs required for updating a joblisting
type UpdateJobPostRequest struct {
	ID      uint64
	JobPost *models.JobPost
}

// UpdateJobPostResponse declares the outputs after attempting to update a joblisting
type UpdateJobPostResponse struct {
	JobPost *models.JobPost
	Err     error
}

func makeDeleteJobPostEndpoint(s interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteJobPostRequest)
		err := s.DeleteJobPost(ctx, req.ID)
		return DeleteJobPostResponse{Err: err}, nil
	}
}

// DeleteJobPostRequest declares the inputs required for deleting a joblisting
type DeleteJobPostRequest struct {
	ID uint64
}

// DeleteJobPostResponse declares the outputs after attempting to delete a joblisting
type DeleteJobPostResponse struct {
	Err error
}

/* -------------- Company -------------- */

func makeCreateCompanyEndpoint(s interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateCompanyRequest)
		m, err := s.CreateCompany(ctx, req.Company)
		return CreateCompanyResponse{Company: m, Err: err}, nil
	}
}

// CreateCompanyRequest declares the inputs required for creating a joblisting
type CreateCompanyRequest struct {
	Company *models.Company
}

// CreateCompanyResponse declares the outputs after attempting to create a joblisting
type CreateCompanyResponse struct {
	Company *models.Company
	Err     error
}

func makeGetAllCompaniesEndpoint(s interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAllCompaniesRequest)
		f := models.CompanyFilters(req)
		m, err := s.GetAllCompanies(ctx, f)
		return GetAllCompaniesResponse{Companies: m, Err: err}, nil
	}
}

// GetAllCompaniesRequest declares the inputs required for getting all joblistings
type GetAllCompaniesRequest struct {
	ID   []uint64
	Name string
}

// GetAllCompaniesResponse declares the outputs after attempting to get all joblistings
type GetAllCompaniesResponse struct {
	Companies []*models.Company
	Err       error
}

func makeUpdateCompanyEndpoint(s interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateCompanyRequest)
		m, err := s.UpdateCompany(ctx, req.Company)
		return UpdateCompanyResponse{Company: m, Err: err}, nil
	}
}

// UpdateCompanyRequest declares the inputs required for updating a joblisting
type UpdateCompanyRequest struct {
	ID      uint64
	Company *models.Company
}

// UpdateCompanyResponse declares the outputs after attempting to update a joblisting
type UpdateCompanyResponse struct {
	Company *models.Company
	Err     error
}

func makeDeleteCompanyEndpoint(s interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteCompanyRequest)
		err := s.DeleteCompany(ctx, req.ID)
		return DeleteCompanyResponse{Err: err}, nil
	}
}

// DeleteCompanyRequest declares the inputs required for deleting a joblisting
type DeleteCompanyRequest struct {
	ID uint64
}

// DeleteCompanyResponse declares the outputs after attempting to delete a joblisting
type DeleteCompanyResponse struct {
	Err error
}

/* -------------- Industry -------------- */

func makeCreateIndustryEndpoint(s interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateIndustryRequest)
		m, err := s.CreateIndustry(ctx, req.Industry)
		return CreateIndustryResponse{Industry: m, Err: err}, nil
	}
}

// CreateIndustryRequest declares the inputs required for creating a joblisting
type CreateIndustryRequest struct {
	Industry *models.Industry
}

// CreateIndustryResponse declares the outputs after attempting to create a joblisting
type CreateIndustryResponse struct {
	Industry *models.Industry
	Err      error
}

func makeGetAllIndustriesEndpoint(s interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAllIndustriesRequest)
		f := models.IndustryFilters(req)
		m, err := s.GetAllIndustries(ctx, f)
		return GetAllIndustriesResponse{Industries: m, Err: err}, nil
	}
}

// GetAllIndustriesRequest declares the inputs required for getting all joblistings
type GetAllIndustriesRequest struct {
	ID   []uint64
	Name string
}

// GetAllIndustriesResponse declares the outputs after attempting to get all joblistings
type GetAllIndustriesResponse struct {
	Industries []*models.Industry
	Err        error
}

func makeDeleteIndustryEndpoint(s interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteIndustryRequest)
		err := s.DeleteIndustry(ctx, req.ID)
		return DeleteIndustryResponse{Err: err}, nil
	}
}

// DeleteIndustryRequest declares the inputs required for deleting a joblisting
type DeleteIndustryRequest struct {
	ID uint64
}

// DeleteIndustryResponse declares the outputs after attempting to delete a joblisting
type DeleteIndustryResponse struct {
	Err error
}

/* -------------- Job Function -------------- */

func makeCreateJobFunctionEndpoint(s interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateJobFunctionRequest)
		m, err := s.CreateJobFunction(ctx, req.JobFunction)
		return CreateJobFunctionResponse{JobFunction: m, Err: err}, nil
	}
}

// CreateJobFunctionRequest declares the inputs required for creating a joblisting
type CreateJobFunctionRequest struct {
	JobFunction *models.JobFunction
}

// CreateJobFunctionResponse declares the outputs after attempting to create a joblisting
type CreateJobFunctionResponse struct {
	JobFunction *models.JobFunction
	Err         error
}

func makeGetAllJobFunctionsEndpoint(s interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAllJobFunctionsRequest)
		f := models.JobFunctionFilters(req)
		m, err := s.GetAllJobFunctions(ctx, f)
		return GetAllJobFunctionsResponse{JobFunctions: m, Err: err}, nil
	}
}

// GetAllJobFunctionsRequest declares the inputs required for getting all joblistings
type GetAllJobFunctionsRequest struct {
	ID   []uint64
	Name string
}

// GetAllJobFunctionsResponse declares the outputs after attempting to get all joblistings
type GetAllJobFunctionsResponse struct {
	JobFunctions []*models.JobFunction
	Err          error
}

func makeDeleteJobFunctionEndpoint(s interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteJobFunctionRequest)
		err := s.DeleteJobFunction(ctx, req.ID)
		return DeleteJobFunctionResponse{Err: err}, nil
	}
}

// DeleteJobFunctionRequest declares the inputs required for deleting a joblisting
type DeleteJobFunctionRequest struct {
	ID uint64
}

// DeleteJobFunctionResponse declares the outputs after attempting to delete a joblisting
type DeleteJobFunctionResponse struct {
	Err error
}

/* -------------- Key Person -------------- */

func makeCreateKeyPersonEndpoint(s interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateKeyPersonRequest)
		m, err := s.CreateKeyPerson(ctx, req.KeyPerson)
		return CreateKeyPersonResponse{KeyPerson: m, Err: err}, nil
	}
}

// CreateKeyPersonRequest declares the inputs required for creating a joblisting
type CreateKeyPersonRequest struct {
	KeyPerson *models.KeyPerson
}

// CreateKeyPersonResponse declares the outputs after attempting to create a joblisting
type CreateKeyPersonResponse struct {
	KeyPerson *models.KeyPerson
	Err       error
}

func makeBulkCreateKeyPersonEndpoint(s interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(BulkCreateKeyPersonRequest)
		m, err := s.BulkCreateKeyPerson(ctx, req.KeyPersons)
		return BulkCreateKeyPersonResponse{KeyPersons: m, Err: err}, nil
	}
}

// BulkCreateKeyPersonRequest declares the inputs required for creating a KeyPerson
type BulkCreateKeyPersonRequest struct {
	KeyPersons []*models.KeyPerson
}

// BulkCreateKeyPersonResponse declares the outputs after attempting to create a KeyPerson
type BulkCreateKeyPersonResponse struct {
	KeyPersons []*models.KeyPerson
	Err        error
}

func makeGetAllKeyPersonsEndpoint(s interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAllKeyPersonsRequest)
		f := models.KeyPersonFilters(req)
		m, err := s.GetAllKeyPersons(ctx, f)
		return GetAllKeyPersonsResponse{KeyPersons: m, Err: err}, nil
	}
}

// GetAllKeyPersonsRequest declares the inputs required for getting all joblistings
type GetAllKeyPersonsRequest struct {
	ID            []uint64
	CompanyID     []uint64
	Name          string
	ContactNumber string
	Email         string
	JobTitle      string
}

// GetAllKeyPersonsResponse declares the outputs after attempting to get all joblistings
type GetAllKeyPersonsResponse struct {
	KeyPersons []*models.KeyPerson
	Err        error
}

func makeGetKeyPersonByIDEndpoint(s interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetKeyPersonByIDRequest)
		m, err := s.GetKeyPersonByID(ctx, req.ID)
		return GetKeyPersonByIDResponse{KeyPerson: m, Err: err}, nil
	}
}

// GetKeyPersonByIDRequest declares the inputs required for getting a single joblisting by ID
type GetKeyPersonByIDRequest struct {
	ID uint64
}

// GetKeyPersonByIDResponse declares the outputs after attempting to get a single joblisting by ID
type GetKeyPersonByIDResponse struct {
	KeyPerson *models.KeyPerson
	Err       error
}

func makeUpdateKeyPersonEndpoint(s interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateKeyPersonRequest)
		m, err := s.UpdateKeyPerson(ctx, req.KeyPerson)
		return UpdateKeyPersonResponse{KeyPerson: m, Err: err}, nil
	}
}

// UpdateKeyPersonRequest declares the inputs required for updating a joblisting
type UpdateKeyPersonRequest struct {
	ID        uint64
	KeyPerson *models.KeyPerson
}

// UpdateKeyPersonResponse declares the outputs after attempting to update a joblisting
type UpdateKeyPersonResponse struct {
	KeyPerson *models.KeyPerson
	Err       error
}

func makeDeleteKeyPersonEndpoint(s interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteKeyPersonRequest)
		err := s.DeleteKeyPerson(ctx, req.ID)
		return DeleteKeyPersonResponse{Err: err}, nil
	}
}

// DeleteKeyPersonRequest declares the inputs required for deleting a joblisting
type DeleteKeyPersonRequest struct {
	ID uint64
}

// DeleteKeyPersonResponse declares the outputs after attempting to delete a joblisting
type DeleteKeyPersonResponse struct {
	Err error
}

/* -------------- Job Platform -------------- */

func makeCreateJobPlatformEndpoint(s interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateJobPlatformRequest)
		m, err := s.CreateJobPlatform(ctx, req.JobPlatform)
		return CreateJobPlatformResponse{JobPlatform: m, Err: err}, nil
	}
}

// CreateJobPlatformRequest declares the inputs required for creating a joblisting
type CreateJobPlatformRequest struct {
	JobPlatform *models.JobPlatform
}

// CreateJobPlatformResponse declares the outputs after attempting to create a joblisting
type CreateJobPlatformResponse struct {
	JobPlatform *models.JobPlatform
	Err         error
}

func makeGetAllJobPlatformsEndpoint(s interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAllJobPlatformsRequest)
		f := models.JobPlatformFilters(req)
		m, err := s.GetAllJobPlatforms(ctx, f)
		return GetAllJobPlatformsResponse{JobPlatforms: m, Err: err}, nil
	}
}

// GetAllJobPlatformsRequest declares the inputs required for getting all joblistings
type GetAllJobPlatformsRequest struct {
	ID   []uint64
	Name string
}

// GetAllJobPlatformsResponse declares the outputs after attempting to get all joblistings
type GetAllJobPlatformsResponse struct {
	JobPlatforms []*models.JobPlatform
	Err          error
}

func makeDeleteJobPlatformEndpoint(s interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteJobPlatformRequest)
		err := s.DeleteJobPlatform(ctx, req.ID)
		return DeleteJobPlatformResponse{Err: err}, nil
	}
}

// DeleteJobPlatformRequest declares the inputs required for deleting a joblisting
type DeleteJobPlatformRequest struct {
	ID uint64
}

// DeleteJobPlatformResponse declares the outputs after attempting to delete a joblisting
type DeleteJobPlatformResponse struct {
	Err error
}
