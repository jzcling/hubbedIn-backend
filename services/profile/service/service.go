package service

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"in-backend/services/profile"
	"in-backend/services/profile/models"
)

// Service implements the profile Service interface
type Service struct {
	repository profile.Repository
	logger     log.Logger
}

// New creates and returns a new Service that implements the profile Service interface
func New(r profile.Repository, l log.Logger) *Service {
	return &Service{
		repository: r,
		logger:     l,
	}
}

/* --------------- Candidate --------------- */

// CreateCandidate creates a new Candidate
func (s *Service) CreateCandidate(ctx context.Context, candidate *models.Candidate) (*models.Candidate, error) {
	logger := log.With(s.logger, "method", "CreateCandidate")

	c, err := s.repository.CreateCandidate(ctx, candidate)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return c, err
}

// GetAllCandidates returns all Candidates
func (s *Service) GetAllCandidates(ctx context.Context) ([]*models.Candidate, error) {
	logger := log.With(s.logger, "method", "GetAllCandidates")

	c, err := s.repository.GetAllCandidates(ctx)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return c, err
}

// GetCandidateByID returns a Candidate by ID
func (s *Service) GetCandidateByID(ctx context.Context, id uint64) (*models.Candidate, error) {
	logger := log.With(s.logger, "method", "GetCandidateByID")

	c, err := s.repository.GetCandidateByID(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return c, err
}

// UpdateCandidate updates a Candidate
func (s *Service) UpdateCandidate(ctx context.Context, candidate *models.Candidate) (*models.Candidate, error) {
	logger := log.With(s.logger, "method", "UpdateCandidate")

	c, err := s.repository.UpdateCandidate(ctx, candidate)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return c, err
}

// DeleteCandidate deletes a Candidate by ID
func (s *Service) DeleteCandidate(ctx context.Context, id uint64) error {
	logger := log.With(s.logger, "method", "DeleteCandidate")

	err := s.repository.DeleteCandidate(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return err
}

/* --------------- Skill --------------- */

// CreateSkill creates a new Skill
func (s *Service) CreateSkill(ctx context.Context, skill *models.Skill) (*models.Skill, error) {
	logger := log.With(s.logger, "method", "CreateSkill")

	sk, err := s.repository.CreateSkill(ctx, skill)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return sk, err
}

// GetSkill returns a Skill by ID
func (s *Service) GetSkill(ctx context.Context, id uint64) (*models.Skill, error) {
	logger := log.With(s.logger, "method", "GetSkill")

	sk, err := s.repository.GetSkill(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return sk, err
}

// GetAllSkills returns all Skills
func (s *Service) GetAllSkills(ctx context.Context) ([]*models.Skill, error) {
	logger := log.With(s.logger, "method", "GetAllSkills")

	sk, err := s.repository.GetAllSkills(ctx)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return sk, err
}

/* --------------- Institution --------------- */

// CreateInstitution creates a new Institution
func (s *Service) CreateInstitution(ctx context.Context, institution *models.Institution) (*models.Institution, error) {
	logger := log.With(s.logger, "method", "CreateInstitution")

	i, err := s.repository.CreateInstitution(ctx, institution)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return i, err
}

// GetInstitution returns a Institution by ID
func (s *Service) GetInstitution(ctx context.Context, id uint64) (*models.Institution, error) {
	logger := log.With(s.logger, "method", "GetInstitution")

	i, err := s.repository.GetInstitution(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return i, err
}

// GetAllInstitutions returns all Institutions
func (s *Service) GetAllInstitutions(ctx context.Context) ([]*models.Institution, error) {
	logger := log.With(s.logger, "method", "GetAllInstitutions")

	i, err := s.repository.GetAllInstitutions(ctx)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return i, err
}

/* --------------- Course --------------- */

// CreateCourse creates a new Course
func (s *Service) CreateCourse(ctx context.Context, course *models.Course) (*models.Course, error) {
	logger := log.With(s.logger, "method", "CreateCourse")

	c, err := s.repository.CreateCourse(ctx, course)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return c, err
}

// GetCourse returns a Course by ID
func (s *Service) GetCourse(ctx context.Context, id uint64) (*models.Course, error) {
	logger := log.With(s.logger, "method", "GetCourse")

	c, err := s.repository.GetCourse(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return c, err
}

// GetAllCourses returns all Courses
func (s *Service) GetAllCourses(ctx context.Context) ([]*models.Course, error) {
	logger := log.With(s.logger, "method", "GetAllCourses")

	c, err := s.repository.GetAllCourses(ctx)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return c, err
}

/* --------------- Academic History --------------- */

// CreateAcademicHistory creates a new AcademicHistory
func (s *Service) CreateAcademicHistory(ctx context.Context, academic *models.AcademicHistory) (*models.AcademicHistory, error) {
	logger := log.With(s.logger, "method", "CreateAcademicHistory")

	a, err := s.repository.CreateAcademicHistory(ctx, academic)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return a, err
}

// GetAcademicHistory returns a AcademicHistory by ID
func (s *Service) GetAcademicHistory(ctx context.Context, id uint64) (*models.AcademicHistory, error) {
	logger := log.With(s.logger, "method", "GetAcademicHistory")

	a, err := s.repository.GetAcademicHistory(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return a, err
}

// UpdateAcademicHistory updates a AcademicHistory
func (s *Service) UpdateAcademicHistory(ctx context.Context, academic *models.AcademicHistory) (*models.AcademicHistory, error) {
	logger := log.With(s.logger, "method", "UpdateAcademicHistory")

	a, err := s.repository.UpdateAcademicHistory(ctx, academic)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return a, err
}

// DeleteAcademicHistory deletes a AcademicHistory by ID
func (s *Service) DeleteAcademicHistory(ctx context.Context, id uint64) error {
	logger := log.With(s.logger, "method", "DeleteAcademicHistory")

	err := s.repository.DeleteAcademicHistory(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return err
}

/* --------------- Company --------------- */

// CreateCompany creates a new Company
func (s *Service) CreateCompany(ctx context.Context, company *models.Company) (*models.Company, error) {
	logger := log.With(s.logger, "method", "CreateCompany")

	c, err := s.repository.CreateCompany(ctx, company)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return c, err
}

// GetCompany returns a Company by ID
func (s *Service) GetCompany(ctx context.Context, id uint64) (*models.Company, error) {
	logger := log.With(s.logger, "method", "GetCompany")

	c, err := s.repository.GetCompany(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return c, err
}

// GetAllCompanies returns all Companies
func (s *Service) GetAllCompanies(ctx context.Context) ([]*models.Company, error) {
	logger := log.With(s.logger, "method", "GetAllCompanies")

	c, err := s.repository.GetAllCompanies(ctx)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return c, err
}

/* --------------- Department --------------- */

// CreateDepartment creates a new Department
func (s *Service) CreateDepartment(ctx context.Context, course *models.Department) (*models.Department, error) {
	logger := log.With(s.logger, "method", "CreateDepartment")

	d, err := s.repository.CreateDepartment(ctx, course)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return d, err
}

// GetDepartment returns a Department by ID
func (s *Service) GetDepartment(ctx context.Context, id uint64) (*models.Department, error) {
	logger := log.With(s.logger, "method", "GetDepartment")

	d, err := s.repository.GetDepartment(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return d, err
}

// GetAllDepartments returns all Departments
func (s *Service) GetAllDepartments(ctx context.Context) ([]*models.Department, error) {
	logger := log.With(s.logger, "method", "GetAllDepartments")

	d, err := s.repository.GetAllDepartments(ctx)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return d, err
}

/* --------------- Job History --------------- */

// CreateJobHistory creates a new JobHistory
func (s *Service) CreateJobHistory(ctx context.Context, job *models.JobHistory) (*models.JobHistory, error) {
	logger := log.With(s.logger, "method", "CreateJobHistory")

	j, err := s.repository.CreateJobHistory(ctx, job)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return j, err
}

// GetJobHistory returns a JobHistory by ID
func (s *Service) GetJobHistory(ctx context.Context, id uint64) (*models.JobHistory, error) {
	logger := log.With(s.logger, "method", "GetJobHistory")

	j, err := s.repository.GetJobHistory(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return j, err
}

// UpdateJobHistory updates a JobHistory
func (s *Service) UpdateJobHistory(ctx context.Context, job *models.JobHistory) (*models.JobHistory, error) {
	logger := log.With(s.logger, "method", "UpdateJobHistory")

	j, err := s.repository.UpdateJobHistory(ctx, job)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return j, err
}

// DeleteJobHistory deletes a JobHistory by ID
func (s *Service) DeleteJobHistory(ctx context.Context, id uint64) error {
	logger := log.With(s.logger, "method", "DeleteJobHistory")

	err := s.repository.DeleteJobHistory(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return err
}
