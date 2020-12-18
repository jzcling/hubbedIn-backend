package middlewares

import (
	"context"
	"errors"
	"in-backend/services/profile/interfaces"
	"in-backend/services/profile/models"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/metadata"
)

type authMiddleware struct {
	next       interfaces.Service
	repository interfaces.Repository
}

var (
	errAuth = errors.New("Forbidden")

	idKey    = "https://hubbedin/id"
	rolesKey = "https://hubbedin/roles"
)

// NewAuthMiddleware creates and returns a new Auth Middleware that implements the profile Service interface
func NewAuthMiddleware(svc interfaces.Service, r interfaces.Repository) interfaces.Service {
	return &authMiddleware{
		next:       svc,
		repository: r,
	}
}

func getRoleAndID(ctx context.Context, ownerID *uint64) (*string, *uint64, error) {
	claims, err := getClaims(ctx)
	if err != nil {
		return nil, nil, err
	}

	var role string = ""
	var id uint64 = 0

	// this should come first so that role gets overwritten if owner is also an admin
	if claims[idKey] != nil {
		id, err = strconv.ParseUint(claims[idKey].(string), 10, 64)
		if err != nil {
			return nil, nil, err
		}
		if ownerID != nil && id == *ownerID {
			role = "Owner"
		}
	}

	if claims[rolesKey] != nil {
		for _, r := range claims[rolesKey].([]interface{}) {
			roleCast := r.(string)
			if roleCast == "Admin" {
				role = "Admin"
			}
		}
	}

	return &role, &id, nil
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
	role, _, err := getRoleAndID(ctx, nil)
	if err != nil {
		return nil, err
	}
	if *role != "Admin" {
		return nil, errAuth
	}
	return mw.next.GetAllCandidates(ctx, f)
}

// GetCandidateByID returns a Candidate by ID
func (mw authMiddleware) GetCandidateByID(ctx context.Context, id uint64) (*models.Candidate, error) {
	role, _, err := getRoleAndID(ctx, &id)
	if err != nil {
		return nil, err
	}
	if *role != "Admin" && *role != "Owner" {
		return nil, errAuth
	}
	return mw.next.GetCandidateByID(ctx, id)
}

// UpdateCandidate updates a Candidate
func (mw authMiddleware) UpdateCandidate(ctx context.Context, m *models.Candidate) (*models.Candidate, error) {
	c, err := mw.repository.GetCandidateByID(ctx, m.ID)
	if err != nil {
		return nil, err
	}
	role, _, err := getRoleAndID(ctx, &c.ID)
	if err != nil {
		return nil, err
	}
	if *role != "Admin" && *role != "Owner" {
		return nil, errAuth
	}
	return mw.next.UpdateCandidate(ctx, m)
}

// DeleteCandidate deletes a Candidate by ID
func (mw authMiddleware) DeleteCandidate(ctx context.Context, id uint64) error {
	role, _, err := getRoleAndID(ctx, &id)
	if err != nil {
		return err
	}
	if *role != "Admin" && *role != "Owner" {
		return errAuth
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
func (mw authMiddleware) CreateUserSkill(ctx context.Context, m *models.UserSkill) (*models.UserSkill, error) {
	role, _, err := getRoleAndID(ctx, &m.CandidateID)
	if err != nil {
		return nil, err
	}
	if *role != "Admin" && *role != "Owner" {
		return nil, errAuth
	}
	return mw.next.CreateUserSkill(ctx, m)
}

// DeleteUserSkill deletes a UserSkill by ID
func (mw authMiddleware) DeleteUserSkill(ctx context.Context, cid, sid uint64) error {
	role, _, err := getRoleAndID(ctx, &cid)
	if err != nil {
		return err
	}
	if *role != "Admin" && *role != "Owner" {
		return errAuth
	}
	return mw.next.DeleteUserSkill(ctx, cid, sid)
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
func (mw authMiddleware) CreateAcademicHistory(ctx context.Context, m *models.AcademicHistory) (*models.AcademicHistory, error) {
	role, _, err := getRoleAndID(ctx, &m.CandidateID)
	if err != nil {
		return nil, err
	}
	if *role != "Admin" && *role != "Owner" {
		return nil, errAuth
	}
	return mw.next.CreateAcademicHistory(ctx, m)
}

// GetAcademicHistory returns a AcademicHistory by ID
func (mw authMiddleware) GetAcademicHistory(ctx context.Context, id uint64) (*models.AcademicHistory, error) {
	ah, err := mw.repository.GetAcademicHistory(ctx, id)
	if err != nil {
		return nil, err
	}
	role, _, err := getRoleAndID(ctx, &ah.CandidateID)
	if err != nil {
		return nil, err
	}
	if *role != "Admin" && *role != "Owner" {
		return nil, errAuth
	}
	return mw.next.GetAcademicHistory(ctx, id)
}

// UpdateAcademicHistory updates a AcademicHistory
func (mw authMiddleware) UpdateAcademicHistory(ctx context.Context, m *models.AcademicHistory) (*models.AcademicHistory, error) {
	ah, err := mw.repository.GetAcademicHistory(ctx, m.ID)
	if err != nil {
		return nil, err
	}
	role, _, err := getRoleAndID(ctx, &ah.CandidateID)
	if err != nil {
		return nil, err
	}
	if *role != "Admin" && *role != "Owner" {
		return nil, errAuth
	}
	return mw.next.UpdateAcademicHistory(ctx, m)
}

// DeleteAcademicHistory deletes a AcademicHistory by ID
func (mw authMiddleware) DeleteAcademicHistory(ctx context.Context, cid, ahid uint64) error {
	role, _, err := getRoleAndID(ctx, &cid)
	if err != nil {
		return err
	}
	if *role != "Admin" && *role != "Owner" {
		return errAuth
	}
	return mw.next.DeleteAcademicHistory(ctx, cid, ahid)
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
func (mw authMiddleware) CreateJobHistory(ctx context.Context, m *models.JobHistory) (*models.JobHistory, error) {
	role, _, err := getRoleAndID(ctx, &m.CandidateID)
	if err != nil {
		return nil, err
	}
	if *role != "Admin" && *role != "Owner" {
		return nil, errAuth
	}
	return mw.next.CreateJobHistory(ctx, m)
}

// GetJobHistory returns a JobHistory by ID
func (mw authMiddleware) GetJobHistory(ctx context.Context, id uint64) (*models.JobHistory, error) {
	jh, err := mw.repository.GetJobHistory(ctx, id)
	if err != nil {
		return nil, err
	}
	role, _, err := getRoleAndID(ctx, &jh.CandidateID)
	if err != nil {
		return nil, err
	}
	if *role != "Admin" && *role != "Owner" {
		return nil, errAuth
	}
	return mw.next.GetJobHistory(ctx, id)
}

// UpdateJobHistory updates a JobHistory
func (mw authMiddleware) UpdateJobHistory(ctx context.Context, m *models.JobHistory) (*models.JobHistory, error) {
	jh, err := mw.repository.GetJobHistory(ctx, m.ID)
	if err != nil {
		return nil, err
	}
	role, _, err := getRoleAndID(ctx, &jh.CandidateID)
	if err != nil {
		return nil, err
	}
	if *role != "Admin" && *role != "Owner" {
		return nil, errAuth
	}
	return mw.next.UpdateJobHistory(ctx, m)
}

// DeleteJobHistory deletes a JobHistory by ID
func (mw authMiddleware) DeleteJobHistory(ctx context.Context, cid, jhid uint64) error {
	role, _, err := getRoleAndID(ctx, &cid)
	if err != nil {
		return err
	}
	if *role != "Admin" && *role != "Owner" {
		return errAuth
	}
	return mw.next.DeleteJobHistory(ctx, cid, jhid)
}
