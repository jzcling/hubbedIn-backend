package service

import (
	"context"

	"in-backend/services/profile/interfaces"
	"in-backend/services/profile/models"
)

// Service implements the profile Service interface
type service struct {
	repository interfaces.Repository
}

// New creates and returns a new Service that implements the profile Service interface
func New(r interfaces.Repository) interfaces.Service {
	return &service{
		repository: r,
	}
}

/* --------------- Candidate --------------- */

// CreateCandidate creates a new Candidate
func (s *service) CreateCandidate(ctx context.Context, candidate *models.Candidate) (*models.Candidate, error) {
	c, err := s.repository.CreateCandidate(ctx, candidate)
	if err != nil {
		return nil, err
	}

	return c, err
}

// GetAllCandidates returns all Candidates
func (s *service) GetAllCandidates(ctx context.Context, f models.CandidateFilters) ([]*models.Candidate, error) {
	c, err := s.repository.GetAllCandidates(ctx, f)
	if err != nil {
		return nil, err
	}
	return c, err
}

// GetCandidateByID returns a Candidate by ID
func (s *service) GetCandidateByID(ctx context.Context, id uint64) (*models.Candidate, error) {
	c, err := s.repository.GetCandidateByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return c, err
}

// UpdateCandidate updates a Candidate
func (s *service) UpdateCandidate(ctx context.Context, candidate *models.Candidate) (*models.Candidate, error) {
	c, err := s.repository.UpdateCandidate(ctx, candidate)
	if err != nil {
		return nil, err
	}
	return c, err
}

// DeleteCandidate deletes a Candidate by ID
func (s *service) DeleteCandidate(ctx context.Context, id uint64) error {
	err := s.repository.DeleteCandidate(ctx, id)
	if err != nil {
		return err
	}
	return err
}

/* --------------- Skill --------------- */

// CreateSkill creates a new Skill
func (s *service) CreateSkill(ctx context.Context, skill *models.Skill) (*models.Skill, error) {
	sk, err := s.repository.CreateSkill(ctx, skill)
	if err != nil {
		return nil, err
	}
	return sk, err
}

// GetSkill returns a Skill by ID
func (s *service) GetSkill(ctx context.Context, id uint64) (*models.Skill, error) {
	sk, err := s.repository.GetSkill(ctx, id)
	if err != nil {
		return nil, err
	}
	return sk, err
}

// GetAllSkills returns all Skills
func (s *service) GetAllSkills(ctx context.Context, f models.SkillFilters) ([]*models.Skill, error) {
	sk, err := s.repository.GetAllSkills(ctx, f)
	if err != nil {
		return nil, err
	}
	return sk, err
}

/* --------------- User Skill --------------- */

// CreateUserSkill creates a new UserSkill
func (s *service) CreateUserSkill(ctx context.Context, us *models.UserSkill) (*models.UserSkill, error) {
	us, err := s.repository.CreateUserSkill(ctx, us)
	if err != nil {
		return nil, err
	}

	return us, err
}

// DeleteUserSkill deletes a UserSkill by ID
func (s *service) DeleteUserSkill(ctx context.Context, cid, sid uint64) error {
	err := s.repository.DeleteUserSkill(ctx, cid, sid)
	if err != nil {
		return err
	}
	return err
}

/* --------------- Institution --------------- */

// CreateInstitution creates a new Institution
func (s *service) CreateInstitution(ctx context.Context, institution *models.Institution) (*models.Institution, error) {
	i, err := s.repository.CreateInstitution(ctx, institution)
	if err != nil {
		return nil, err
	}
	return i, err
}

// GetInstitution returns a Institution by ID
func (s *service) GetInstitution(ctx context.Context, id uint64) (*models.Institution, error) {
	i, err := s.repository.GetInstitution(ctx, id)
	if err != nil {
		return nil, err
	}
	return i, err
}

// GetAllInstitutions returns all Institutions
func (s *service) GetAllInstitutions(ctx context.Context, f models.InstitutionFilters) ([]*models.Institution, error) {
	i, err := s.repository.GetAllInstitutions(ctx, f)
	if err != nil {
		return nil, err
	}
	return i, err
}

/* --------------- Course --------------- */

// CreateCourse creates a new Course
func (s *service) CreateCourse(ctx context.Context, course *models.Course) (*models.Course, error) {
	c, err := s.repository.CreateCourse(ctx, course)
	if err != nil {
		return nil, err
	}
	return c, err
}

// GetCourse returns a Course by ID
func (s *service) GetCourse(ctx context.Context, id uint64) (*models.Course, error) {
	c, err := s.repository.GetCourse(ctx, id)
	if err != nil {
		return nil, err
	}
	return c, err
}

// GetAllCourses returns all Courses
func (s *service) GetAllCourses(ctx context.Context, f models.CourseFilters) ([]*models.Course, error) {
	c, err := s.repository.GetAllCourses(ctx, f)
	if err != nil {
		return nil, err
	}
	return c, err
}

/* --------------- Academic History --------------- */

// CreateAcademicHistory creates a new AcademicHistory
func (s *service) CreateAcademicHistory(ctx context.Context, academic *models.AcademicHistory) (*models.AcademicHistory, error) {
	a, err := s.repository.CreateAcademicHistory(ctx, academic)
	if err != nil {
		return nil, err
	}
	return a, err
}

// GetAcademicHistory returns a AcademicHistory by ID
func (s *service) GetAcademicHistory(ctx context.Context, id uint64) (*models.AcademicHistory, error) {
	a, err := s.repository.GetAcademicHistory(ctx, id)
	if err != nil {
		return nil, err
	}
	return a, err
}

// UpdateAcademicHistory updates a AcademicHistory
func (s *service) UpdateAcademicHistory(ctx context.Context, academic *models.AcademicHistory) (*models.AcademicHistory, error) {
	a, err := s.repository.UpdateAcademicHistory(ctx, academic)
	if err != nil {
		return nil, err
	}
	return a, err
}

// DeleteAcademicHistory deletes a AcademicHistory by ID
func (s *service) DeleteAcademicHistory(ctx context.Context, cid, ahid uint64) error {
	err := s.repository.DeleteAcademicHistory(ctx, cid, ahid)
	if err != nil {
		return err
	}
	return err
}

/* --------------- Company --------------- */

// CreateCompany creates a new Company
func (s *service) CreateCompany(ctx context.Context, company *models.Company) (*models.Company, error) {
	c, err := s.repository.CreateCompany(ctx, company)
	if err != nil {
		return nil, err
	}
	return c, err
}

// GetCompany returns a Company by ID
func (s *service) GetCompany(ctx context.Context, id uint64) (*models.Company, error) {

	c, err := s.repository.GetCompany(ctx, id)
	if err != nil {
		return nil, err
	}
	return c, err
}

// GetAllCompanies returns all Companies
func (s *service) GetAllCompanies(ctx context.Context, f models.CompanyFilters) ([]*models.Company, error) {
	c, err := s.repository.GetAllCompanies(ctx, f)
	if err != nil {
		return nil, err
	}
	return c, err
}

/* --------------- Department --------------- */

// CreateDepartment creates a new Department
func (s *service) CreateDepartment(ctx context.Context, department *models.Department) (*models.Department, error) {
	d, err := s.repository.CreateDepartment(ctx, department)
	if err != nil {
		return nil, err
	}
	return d, err
}

// GetDepartment returns a Department by ID
func (s *service) GetDepartment(ctx context.Context, id uint64) (*models.Department, error) {
	d, err := s.repository.GetDepartment(ctx, id)
	if err != nil {
		return nil, err
	}
	return d, err
}

// GetAllDepartments returns all Departments
func (s *service) GetAllDepartments(ctx context.Context, f models.DepartmentFilters) ([]*models.Department, error) {
	d, err := s.repository.GetAllDepartments(ctx, f)
	if err != nil {
		return nil, err
	}
	return d, err
}

/* --------------- Job History --------------- */

// CreateJobHistory creates a new JobHistory
func (s *service) CreateJobHistory(ctx context.Context, job *models.JobHistory) (*models.JobHistory, error) {
	j, err := s.repository.CreateJobHistory(ctx, job)
	if err != nil {
		return nil, err
	}
	return j, err
}

// GetJobHistory returns a JobHistory by ID
func (s *service) GetJobHistory(ctx context.Context, id uint64) (*models.JobHistory, error) {
	j, err := s.repository.GetJobHistory(ctx, id)
	if err != nil {
		return nil, err
	}
	return j, err
}

// UpdateJobHistory updates a JobHistory
func (s *service) UpdateJobHistory(ctx context.Context, job *models.JobHistory) (*models.JobHistory, error) {
	j, err := s.repository.UpdateJobHistory(ctx, job)
	if err != nil {
		return nil, err
	}
	return j, err
}

// DeleteJobHistory deletes a JobHistory by ID
func (s *service) DeleteJobHistory(ctx context.Context, cid, jhid uint64) error {
	err := s.repository.DeleteJobHistory(ctx, cid, jhid)
	if err != nil {
		return err
	}
	return err
}
