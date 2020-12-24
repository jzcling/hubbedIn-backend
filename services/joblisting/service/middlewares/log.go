package middlewares

import (
	"context"
	"fmt"
	"in-backend/services/joblisting/interfaces"
	"in-backend/services/joblisting/models"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type logMiddleware struct {
	logger log.Logger
	next   interfaces.Service
}

// NewLogMiddleware creates and returns a new Log Middleware that implements the joblisting Service interface
func NewLogMiddleware(logger log.Logger, svc interfaces.Service) interfaces.Service {
	return &logMiddleware{
		logger: logger,
		next:   svc,
	}
}

func (mw logMiddleware) log(method string, begin time.Time, input, output interface{}, err *error) {
	var logger log.Logger
	if err != nil {
		logger = level.Error(mw.logger)
	} else {
		logger = level.Info(mw.logger)
	}
	logger.Log(
		"method", method,
		"input", fmt.Sprintf("%v", input),
		"output", fmt.Sprintf("%v", output),
		"err", err,
		"took", time.Since(begin),
	)
}

/* --------------- Job Post --------------- */

// CreateJobPost creates a new JobPost
func (mw logMiddleware) CreateJobPost(ctx context.Context, input *models.JobPost) (output *models.JobPost, err error) {
	defer mw.log("CreateJobPost", time.Now(), input, output, &err)
	output, err = mw.next.CreateJobPost(ctx, input)
	return
}

// BulkCreateJobPost creates multiple JobPosts
func (mw logMiddleware) BulkCreateJobPost(ctx context.Context, input []*models.JobPost) (output []*models.JobPost, err error) {
	defer mw.log("BulkCreateJobPost", time.Now(), input, output, &err)
	output, err = mw.next.BulkCreateJobPost(ctx, input)
	return
}

// GetAllJobPosts returns all JobPosts that match the filters
func (mw logMiddleware) GetAllJobPosts(ctx context.Context, input models.JobPostFilters) (output []*models.JobPost, err error) {
	defer mw.log("GetAllJobPosts", time.Now(), input, output, &err)
	output, err = mw.next.GetAllJobPosts(ctx, input)
	return
}

// GetJobPostByID finds and returns a JobPost by ID
func (mw logMiddleware) GetJobPostByID(ctx context.Context, input uint64) (output *models.JobPost, err error) {
	defer mw.log("GetJobPostByID", time.Now(), input, output, &err)
	output, err = mw.next.GetJobPostByID(ctx, input)
	return
}

// UpdateJobPost updates a JobPost
func (mw logMiddleware) UpdateJobPost(ctx context.Context, input *models.JobPost) (output *models.JobPost, err error) {
	defer mw.log("UpdateJobPost", time.Now(), input, output, &err)
	output, err = mw.next.UpdateJobPost(ctx, input)
	return
}

// DeleteJobPost deletes a JobPost by ID
func (mw logMiddleware) DeleteJobPost(ctx context.Context, input uint64) (err error) {
	defer mw.log("DeleteJobPost", time.Now(), input, nil, &err)
	err = mw.next.DeleteJobPost(ctx, input)
	return
}

/* --------------- Company --------------- */

// CreateCompany creates a new Company
func (mw logMiddleware) CreateCompany(ctx context.Context, input *models.Company) (output *models.Company, err error) {
	defer mw.log("CreateCompany", time.Now(), input, output, &err)
	output, err = mw.next.CreateCompany(ctx, input)
	return
}

// GetAllCompanies returns all Companies that match the filters
func (mw logMiddleware) GetAllCompanies(ctx context.Context, input models.CompanyFilters) (output []*models.Company, err error) {
	defer mw.log("GetAllCompanies", time.Now(), input, output, &err)
	output, err = mw.next.GetAllCompanies(ctx, input)
	return
}

// UpdateCompany updates a Company
func (mw logMiddleware) UpdateCompany(ctx context.Context, input *models.Company) (output *models.Company, err error) {
	defer mw.log("UpdateCompany", time.Now(), input, output, &err)
	output, err = mw.next.UpdateCompany(ctx, input)
	return
}

// DeleteCompany deletes a Company by ID
func (mw logMiddleware) DeleteCompany(ctx context.Context, input uint64) (err error) {
	defer mw.log("DeleteCompany", time.Now(), input, nil, &err)
	err = mw.next.DeleteCompany(ctx, input)
	return
}

/* --------------- Industry --------------- */

// CreateIndustry creates a new Industry
func (mw logMiddleware) CreateIndustry(ctx context.Context, input *models.Industry) (output *models.Industry, err error) {
	defer mw.log("CreateIndustry", time.Now(), input, output, &err)
	output, err = mw.next.CreateIndustry(ctx, input)
	return
}

// GetAllIndustries returns all Industries
func (mw logMiddleware) GetAllIndustries(ctx context.Context, input models.IndustryFilters) (output []*models.Industry, err error) {
	defer mw.log("GetAllIndustries", time.Now(), input, output, &err)
	output, err = mw.next.GetAllIndustries(ctx, input)
	return
}

// DeleteIndustry deletes a Industry by ID
func (mw logMiddleware) DeleteIndustry(ctx context.Context, input uint64) (err error) {
	defer mw.log("DeleteIndustry", time.Now(), input, nil, &err)
	err = mw.next.DeleteIndustry(ctx, input)
	return
}

/* --------------- Job Function --------------- */

// CreateJobFunction creates a new JobFunction
func (mw logMiddleware) CreateJobFunction(ctx context.Context, input *models.JobFunction) (output *models.JobFunction, err error) {
	defer mw.log("CreateJobFunction", time.Now(), input, output, &err)
	output, err = mw.next.CreateJobFunction(ctx, input)
	return
}

// GetAllJobFunctions returns all JobFunctions
func (mw logMiddleware) GetAllJobFunctions(ctx context.Context, input models.JobFunctionFilters) (output []*models.JobFunction, err error) {
	defer mw.log("GetAllJobFunctions", time.Now(), input, output, &err)
	output, err = mw.next.GetAllJobFunctions(ctx, input)
	return
}

// DeleteJobFunction deletes a JobFunction by ID
func (mw logMiddleware) DeleteJobFunction(ctx context.Context, input uint64) (err error) {
	defer mw.log("DeleteJobFunction", time.Now(), input, nil, &err)
	err = mw.next.DeleteJobFunction(ctx, input)
	return
}

/* --------------- Key Person --------------- */

// CreateKeyPerson creates a new KeyPerson
func (mw logMiddleware) CreateKeyPerson(ctx context.Context, input *models.KeyPerson) (output *models.KeyPerson, err error) {
	defer mw.log("CreateKeyPerson", time.Now(), input, output, &err)
	output, err = mw.next.CreateKeyPerson(ctx, input)
	return
}

// BulkCreateKeyPerson creates multiple KeyPersons
func (mw logMiddleware) BulkCreateKeyPerson(ctx context.Context, input []*models.KeyPerson) (output []*models.KeyPerson, err error) {
	defer mw.log("BulkCreateKeyPerson", time.Now(), input, output, &err)
	output, err = mw.next.BulkCreateKeyPerson(ctx, input)
	return
}

// GetAllKeyPersons returns all KeyPersons that match the filters
func (mw logMiddleware) GetAllKeyPersons(ctx context.Context, input models.KeyPersonFilters) (output []*models.KeyPerson, err error) {
	defer mw.log("GetAllKeyPersons", time.Now(), input, output, &err)
	output, err = mw.next.GetAllKeyPersons(ctx, input)
	return
}

// GetKeyPersonByID finds and returns a KeyPerson by ID
func (mw logMiddleware) GetKeyPersonByID(ctx context.Context, input uint64) (output *models.KeyPerson, err error) {
	defer mw.log("GetKeyPersonByID", time.Now(), input, output, &err)
	output, err = mw.next.GetKeyPersonByID(ctx, input)
	return
}

// UpdateKeyPerson updates a KeyPerson
func (mw logMiddleware) UpdateKeyPerson(ctx context.Context, input *models.KeyPerson) (output *models.KeyPerson, err error) {
	defer mw.log("UpdateKeyPerson", time.Now(), input, output, &err)
	output, err = mw.next.UpdateKeyPerson(ctx, input)
	return
}

// DeleteKeyPerson deletes a KeyPerson by ID
func (mw logMiddleware) DeleteKeyPerson(ctx context.Context, input uint64) (err error) {
	defer mw.log("DeleteKeyPerson", time.Now(), input, nil, &err)
	err = mw.next.DeleteKeyPerson(ctx, input)
	return
}

/* --------------- Job Platform --------------- */

// CreateJobPlatform creates a new JobPlatform
func (mw logMiddleware) CreateJobPlatform(ctx context.Context, input *models.JobPlatform) (output *models.JobPlatform, err error) {
	defer mw.log("CreateJobPlatform", time.Now(), input, output, &err)
	output, err = mw.next.CreateJobPlatform(ctx, input)
	return
}

// GetAllJobPlatforms returns all JobPlatforms
func (mw logMiddleware) GetAllJobPlatforms(ctx context.Context, input models.JobPlatformFilters) (output []*models.JobPlatform, err error) {
	defer mw.log("GetAllJobPlatforms", time.Now(), input, output, &err)
	output, err = mw.next.GetAllJobPlatforms(ctx, input)
	return
}

// DeleteJobPlatform deletes a JobPlatform by ID
func (mw logMiddleware) DeleteJobPlatform(ctx context.Context, input uint64) (err error) {
	defer mw.log("DeleteJobPlatform", time.Now(), input, nil, &err)
	err = mw.next.DeleteJobPlatform(ctx, input)
	return
}
