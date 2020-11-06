package profile

import (
	"context"
)

// Repository declares the repository for candidate profiles
type Repository interface {
	/* --------------- Candidate --------------- */

	// CreateCandidate creates a new candidate
	CreateCandidate(ctx context.Context, c *orm.Candidate) (*orm.Candidate, error)

	// GetAllCandidates returns all candidates
	GetAllCandidates(ctx context.Context) ([]*orm.Candidate, error)

	// GetCandidateByID finds and returns a candidate by ID
	GetCandidateByID(ctx context.Context, id uint64) (*orm.Candidate, error)

	// UpdateCandidate updates a candidate
	UpdateCandidate(ctx context.Context, c *orm.Candidate) (*orm.Candidate, error)

	// DeleteCandidate deletes a candidate by ID
	DeleteCandidate(ctx context.Context, id uint64) error

	/* --------------- Skill --------------- */

	// CreateSkill creates a new Skill
	CreateSkill(ctx context.Context, s *orm.Skill) (*orm.Skill, error)

	// GetSkill returns a Skill by ID
	GetSkill(ctx context.Context, id uint64) (*orm.Skill, error)

	// GetAllSkills returns all Skills
	GetAllSkills(ctx context.Context) (*orm.Skill, error)

	/* --------------- Institution --------------- */

	// CreateInstitution creates a new Institution
	CreateInstitution(ctx context.Context, i *orm.Institution) (*orm.Institution, error)

	// GetInstitution returns a Institution by ID
	GetInstitution(ctx context.Context, id uint64) (*orm.Institution, error)

	// GetAllInstitutions returns all Institutions
	GetAllInstitutions(ctx context.Context) (*orm.Institution, error)

	/* --------------- Course --------------- */

	// CreateCourse creates a new Course
	CreateCourse(ctx context.Context, c *orm.Course) (*orm.Course, error)

	// GetCourse returns a Course by ID
	GetCourse(ctx context.Context, id uint64) (*orm.Course, error)

	// GetAllCourses returns all Courses
	GetAllCourses(ctx context.Context) (*orm.Course, error)

	/* --------------- Academic History --------------- */

	// CreateAcademicHistory creates a new AcademicHistory
	CreateAcademicHistory(ctx context.Context, a *orm.AcademicHistory) (*orm.AcademicHistory, error)

	// GetAcademicHistory returns a AcademicHistory by ID
	GetAcademicHistory(ctx context.Context, id uint64) (*orm.AcademicHistory, error)

	// UpdateAcademicHistory updates a AcademicHistory
	UpdateAcademicHistory(ctx context.Context, a *orm.AcademicHistory) (*orm.AcademicHistory, error)

	// DeleteAcademicHistory deletes a AcademicHistory by ID
	DeleteAcademicHistory(ctx context.Context, id uint64) error

	/* --------------- Company --------------- */

	// CreateCompany creates a new Company
	CreateCompany(ctx context.Context, c *orm.Company) (*orm.Company, error)

	// GetCompany returns a Company by ID
	GetCompany(ctx context.Context, id uint64) (*orm.Company, error)

	// GetAllCompanies returns all Companies
	GetAllCompanies(ctx context.Context) (*orm.Company, error)

	/* --------------- Department --------------- */

	// CreateDepartment creates a new Department
	CreateDepartment(ctx context.Context, c *orm.Department) (*orm.Department, error)

	// GetDepartment returns a Department by ID
	GetDepartment(ctx context.Context, id uint64) (*orm.Department, error)

	// GetAllDepartments returns all Departments
	GetAllDepartments(ctx context.Context) (*orm.Department, error)

	/* --------------- Job History --------------- */

	// CreateJobHistory creates a new JobHistory
	CreateJobHistory(ctx context.Context, a *orm.JobHistory) (*orm.JobHistory, error)

	// GetJobHistory returns a JobHistory by ID
	GetJobHistory(ctx context.Context, id uint64) (*orm.JobHistory, error)

	// UpdateJobHistory updates a JobHistory
	UpdateJobHistory(ctx context.Context, j *orm.JobHistory) (*orm.JobHistory, error)

	// DeleteJobHistory deletes a JobHistory by ID
	DeleteJobHistory(ctx context.Context, id uint64) error
}
