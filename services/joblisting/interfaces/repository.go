package interfaces

import (
	"context"
	"in-backend/services/joblisting/models"
)

// Repository declares the repository for joblistings
type Repository interface {
	/* --------------- Job Post --------------- */

	// CreateJobPost creates a new JobPost
	CreateJobPost(ctx context.Context, m *models.JobPost) (*models.JobPost, error)

	// BulkCreateJobPost creates multiple JobPosts
	BulkCreateJobPost(ctx context.Context, m []*models.JobPost) ([]*models.JobPost, error)

	// GetAllJobPosts returns all JobPosts that match the filters
	GetAllJobPosts(ctx context.Context, f models.JobPostFilters) ([]*models.JobPost, error)

	// GetJobPostByID finds and returns a JobPost by ID
	GetJobPostByID(ctx context.Context, id uint64) (*models.JobPost, error)

	// UpdateJobPost updates a JobPost
	UpdateJobPost(ctx context.Context, m *models.JobPost) (*models.JobPost, error)

	// DeleteJobPost deletes a JobPost by ID
	DeleteJobPost(ctx context.Context, id uint64) error

	/* --------------- Company --------------- */

	// CreateCompany creates a new Company
	CreateCompany(ctx context.Context, m *models.Company) (*models.Company, error)

	// GetAllCompanies returns all Companies that match the filters
	GetAllCompanies(ctx context.Context, f models.CompanyFilters) ([]*models.Company, error)

	// UpdateCompany updates a Company
	UpdateCompany(ctx context.Context, m *models.Company) (*models.Company, error)

	// DeleteCompany deletes a Company by ID
	DeleteCompany(ctx context.Context, id uint64) error

	/* --------------- Industry --------------- */

	// CreateIndustry creates a new Industry
	CreateIndustry(ctx context.Context, m *models.Industry) (*models.Industry, error)

	// GetAllIndustries returns all Industries
	GetAllIndustries(ctx context.Context, f models.IndustryFilters) ([]*models.Industry, error)

	// DeleteIndustry deletes a Industry by ID
	DeleteIndustry(ctx context.Context, id uint64) error

	/* --------------- Job Function --------------- */

	// CreateJobFunction creates a new JobFunction
	CreateJobFunction(ctx context.Context, m *models.JobFunction) (*models.JobFunction, error)

	// GetAllJobFunctions returns all JobFunctions
	GetAllJobFunctions(ctx context.Context, f models.JobFunctionFilters) ([]*models.JobFunction, error)

	// DeleteJobFunction deletes a JobFunction by ID
	DeleteJobFunction(ctx context.Context, id uint64) error

	/* --------------- Key Person --------------- */

	// CreateKeyPerson creates a new KeyPerson
	CreateKeyPerson(ctx context.Context, m *models.KeyPerson) (*models.KeyPerson, error)

	// BulkCreateKeyPerson creates multiple KeyPersons
	BulkCreateKeyPerson(ctx context.Context, m []*models.KeyPerson) ([]*models.KeyPerson, error)

	// GetAllKeyPersons returns all KeyPersons that match the filters
	GetAllKeyPersons(ctx context.Context, f models.KeyPersonFilters) ([]*models.KeyPerson, error)

	// GetKeyPersonByID finds and returns a KeyPerson by ID
	GetKeyPersonByID(ctx context.Context, id uint64) (*models.KeyPerson, error)

	// UpdateKeyPerson updates a KeyPerson
	UpdateKeyPerson(ctx context.Context, m *models.KeyPerson) (*models.KeyPerson, error)

	// DeleteKeyPerson deletes a KeyPerson by ID
	DeleteKeyPerson(ctx context.Context, id uint64) error

	/* --------------- Job Platform --------------- */

	// CreateJobPlatform creates a new JobPlatform
	CreateJobPlatform(ctx context.Context, m *models.JobPlatform) (*models.JobPlatform, error)

	// GetAllJobPlatforms returns all JobPlatforms
	GetAllJobPlatforms(ctx context.Context, f models.JobPlatformFilters) ([]*models.JobPlatform, error)

	// DeleteJobPlatform deletes a JobPlatform by ID
	DeleteJobPlatform(ctx context.Context, id uint64) error
}
