package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"in-backend/services/profile"
	"in-backend/services/profile/models"
)

// Endpoints holds all Go kit endpoints for the Profile Service.
type Endpoints struct {
	CreateCandidate  endpoint.Endpoint
	GetAllCandidates endpoint.Endpoint
	GetCandidateByID endpoint.Endpoint
	UpdateCandidate  endpoint.Endpoint
	DeleteCandidate  endpoint.Endpoint

	CreateSkill  endpoint.Endpoint
	GetSkill     endpoint.Endpoint
	GetAllSkills endpoint.Endpoint

	CreateUserSkill endpoint.Endpoint
	DeleteUserSkill endpoint.Endpoint

	CreateInstitution  endpoint.Endpoint
	GetInstitution     endpoint.Endpoint
	GetAllInstitutions endpoint.Endpoint

	CreateCourse  endpoint.Endpoint
	GetCourse     endpoint.Endpoint
	GetAllCourses endpoint.Endpoint

	CreateAcademicHistory endpoint.Endpoint
	GetAcademicHistory    endpoint.Endpoint
	UpdateAcademicHistory endpoint.Endpoint
	DeleteAcademicHistory endpoint.Endpoint

	CreateCompany   endpoint.Endpoint
	GetCompany      endpoint.Endpoint
	GetAllCompanies endpoint.Endpoint

	CreateDepartment  endpoint.Endpoint
	GetDepartment     endpoint.Endpoint
	GetAllDepartments endpoint.Endpoint

	CreateJobHistory endpoint.Endpoint
	GetJobHistory    endpoint.Endpoint
	UpdateJobHistory endpoint.Endpoint
	DeleteJobHistory endpoint.Endpoint
}

// MakeEndpoints initializes all Go kit endpoints for the Profile service.
func MakeEndpoints(s profile.Service) Endpoints {
	return Endpoints{
		CreateCandidate:  makeCreateCandidateEndpoint(s),
		GetAllCandidates: makeGetAllCandidatesEndpoint(s),
		GetCandidateByID: makeGetCandidateByIDEndpoint(s),
		UpdateCandidate:  makeUpdateCandidateEndpoint(s),
		DeleteCandidate:  makeDeleteCandidateEndpoint(s),

		CreateSkill:  makeCreateSkillEndpoint(s),
		GetSkill:     makeGetSkillEndpoint(s),
		GetAllSkills: makeGetAllSkillsEndpoint(s),

		CreateUserSkill: makeCreateUserSkillEndpoint(s),
		DeleteUserSkill: makeDeleteUserSkillEndpoint(s),

		CreateInstitution:  makeCreateInstitutionEndpoint(s),
		GetInstitution:     makeGetInstitutionEndpoint(s),
		GetAllInstitutions: makeGetAllInstitutionsEndpoint(s),

		CreateCourse:  makeCreateCourseEndpoint(s),
		GetCourse:     makeGetCourseEndpoint(s),
		GetAllCourses: makeGetAllCoursesEndpoint(s),

		CreateAcademicHistory: makeCreateAcademicHistoryEndpoint(s),
		GetAcademicHistory:    makeGetAcademicHistoryEndpoint(s),
		UpdateAcademicHistory: makeUpdateAcademicHistoryEndpoint(s),
		DeleteAcademicHistory: makeDeleteAcademicHistoryEndpoint(s),

		CreateCompany:   makeCreateCompanyEndpoint(s),
		GetCompany:      makeGetCompanyEndpoint(s),
		GetAllCompanies: makeGetAllCompaniesEndpoint(s),

		CreateDepartment:  makeCreateDepartmentEndpoint(s),
		GetDepartment:     makeGetDepartmentEndpoint(s),
		GetAllDepartments: makeGetAllDepartmentsEndpoint(s),

		CreateJobHistory: makeCreateJobHistoryEndpoint(s),
		GetJobHistory:    makeGetJobHistoryEndpoint(s),
		UpdateJobHistory: makeUpdateJobHistoryEndpoint(s),
		DeleteJobHistory: makeDeleteJobHistoryEndpoint(s),
	}
}

/* -------------- Candidate -------------- */

func makeCreateCandidateEndpoint(s profile.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateCandidateRequest)
		c, err := s.CreateCandidate(ctx, req.Candidate)
		return CreateCandidateResponse{Candidate: c, Err: err}, nil
	}
}

// CreateCandidateRequest declares the inputs required for creating a candidate
type CreateCandidateRequest struct {
	Candidate *models.Candidate
}

// CreateCandidateResponse declares the outputs after attempting to create a candidate
type CreateCandidateResponse struct {
	Candidate *models.Candidate
	Err       error
}

func makeGetAllCandidatesEndpoint(s profile.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAllCandidatesRequest)
		f := models.CandidateFilters(req)
		c, err := s.GetAllCandidates(ctx, f)
		return GetAllCandidatesResponse{Candidates: c, Err: err}, nil
	}
}

// GetAllCandidatesRequest declares the inputs required for getting all candidates
type GetAllCandidatesRequest struct {
	ID              []uint64
	FirstName       string
	LastName        string
	Email           string
	ContactNumber   string
	Gender          []string
	Nationality     []string
	ResidenceCity   []string
	MinSalary       uint32
	MaxSalary       uint32
	EducationLevel  []string
	MaxNoticePeriod uint32
}

// GetAllCandidatesResponse declares the outputs after attempting to get all candidates
type GetAllCandidatesResponse struct {
	Candidates []*models.Candidate
	Err        error
}

func makeGetCandidateByIDEndpoint(s profile.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetCandidateByIDRequest)
		c, err := s.GetCandidateByID(ctx, req.ID)
		return GetCandidateByIDResponse{Candidate: c, Err: err}, nil
	}
}

// GetCandidateByIDRequest declares the inputs required for getting a single candidate by ID
type GetCandidateByIDRequest struct {
	ID uint64
}

// GetCandidateByIDResponse declares the outputs after attempting to get a single candidate by ID
type GetCandidateByIDResponse struct {
	Candidate *models.Candidate
	Err       error
}

func makeUpdateCandidateEndpoint(s profile.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateCandidateRequest)
		c, err := s.UpdateCandidate(ctx, req.Candidate)
		return UpdateCandidateResponse{Candidate: c, Err: err}, nil
	}
}

// UpdateCandidateRequest declares the inputs required for updating a candidate
type UpdateCandidateRequest struct {
	ID        uint64
	Candidate *models.Candidate
}

// UpdateCandidateResponse declares the outputs after attempting to update a candidate
type UpdateCandidateResponse struct {
	Candidate *models.Candidate
	Err       error
}

func makeDeleteCandidateEndpoint(s profile.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteCandidateRequest)
		err := s.DeleteCandidate(ctx, req.ID)
		return DeleteCandidateResponse{Err: err}, nil
	}
}

// DeleteCandidateRequest declares the inputs required for deleting a candidate
type DeleteCandidateRequest struct {
	ID uint64
}

// DeleteCandidateResponse declares the outputs after attempting to delete a candidate
type DeleteCandidateResponse struct {
	Err error
}

/* -------------- Skill -------------- */

func makeCreateSkillEndpoint(s profile.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateSkillRequest)
		sk, err := s.CreateSkill(ctx, req.Skill)
		return CreateSkillResponse{Skill: sk, Err: err}, nil
	}
}

// CreateSkillRequest declares the inputs required for creating a skill
type CreateSkillRequest struct {
	Skill *models.Skill
}

// CreateSkillResponse declares the outputs after attempting to create a skill
type CreateSkillResponse struct {
	Skill *models.Skill
	Err   error
}

func makeGetSkillEndpoint(s profile.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetSkillRequest)
		sk, err := s.GetSkill(ctx, req.ID)
		return GetSkillResponse{Skill: sk, Err: err}, nil
	}
}

// GetSkillRequest declares the inputs required for getting a single skill by ID
type GetSkillRequest struct {
	ID uint64
}

// GetSkillResponse declares the outputs after attempting to get a single skill by ID
type GetSkillResponse struct {
	Skill *models.Skill
	Err   error
}

func makeGetAllSkillsEndpoint(s profile.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAllSkillsRequest)
		f := models.SkillFilters(req)
		sk, err := s.GetAllSkills(ctx, f)
		return GetAllSkillsResponse{Skills: sk, Err: err}, nil
	}
}

// GetAllSkillsRequest declares the inputs required for getting all skills
type GetAllSkillsRequest struct {
	Name []string
}

// GetAllSkillsResponse declares the outputs after attempting to get all skills
type GetAllSkillsResponse struct {
	Skills []*models.Skill
	Err    error
}

/* -------------- User Skill -------------- */

func makeCreateUserSkillEndpoint(s profile.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserSkillRequest)
		us, err := s.CreateUserSkill(ctx, req.UserSkill)
		return CreateUserSkillResponse{UserSkill: us, Err: err}, nil
	}
}

// CreateUserSkillRequest declares the inputs required for creating a UserSkill
type CreateUserSkillRequest struct {
	UserSkill *models.UserSkill
}

// CreateUserSkillResponse declares the outputs after attempting to create a UserSkill
type CreateUserSkillResponse struct {
	UserSkill *models.UserSkill
	Err       error
}

func makeDeleteUserSkillEndpoint(s profile.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteUserSkillRequest)
		err := s.DeleteUserSkill(ctx, req.ID)
		return DeleteUserSkillResponse{Err: err}, nil
	}
}

// DeleteUserSkillRequest declares the inputs required for deleting a UserSkill
type DeleteUserSkillRequest struct {
	ID uint64
}

// DeleteUserSkillResponse declares the outputs after attempting to delete a UserSkill
type DeleteUserSkillResponse struct {
	Err error
}

/* -------------- Institution -------------- */

func makeCreateInstitutionEndpoint(s profile.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateInstitutionRequest)
		i, err := s.CreateInstitution(ctx, req.Institution)
		return CreateInstitutionResponse{Institution: i, Err: err}, nil
	}
}

// CreateInstitutionRequest declares the inputs required for creating a Institution
type CreateInstitutionRequest struct {
	Institution *models.Institution
}

// CreateInstitutionResponse declares the outputs after attempting to create a Institution
type CreateInstitutionResponse struct {
	Institution *models.Institution
	Err         error
}

func makeGetInstitutionEndpoint(s profile.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetInstitutionRequest)
		i, err := s.GetInstitution(ctx, req.ID)
		return GetInstitutionResponse{Institution: i, Err: err}, nil
	}
}

// GetInstitutionRequest declares the inputs required for getting a single Institution by ID
type GetInstitutionRequest struct {
	ID uint64
}

// GetInstitutionResponse declares the outputs after attempting to get a single Institution by ID
type GetInstitutionResponse struct {
	Institution *models.Institution
	Err         error
}

func makeGetAllInstitutionsEndpoint(s profile.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAllInstitutionsRequest)
		f := models.InstitutionFilters(req)
		i, err := s.GetAllInstitutions(ctx, f)
		return GetAllInstitutionsResponse{Institutions: i, Err: err}, nil
	}
}

// GetAllInstitutionsRequest declares the inputs required for getting all Institutions
type GetAllInstitutionsRequest struct {
	Name    []string
	Country []string
}

// GetAllInstitutionsResponse declares the outputs after attempting to get all Institutions
type GetAllInstitutionsResponse struct {
	Institutions []*models.Institution
	Err          error
}

/* -------------- Course -------------- */

func makeCreateCourseEndpoint(s profile.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateCourseRequest)
		c, err := s.CreateCourse(ctx, req.Course)
		return CreateCourseResponse{Course: c, Err: err}, nil
	}
}

// CreateCourseRequest declares the inputs required for creating a Course
type CreateCourseRequest struct {
	Course *models.Course
}

// CreateCourseResponse declares the outputs after attempting to create a Course
type CreateCourseResponse struct {
	Course *models.Course
	Err    error
}

func makeGetCourseEndpoint(s profile.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetCourseRequest)
		c, err := s.GetCourse(ctx, req.ID)
		return GetCourseResponse{Course: c, Err: err}, nil
	}
}

// GetCourseRequest declares the inputs required for getting a single Course by ID
type GetCourseRequest struct {
	ID uint64
}

// GetCourseResponse declares the outputs after attempting to get a single Course by ID
type GetCourseResponse struct {
	Course *models.Course
	Err    error
}

func makeGetAllCoursesEndpoint(s profile.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAllCoursesRequest)
		f := models.CourseFilters(req)
		c, err := s.GetAllCourses(ctx, f)
		return GetAllCoursesResponse{Courses: c, Err: err}, nil
	}
}

// GetAllCoursesRequest declares the inputs required for getting all Courses
type GetAllCoursesRequest struct {
	Name  []string
	Level []string
}

// GetAllCoursesResponse declares the outputs after attempting to get all Courses
type GetAllCoursesResponse struct {
	Courses []*models.Course
	Err     error
}

/* -------------- Academic History -------------- */

func makeCreateAcademicHistoryEndpoint(s profile.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateAcademicHistoryRequest)
		a, err := s.CreateAcademicHistory(ctx, req.AcademicHistory)
		return CreateAcademicHistoryResponse{AcademicHistory: a, Err: err}, nil
	}
}

// CreateAcademicHistoryRequest declares the inputs required for creating a AcademicHistory
type CreateAcademicHistoryRequest struct {
	AcademicHistory *models.AcademicHistory
}

// CreateAcademicHistoryResponse declares the outputs after attempting to create a AcademicHistory
type CreateAcademicHistoryResponse struct {
	AcademicHistory *models.AcademicHistory
	Err             error
}

func makeGetAcademicHistoryEndpoint(s profile.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAcademicHistoryRequest)
		a, err := s.GetAcademicHistory(ctx, req.ID)
		return GetAcademicHistoryResponse{AcademicHistory: a, Err: err}, nil
	}
}

// GetAcademicHistoryRequest declares the inputs required for getting a single AcademicHistory by ID
type GetAcademicHistoryRequest struct {
	ID uint64
}

// GetAcademicHistoryResponse declares the outputs after attempting to get a single AcademicHistory by ID
type GetAcademicHistoryResponse struct {
	AcademicHistory *models.AcademicHistory
	Err             error
}

func makeUpdateAcademicHistoryEndpoint(s profile.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateAcademicHistoryRequest)
		a, err := s.UpdateAcademicHistory(ctx, req.AcademicHistory)
		return UpdateAcademicHistoryResponse{AcademicHistory: a, Err: err}, nil
	}
}

// UpdateAcademicHistoryRequest declares the inputs required for updating a AcademicHistory
type UpdateAcademicHistoryRequest struct {
	AcademicHistory *models.AcademicHistory
}

// UpdateAcademicHistoryResponse declares the outputs after attempting to update a AcademicHistory
type UpdateAcademicHistoryResponse struct {
	AcademicHistory *models.AcademicHistory
	Err             error
}

func makeDeleteAcademicHistoryEndpoint(s profile.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteAcademicHistoryRequest)
		err := s.DeleteAcademicHistory(ctx, req.ID)
		return DeleteAcademicHistoryResponse{Err: err}, nil
	}
}

// DeleteAcademicHistoryRequest declares the inputs required for deleting a AcademicHistory
type DeleteAcademicHistoryRequest struct {
	ID uint64
}

// DeleteAcademicHistoryResponse declares the outputs after attempting to delete a AcademicHistory
type DeleteAcademicHistoryResponse struct {
	Err error
}

/* -------------- Company -------------- */

func makeCreateCompanyEndpoint(s profile.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateCompanyRequest)
		c, err := s.CreateCompany(ctx, req.Company)
		return CreateCompanyResponse{Company: c, Err: err}, nil
	}
}

// CreateCompanyRequest declares the inputs required for creating a Company
type CreateCompanyRequest struct {
	Company *models.Company
}

// CreateCompanyResponse declares the outputs after attempting to create a Company
type CreateCompanyResponse struct {
	Company *models.Company
	Err     error
}

func makeGetCompanyEndpoint(s profile.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetCompanyRequest)
		c, err := s.GetCompany(ctx, req.ID)
		return GetCompanyResponse{Company: c, Err: err}, nil
	}
}

// GetCompanyRequest declares the inputs required for getting a single Company by ID
type GetCompanyRequest struct {
	ID uint64
}

// GetCompanyResponse declares the outputs after attempting to get a single Company by ID
type GetCompanyResponse struct {
	Company *models.Company
	Err     error
}

func makeGetAllCompaniesEndpoint(s profile.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAllCompaniesRequest)
		f := models.CompanyFilters(req)
		c, err := s.GetAllCompanies(ctx, f)
		return GetAllCompaniesResponse{Companies: c, Err: err}, nil
	}
}

// GetAllCompaniesRequest declares the inputs required for getting all Companies
type GetAllCompaniesRequest struct {
	Name []string
}

// GetAllCompaniesResponse declares the outputs after attempting to get all Companies
type GetAllCompaniesResponse struct {
	Companies []*models.Company
	Err       error
}

/* -------------- Department -------------- */

func makeCreateDepartmentEndpoint(s profile.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateDepartmentRequest)
		d, err := s.CreateDepartment(ctx, req.Department)
		return CreateDepartmentResponse{Department: d, Err: err}, nil
	}
}

// CreateDepartmentRequest declares the inputs required for creating a Department
type CreateDepartmentRequest struct {
	Department *models.Department
}

// CreateDepartmentResponse declares the outputs after attempting to create a Department
type CreateDepartmentResponse struct {
	Department *models.Department
	Err        error
}

func makeGetDepartmentEndpoint(s profile.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetDepartmentRequest)
		d, err := s.GetDepartment(ctx, req.ID)
		return GetDepartmentResponse{Department: d, Err: err}, nil
	}
}

// GetDepartmentRequest declares the inputs required for getting a single Department by ID
type GetDepartmentRequest struct {
	ID uint64
}

// GetDepartmentResponse declares the outputs after attempting to get a single Department by ID
type GetDepartmentResponse struct {
	Department *models.Department
	Err        error
}

func makeGetAllDepartmentsEndpoint(s profile.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAllDepartmentsRequest)
		f := models.DepartmentFilters(req)
		d, err := s.GetAllDepartments(ctx, f)
		return GetAllDepartmentsResponse{Departments: d, Err: err}, nil
	}
}

// GetAllDepartmentsRequest declares the inputs required for getting all Departments
type GetAllDepartmentsRequest struct {
	Name []string
}

// GetAllDepartmentsResponse declares the outputs after attempting to get all Departments
type GetAllDepartmentsResponse struct {
	Departments []*models.Department
	Err         error
}

/* -------------- Job History -------------- */

func makeCreateJobHistoryEndpoint(s profile.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateJobHistoryRequest)
		j, err := s.CreateJobHistory(ctx, req.JobHistory)
		return CreateJobHistoryResponse{JobHistory: j, Err: err}, nil
	}
}

// CreateJobHistoryRequest declares the inputs required for creating a JobHistory
type CreateJobHistoryRequest struct {
	JobHistory *models.JobHistory
}

// CreateJobHistoryResponse declares the outputs after attempting to create a JobHistory
type CreateJobHistoryResponse struct {
	JobHistory *models.JobHistory
	Err        error
}

func makeGetJobHistoryEndpoint(s profile.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetJobHistoryRequest)
		j, err := s.GetJobHistory(ctx, req.ID)
		return GetJobHistoryResponse{JobHistory: j, Err: err}, nil
	}
}

// GetJobHistoryRequest declares the inputs required for getting a single JobHistory by ID
type GetJobHistoryRequest struct {
	ID uint64
}

// GetJobHistoryResponse declares the outputs after attempting to get a single JobHistory by ID
type GetJobHistoryResponse struct {
	JobHistory *models.JobHistory
	Err        error
}

func makeUpdateJobHistoryEndpoint(s profile.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateJobHistoryRequest)
		j, err := s.UpdateJobHistory(ctx, req.JobHistory)
		return UpdateJobHistoryResponse{JobHistory: j, Err: err}, nil
	}
}

// UpdateJobHistoryRequest declares the inputs required for updating a JobHistory
type UpdateJobHistoryRequest struct {
	JobHistory *models.JobHistory
}

// UpdateJobHistoryResponse declares the outputs after attempting to update a JobHistory
type UpdateJobHistoryResponse struct {
	JobHistory *models.JobHistory
	Err        error
}

func makeDeleteJobHistoryEndpoint(s profile.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteJobHistoryRequest)
		err := s.DeleteJobHistory(ctx, req.ID)
		return DeleteJobHistoryResponse{Err: err}, nil
	}
}

// DeleteJobHistoryRequest declares the inputs required for deleting a JobHistory
type DeleteJobHistoryRequest struct {
	ID uint64
}

// DeleteJobHistoryResponse declares the outputs after attempting to delete a JobHistory
type DeleteJobHistoryResponse struct {
	Err error
}
