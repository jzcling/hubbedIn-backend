package middlewares

import (
	"context"
	"errors"
	"in-backend/services/assessment/interfaces"
	"in-backend/services/assessment/models"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/metadata"
)

type authMiddleware struct {
	next interfaces.Service
}

var (
	errAuth = errors.New("Forbidden")

	namespace = "https://hubbedin/id"
)

// NewAuthMiddleware creates and returns a new Auth Middleware that implements the assessment Service interface
func NewAuthMiddleware(svc interfaces.Service) interfaces.Service {
	return &authMiddleware{
		next: svc,
	}
}

func getClaims(ctx context.Context) (jwt.MapClaims, error) {
	headers, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errAuth
	}
	tokenString := strings.Split(headers["authorization"][0], " ")[1]
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return nil, errAuth
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errAuth
	}
	return claims, nil
}

/* --------------- Candidate --------------- */

// CreateCandidate creates a new Candidate
func (mw authMiddleware) CreateCandidate(ctx context.Context, candidate *models.Candidate) (*models.Candidate, error) {
	return mw.next.CreateCandidate(ctx, candidate)
}

// GetAllCandidates returns all Candidates
func (mw authMiddleware) GetAllCandidates(ctx context.Context, f models.CandidateFilters) ([]*models.Candidate, error) {
	return mw.next.GetAllCandidates(ctx, f)
}

// GetCandidateByID returns a Candidate by ID
func (mw authMiddleware) GetCandidateByID(ctx context.Context, id uint64) (*models.Candidate, error) {
	claims, err := getClaims(ctx)
	if err != nil {
		return nil, err
	}
	id, err = strconv.ParseUint(claims[namespace].(string), 10, 64)
	if err != nil {
		return nil, err
	}
	return mw.next.GetCandidateByID(ctx, id)
}

// UpdateCandidate updates a Candidate
func (mw authMiddleware) UpdateCandidate(ctx context.Context, candidate *models.Candidate) (*models.Candidate, error) {
	claims, err := getClaims(ctx)
	if err != nil {
		return nil, err
	}
	id, err := strconv.ParseUint(claims[namespace].(string), 10, 64)
	if err != nil {
		return nil, err
	}
	if id != candidate.ID {
		return nil, errAuth
	}
	return mw.next.UpdateCandidate(ctx, candidate)
}

// DeleteCandidate deletes a Candidate by ID
func (mw authMiddleware) DeleteCandidate(ctx context.Context, id uint64) error {
	claims, err := getClaims(ctx)
	if err != nil {
		return err
	}
	id, err = strconv.ParseUint(claims[namespace].(string), 10, 64)
	if err != nil {
		return err
	}
	return mw.next.DeleteCandidate(ctx, id)
}

/* --------------- Skill --------------- */

// CreateSkill creates a new Skill
func (mw authMiddleware) CreateSkill(ctx context.Context, skill *models.Skill) (*models.Skill, error) {
	return mw.next.CreateSkill(ctx, skill)
}

// GetSkill returns a Skill by ID
func (mw authMiddleware) GetSkill(ctx context.Context, id uint64) (*models.Skill, error) {
	return mw.next.GetSkill(ctx, id)
}

// GetAllSkills returns all Skills
func (mw authMiddleware) GetAllSkills(ctx context.Context, f models.SkillFilters) ([]*models.Skill, error) {
	return mw.next.GetAllSkills(ctx, f)
}

/* --------------- User Skill --------------- */

// CreateUserSkill creates a new UserSkill
func (mw authMiddleware) CreateUserSkill(ctx context.Context, us *models.UserSkill) (*models.UserSkill, error) {
	claims, err := getClaims(ctx)
	if err != nil {
		return nil, err
	}
	id, err := strconv.ParseUint(claims[namespace].(string), 10, 64)
	if err != nil {
		return nil, err
	}
	if id != us.CandidateID {
		return nil, errAuth
	}
	return mw.next.CreateUserSkill(ctx, us)
}

// DeleteUserSkill deletes a UserSkill by ID
func (mw authMiddleware) DeleteUserSkill(ctx context.Context, id uint64) error {
	return mw.next.DeleteUserSkill(ctx, id)
}

/* --------------- Institution --------------- */

// CreateInstitution creates a new Institution
func (mw authMiddleware) CreateInstitution(ctx context.Context, institution *models.Institution) (*models.Institution, error) {
	return mw.next.CreateInstitution(ctx, institution)
}

// GetInstitution returns a Institution by ID
func (mw authMiddleware) GetInstitution(ctx context.Context, id uint64) (*models.Institution, error) {
	return mw.next.GetInstitution(ctx, id)
}

// GetAllInstitutions returns all Institutions
func (mw authMiddleware) GetAllInstitutions(ctx context.Context, f models.InstitutionFilters) ([]*models.Institution, error) {
	return mw.next.GetAllInstitutions(ctx, f)
}

/* --------------- Course --------------- */

// CreateCourse creates a new Course
func (mw authMiddleware) CreateCourse(ctx context.Context, course *models.Course) (*models.Course, error) {
	return mw.next.CreateCourse(ctx, course)
}

// GetCourse returns a Course by ID
func (mw authMiddleware) GetCourse(ctx context.Context, id uint64) (*models.Course, error) {
	return mw.next.GetCourse(ctx, id)
}

// GetAllCourses returns all Courses
func (mw authMiddleware) GetAllCourses(ctx context.Context, f models.CourseFilters) ([]*models.Course, error) {
	return mw.next.GetAllCourses(ctx, f)
}

/* --------------- Academic History --------------- */

// CreateAcademicHistory creates a new AcademicHistory
func (mw authMiddleware) CreateAcademicHistory(ctx context.Context, academic *models.AcademicHistory) (*models.AcademicHistory, error) {
	claims, err := getClaims(ctx)
	if err != nil {
		return nil, err
	}
	id, err := strconv.ParseUint(claims[namespace].(string), 10, 64)
	if err != nil {
		return nil, err
	}
	if id != academic.CandidateID {
		return nil, errAuth
	}
	return mw.next.CreateAcademicHistory(ctx, academic)
}

// GetAcademicHistory returns a AcademicHistory by ID
func (mw authMiddleware) GetAcademicHistory(ctx context.Context, id uint64) (*models.AcademicHistory, error) {
	return mw.next.GetAcademicHistory(ctx, id)
}

// UpdateAcademicHistory updates a AcademicHistory
func (mw authMiddleware) UpdateAcademicHistory(ctx context.Context, academic *models.AcademicHistory) (*models.AcademicHistory, error) {
	claims, err := getClaims(ctx)
	if err != nil {
		return nil, err
	}
	id, err := strconv.ParseUint(claims[namespace].(string), 10, 64)
	if err != nil {
		return nil, err
	}
	if id != academic.CandidateID {
		return nil, errAuth
	}
	return mw.next.UpdateAcademicHistory(ctx, academic)
}

// DeleteAcademicHistory deletes a AcademicHistory by ID
func (mw authMiddleware) DeleteAcademicHistory(ctx context.Context, id uint64) error {
	return mw.next.DeleteAcademicHistory(ctx, id)
}

/* --------------- Company --------------- */

// CreateCompany creates a new Company
func (mw authMiddleware) CreateCompany(ctx context.Context, company *models.Company) (*models.Company, error) {
	return mw.next.CreateCompany(ctx, company)
}

// GetCompany returns a Company by ID
func (mw authMiddleware) GetCompany(ctx context.Context, id uint64) (*models.Company, error) {
	return mw.next.GetCompany(ctx, id)
}

// GetAllCompanies returns all Companies
func (mw authMiddleware) GetAllCompanies(ctx context.Context, f models.CompanyFilters) ([]*models.Company, error) {
	return mw.next.GetAllCompanies(ctx, f)
}

/* --------------- Department --------------- */

// CreateDepartment creates a new Department
func (mw authMiddleware) CreateDepartment(ctx context.Context, department *models.Department) (*models.Department, error) {
	return mw.next.CreateDepartment(ctx, department)
}

// GetDepartment returns a Department by ID
func (mw authMiddleware) GetDepartment(ctx context.Context, id uint64) (*models.Department, error) {
	return mw.next.GetDepartment(ctx, id)
}

// GetAllDepartments returns all Departments
func (mw authMiddleware) GetAllDepartments(ctx context.Context, f models.DepartmentFilters) ([]*models.Department, error) {
	return mw.next.GetAllDepartments(ctx, f)
}

/* --------------- Job History --------------- */

// CreateJobHistory creates a new JobHistory
func (mw authMiddleware) CreateJobHistory(ctx context.Context, job *models.JobHistory) (*models.JobHistory, error) {
	claims, err := getClaims(ctx)
	if err != nil {
		return nil, err
	}
	id, err := strconv.ParseUint(claims[namespace].(string), 10, 64)
	if err != nil {
		return nil, err
	}
	if id != job.CandidateID {
		return nil, errAuth
	}
	return mw.next.CreateJobHistory(ctx, job)
}

// GetJobHistory returns a JobHistory by ID
func (mw authMiddleware) GetJobHistory(ctx context.Context, id uint64) (*models.JobHistory, error) {
	return mw.next.GetJobHistory(ctx, id)
}

// UpdateJobHistory updates a JobHistory
func (mw authMiddleware) UpdateJobHistory(ctx context.Context, job *models.JobHistory) (*models.JobHistory, error) {
	claims, err := getClaims(ctx)
	if err != nil {
		return nil, err
	}
	id, err := strconv.ParseUint(claims[namespace].(string), 10, 64)
	if err != nil {
		return nil, err
	}
	if id != job.CandidateID {
		return nil, errAuth
	}
	return mw.next.UpdateJobHistory(ctx, job)
}

// DeleteJobHistory deletes a JobHistory by ID
func (mw authMiddleware) DeleteJobHistory(ctx context.Context, id uint64) error {
	return mw.next.DeleteJobHistory(ctx, id)
}
