package middlewares

import (
	"context"
	"errors"
	"in-backend/services/joblisting/interfaces"
	"in-backend/services/joblisting/models"
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

	idKey        = "https://hubbedin/id"
	companyIDKey = "https://hubbedin/companyId"
	rolesKey     = "https://hubbedin/roles"
)

// NewAuthMiddleware creates and returns a new Auth Middleware that implements the joblisting Service interface
func NewAuthMiddleware(svc interfaces.Service, r interfaces.Repository) interfaces.Service {
	return &authMiddleware{
		next:       svc,
		repository: r,
	}
}

func getRoleAndID(ctx context.Context, ownerID *uint64) (role string, id uint64, owns bool, err error) {
	role = ""
	id = 0
	owns = false
	var companyID uint64 = 0

	claims, err := getClaims(ctx)
	if err != nil {
		return
	}

	if claims[rolesKey] != nil {
		for _, r := range claims[rolesKey].([]interface{}) {
			role = r.(string)
		}
	}

	if claims[companyIDKey] != nil {
		companyID, err = strconv.ParseUint(claims[companyIDKey].(string), 10, 64)
		if err != nil {
			return
		}
		if ownerID != nil && companyID == *ownerID {
			owns = true
		}
	}

	if claims[idKey] != nil {
		id, err = strconv.ParseUint(claims[idKey].(string), 10, 64)
		if err != nil {
			return
		}
	}

	return
}

func getClaims(ctx context.Context) (jwt.MapClaims, error) {
	headers, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errAuth
	}

	if len(headers["authorization"]) == 0 {
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

/* --------------- Job Post --------------- */

// CreateJobPost creates a new JobPost
func (mw authMiddleware) CreateJobPost(ctx context.Context, model *models.JobPost) (*models.JobPost, error) {
	role, _, owns, err := getRoleAndID(ctx, &model.CompanyID)
	if err != nil {
		return nil, err
	}
	if role != "Admin" && owns == false {
		return nil, errAuth
	}
	return mw.next.CreateJobPost(ctx, model)
}

// BulkCreateJobPost creates multiple JobPosts
func (mw authMiddleware) BulkCreateJobPost(ctx context.Context, models []*models.JobPost) ([]*models.JobPost, error) {
	role, _, owns, err := getRoleAndID(ctx, nil)
	if err != nil {
		return nil, err
	}
	if role != "Admin" && owns == false {
		return nil, errAuth
	}
	return mw.next.BulkCreateJobPost(ctx, models)
}

// GetAllJobPosts returns all JobPosts that match the filters
func (mw authMiddleware) GetAllJobPosts(ctx context.Context, f models.JobPostFilters) ([]*models.JobPost, error) {
	return mw.next.GetAllJobPosts(ctx, f)
}

// GetJobPostByID finds and returns a JobPost by ID
func (mw authMiddleware) GetJobPostByID(ctx context.Context, id uint64) (*models.JobPost, error) {
	return mw.next.GetJobPostByID(ctx, id)
}

// UpdateJobPost updates a JobPost
func (mw authMiddleware) UpdateJobPost(ctx context.Context, model *models.JobPost) (*models.JobPost, error) {
	j, err := mw.repository.GetJobPostByID(ctx, model.ID)
	if err != nil {
		return nil, err
	}

	if j.CompanyID != model.CompanyID {
		return nil, errAuth
	}

	role, _, owns, err := getRoleAndID(ctx, &j.CompanyID)
	if err != nil {
		return nil, err
	}
	if role != "Admin" && owns == false {
		return nil, errAuth
	}
	return mw.next.UpdateJobPost(ctx, model)
}

// DeleteJobPost deletes a JobPost by ID
func (mw authMiddleware) DeleteJobPost(ctx context.Context, id uint64) error {
	j, err := mw.repository.GetJobPostByID(ctx, id)
	if err != nil {
		return err
	}

	if j.CompanyID != id {
		return errAuth
	}

	role, _, owns, err := getRoleAndID(ctx, &j.CompanyID)
	if err != nil {
		return err
	}
	if role != "Admin" && owns == false {
		return errAuth
	}
	return mw.next.DeleteJobPost(ctx, id)
}

/* --------------- Company --------------- */

// CreateCompany creates a new Company
func (mw authMiddleware) CreateCompany(ctx context.Context, model *models.Company) (*models.Company, error) {
	role, _, _, err := getRoleAndID(ctx, nil)
	if err != nil {
		return nil, err
	}
	if role != "Admin" && role != "Company" {
		return nil, errAuth
	}
	return mw.next.CreateCompany(ctx, model)
}

// LocalCreateCompany creates a new Company
func (mw authMiddleware) LocalCreateCompany(ctx context.Context, model *models.Company) (*models.Company, error) {
	return mw.next.LocalCreateCompany(ctx, model)
}

// GetAllCompanies returns all Companies that match the filters
func (mw authMiddleware) GetAllCompanies(ctx context.Context, f models.CompanyFilters) ([]*models.Company, error) {
	return mw.next.GetAllCompanies(ctx, f)
}

// UpdateCompany updates a Company
func (mw authMiddleware) UpdateCompany(ctx context.Context, model *models.Company) (*models.Company, error) {
	role, _, owns, err := getRoleAndID(ctx, &model.ID)
	if err != nil {
		return nil, err
	}
	if role != "Admin" && owns == false {
		return nil, errAuth
	}
	return mw.next.UpdateCompany(ctx, model)
}

// LocalUpdateCompany updates a new Company
func (mw authMiddleware) LocalUpdateCompany(ctx context.Context, model *models.Company) (*models.Company, error) {
	return mw.next.LocalUpdateCompany(ctx, model)
}

// DeleteCompany deletes a Company by ID
func (mw authMiddleware) DeleteCompany(ctx context.Context, id uint64) error {
	role, _, owns, err := getRoleAndID(ctx, &id)
	if err != nil {
		return err
	}
	if role != "Admin" && owns == false {
		return errAuth
	}
	return mw.next.DeleteCompany(ctx, id)
}

/* --------------- Industry --------------- */

// CreateIndustry creates a new Industry
func (mw authMiddleware) CreateIndustry(ctx context.Context, model *models.Industry) (*models.Industry, error) {
	role, _, _, err := getRoleAndID(ctx, nil)
	if err != nil {
		return nil, err
	}
	if role != "Admin" && role != "Company" {
		return nil, errAuth
	}
	return mw.next.CreateIndustry(ctx, model)
}

// GetAllIndustries returns all Industries
func (mw authMiddleware) GetAllIndustries(ctx context.Context, f models.IndustryFilters) ([]*models.Industry, error) {
	return mw.next.GetAllIndustries(ctx, f)
}

// DeleteIndustry deletes a Industry by ID
func (mw authMiddleware) DeleteIndustry(ctx context.Context, id uint64) error {
	role, _, _, err := getRoleAndID(ctx, nil)
	if err != nil {
		return err
	}
	if role != "Admin" {
		return errAuth
	}
	return mw.next.DeleteIndustry(ctx, id)
}

/* --------------- Job Function --------------- */

// CreateJobFunction creates a new JobFunction
func (mw authMiddleware) CreateJobFunction(ctx context.Context, model *models.JobFunction) (*models.JobFunction, error) {
	role, _, _, err := getRoleAndID(ctx, nil)
	if err != nil {
		return nil, err
	}
	if role != "Admin" && role != "Company" {
		return nil, errAuth
	}
	return mw.next.CreateJobFunction(ctx, model)
}

// GetAllJobFunctions returns all JobFunctions
func (mw authMiddleware) GetAllJobFunctions(ctx context.Context, f models.JobFunctionFilters) ([]*models.JobFunction, error) {
	return mw.next.GetAllJobFunctions(ctx, f)
}

// DeleteJobFunction deletes a JobFunction by ID
func (mw authMiddleware) DeleteJobFunction(ctx context.Context, id uint64) error {
	role, _, _, err := getRoleAndID(ctx, nil)
	if err != nil {
		return err
	}
	if role != "Admin" {
		return errAuth
	}
	return mw.next.DeleteJobFunction(ctx, id)
}

/* --------------- Key Person --------------- */

// CreateKeyPerson creates a new KeyPerson
func (mw authMiddleware) CreateKeyPerson(ctx context.Context, model *models.KeyPerson) (*models.KeyPerson, error) {
	role, _, owns, err := getRoleAndID(ctx, &model.CompanyID)
	if err != nil {
		return nil, err
	}
	if role != "Admin" && owns == false {
		return nil, errAuth
	}
	return mw.next.CreateKeyPerson(ctx, model)
}

// BulkCreateKeyPerson creates multiple KeyPersons
func (mw authMiddleware) BulkCreateKeyPerson(ctx context.Context, models []*models.KeyPerson) ([]*models.KeyPerson, error) {
	role, _, _, err := getRoleAndID(ctx, nil)
	if err != nil {
		return nil, err
	}
	if role != "Admin" {
		return nil, errAuth
	}
	return mw.next.BulkCreateKeyPerson(ctx, models)
}

// GetAllKeyPersons returns all KeyPersons that match the filters
func (mw authMiddleware) GetAllKeyPersons(ctx context.Context, f models.KeyPersonFilters) ([]*models.KeyPerson, error) {
	role, _, _, err := getRoleAndID(ctx, nil)
	if err != nil {
		return nil, err
	}
	if role != "Admin" {
		return nil, errAuth
	}
	return mw.next.GetAllKeyPersons(ctx, f)
}

// GetKeyPersonByID finds and returns a KeyPerson by ID
func (mw authMiddleware) GetKeyPersonByID(ctx context.Context, id uint64) (*models.KeyPerson, error) {
	kp, err := mw.repository.GetKeyPersonByID(ctx, id)
	if err != nil {
		return nil, err
	}

	role, _, owns, err := getRoleAndID(ctx, &kp.CompanyID)
	if err != nil {
		return nil, err
	}
	if role != "Admin" && owns == false {
		return nil, errAuth
	}
	return mw.next.GetKeyPersonByID(ctx, id)
}

// UpdateKeyPerson updates a KeyPerson
func (mw authMiddleware) UpdateKeyPerson(ctx context.Context, model *models.KeyPerson) (*models.KeyPerson, error) {
	kp, err := mw.repository.GetKeyPersonByID(ctx, model.ID)
	if err != nil {
		return nil, err
	}

	role, _, owns, err := getRoleAndID(ctx, &kp.CompanyID)
	if err != nil {
		return nil, err
	}
	if role != "Admin" && owns == false {
		return nil, errAuth
	}
	return mw.next.UpdateKeyPerson(ctx, model)
}

// DeleteKeyPerson deletes a KeyPerson by ID
func (mw authMiddleware) DeleteKeyPerson(ctx context.Context, id uint64) error {
	kp, err := mw.repository.GetKeyPersonByID(ctx, id)
	if err != nil {
		return err
	}

	role, _, owns, err := getRoleAndID(ctx, &kp.CompanyID)
	if err != nil {
		return err
	}
	if role != "Admin" && owns == false {
		return errAuth
	}
	return mw.next.DeleteKeyPerson(ctx, id)
}

/* --------------- Job Platform --------------- */

// CreateJobPlatform creates a new JobPlatform
func (mw authMiddleware) CreateJobPlatform(ctx context.Context, model *models.JobPlatform) (*models.JobPlatform, error) {
	role, _, _, err := getRoleAndID(ctx, nil)
	if err != nil {
		return nil, err
	}
	if role != "Admin" {
		return nil, errAuth
	}
	return mw.next.CreateJobPlatform(ctx, model)
}

// GetAllJobPlatforms returns all JobPlatforms
func (mw authMiddleware) GetAllJobPlatforms(ctx context.Context, f models.JobPlatformFilters) ([]*models.JobPlatform, error) {
	return mw.next.GetAllJobPlatforms(ctx, f)
}

// DeleteJobPlatform deletes a JobPlatform by ID
func (mw authMiddleware) DeleteJobPlatform(ctx context.Context, id uint64) error {
	role, _, _, err := getRoleAndID(ctx, nil)
	if err != nil {
		return err
	}
	if role != "Admin" {
		return errAuth
	}
	return mw.next.DeleteJobPlatform(ctx, id)
}
