package database

import (
	"context"
	"fmt"
	"strings"
	"time"

	pg "github.com/go-pg/pg/v10"
	"github.com/pkg/errors"

	"in-backend/services/joblisting/interfaces"
	"in-backend/services/joblisting/models"
)

// Repository implements the joblisting Repository interface
type repository struct {
	DB *pg.DB
}

// NewRepository declares a new Repository that implements joblisting Repository
func NewRepository(db *pg.DB) interfaces.Repository {
	return &repository{
		DB: db,
	}
}

/* --------------- Job Post --------------- */

// CreateJobPost creates a new JobPost
func (r *repository) CreateJobPost(ctx context.Context, m *models.JobPost) (*models.JobPost, error) {
	if m == nil {
		return nil, errors.New("Input parameter job post is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(m).
		Returning("*").
		Insert()
	if err != nil {
		err = errors.Wrapf(err, "Failed to insert job post %v", m)
		return nil, err
	}

	return m, nil
}

// BulkCreateJobPost creates multiple JobPosts
func (r *repository) BulkCreateJobPost(ctx context.Context, m []*models.JobPost) ([]*models.JobPost, error) {
	empty := []*models.JobPost{}
	if len(m) == 0 {
		return empty, errors.New("Input parameter job posts is empty")
	}

	// create all foreign relations first
	var companies []*models.Company
	var functions []*models.JobFunction
	var industries []*models.Industry
	var jobPlatforms []*models.JobPlatform
	var hrContacts []*models.KeyPerson
	var hiringManagers []*models.KeyPerson
	for _, jp := range m {
		companies = append(companies, jp.Company)
		functions = append(functions, jp.Function)
		industries = append(industries, jp.Industry)
		jobPlatforms = append(jobPlatforms, jp.JobPlatform)
		hrContacts = append(hrContacts, jp.HRContact)
		hiringManagers = append(hiringManagers, jp.HiringManager)
	}

	tx, err := r.DB.BeginContext(ctx)
	defer tx.Close()

	_, err = tx.Model(&companies).OnConflict("(name) DO UPDATE").Set("name = EXCLUDED.name").Insert()
	if err != nil {
		tx.Rollback()
		err = errors.Wrapf(err, "Failed to insert companies %v", companies)
		return empty, err
	}

	_, err = tx.Model(&functions).OnConflict("(name) DO UPDATE").Set("name = EXCLUDED.name").Insert()
	if err != nil {
		tx.Rollback()
		err = errors.Wrapf(err, "Failed to insert functions %v", functions)
		return empty, err
	}

	_, err = tx.Model(&industries).OnConflict("(name) DO UPDATE").Set("name = EXCLUDED.name").Insert()
	if err != nil {
		tx.Rollback()
		err = errors.Wrapf(err, "Failed to insert industries %v", industries)
		return empty, err
	}

	for i, co := range companies {
		hrContacts[i].CompanyID = co.ID
		hiringManagers[i].CompanyID = co.ID
	}
	_, err = tx.Model(&hrContacts).OnConflict("(company_id, name) DO UPDATE").Set("job_title = EXCLUDED.job_title").Insert()
	if err != nil {
		tx.Rollback()
		err = errors.Wrapf(err, "Failed to insert hr contacts %v", hrContacts)
		return empty, err
	}

	_, err = tx.Model(&hiringManagers).OnConflict("(company_id, name) DO UPDATE").Set("job_title = EXCLUDED.job_title").Insert()
	if err != nil {
		tx.Rollback()
		err = errors.Wrapf(err, "Failed to insert hiring managers %v", hiringManagers)
		return empty, err
	}

	for i := range m {
		m[i].CompanyID = companies[i].ID
		m[i].FunctionID = functions[i].ID
		m[i].IndustryID = industries[i].ID
		m[i].HRContactID = hrContacts[i].ID
		m[i].HiringManagerID = hiringManagers[i].ID
	}
	_, err = tx.Model(&m).
		Returning("*").
		Insert()
	if err != nil {
		tx.Rollback()
		err = errors.Wrapf(err, "Failed to insert job posts %v", m)
		return empty, err
	}

	if err := tx.Commit(); err != nil {
		return empty, err
	}

	return m, nil
}

// GetAllJobPosts returns all JobPosts that match the filters
func (r *repository) GetAllJobPosts(ctx context.Context, f models.JobPostFilters) ([]*models.JobPost, error) {
	var m []*models.JobPost
	q := r.DB.WithContext(ctx).Model(&m)
	if len(f.ID) > 0 {
		q = q.Where("jp.id in (?)", pg.In(f.ID))
	}
	if len(f.CompanyID) > 0 {
		q = q.Where("jp.company_id in (?)", pg.In(f.CompanyID))
	}
	if len(f.HRContactID) > 0 {
		q = q.Where("jp.hiring_contact_id in (?)", pg.In(f.HRContactID))
	}
	if len(f.HiringManagerID) > 0 {
		q = q.Where("jp.hiring_manager_id in (?)", pg.In(f.HiringManagerID))
	}
	if len(f.JobPlatformID) > 0 {
		q = q.Where("jp.job_platform_id in (?)", f.JobPlatformID)
	}
	if len(f.SkillID) > 0 {
		q = q.Where("jp.skillID @>", pg.Array(f.SkillID))
	}
	if len(f.Title) > 0 {
		q = q.Where("lower(jp.title) like ?", "%"+strings.ToLower(f.Title)+"%")
	}
	if len(f.SeniorityLevel) > 0 {
		q = q.Where("jp.education_level in (?)", pg.In(f.SeniorityLevel))
	}
	if f.MinYearsExperience > 0 {
		q = q.Where("jp.years_experience >= ?", f.MinYearsExperience)
	}
	if f.MaxYearsExperience > 0 {
		q = q.Where("jp.years_experience <= ?", f.MaxYearsExperience)
	}
	if len(f.EmploymentType) > 0 {
		q = q.Where("jp.employment_type in (?)", pg.In(f.EmploymentType))
	}
	if len(f.FunctionID) > 0 {
		q = q.Where("jp.function_id in (?)", pg.In(f.FunctionID))
	}
	if len(f.IndustryID) > 0 {
		q = q.Where("jp.industry_id in (?)", pg.In(f.IndustryID))
	}
	if len(f.SeniorityLevel) > 0 {
		q = q.Where("jp.education_level in (?)", pg.In(f.SeniorityLevel))
	}
	if f.Remote == true {
		q = q.Where("jp.remote = true")
	}
	if f.Salary > 0 {
		q = q.Where("jp.min_salary >= ?", f.Salary).
			Where("jp.max_salary <= ?", f.Salary)
	}
	if f.UpdatedAt != nil {
		q = q.Where("jp.updated_at >= ?", f.UpdatedAt.Format("2006-01-02"))
	}
	if f.ExpireAt != nil {
		q = q.Where("jp.expire_at >= ?", f.ExpireAt.Format("2006-01-02"))
	}
	err := q.Where("expire_at >= ?", time.Now().Format("2006-01-02")).
		Relation(relCompany).
		Relation(relFunction).
		Relation(relIndustry).
		Returning("*").
		Order("jp.updated_at desc").
		Select()
	return m, err
}

// GetJobPostByID finds and returns a JobPost by ID
func (r *repository) GetJobPostByID(ctx context.Context, id uint64) (*models.JobPost, error) {
	m := models.JobPost{ID: id}
	err := r.DB.WithContext(ctx).Model(&m).
		Where("id = ?", id).
		Relation(relCompany).
		Relation(relFunction).
		Relation(relIndustry).
		Relation(relHiringManager).
		Relation(relHRContact).
		Returning("*").
		First()
	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return &m, err
}

// UpdateJobPost updates a JobPost
func (r *repository) UpdateJobPost(ctx context.Context, m *models.JobPost) (*models.JobPost, error) {
	if m == nil {
		return nil, errors.New("Job post is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(m).WherePK().
		Relation(relCompany).
		Relation(relFunction).
		Relation(relIndustry).
		Relation(relHiringManager).
		Relation(relHRContact).
		Returning("*").
		Update()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Cannot update job post with id %v", m.ID))
	}

	return m, nil
}

// DeleteJobPost deletes a JobPost by ID
func (r *repository) DeleteJobPost(ctx context.Context, id uint64) error {
	m := &models.JobPost{ID: id}
	_, err := r.DB.WithContext(ctx).Model(m).WherePK().Delete()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Cannot delete job post with id %v", id))
	}
	return nil
}

/* --------------- Company --------------- */

// CreateCompany creates a new Company
func (r *repository) CreateCompany(ctx context.Context, m *models.Company) (*models.Company, error) {
	if m == nil {
		return nil, errors.New("Input parameter company is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(m).
		Returning("*").
		Insert()
	if err != nil {
		err = errors.Wrapf(err, "Failed to insert company %v", m)
		return nil, err
	}

	return m, nil
}

// GetAllCompanies returns all Companies that match the filters
func (r *repository) GetAllCompanies(ctx context.Context, f models.CompanyFilters) ([]*models.Company, error) {
	var m []*models.Company
	q := r.DB.WithContext(ctx).Model(&m)
	if len(f.ID) > 0 {
		q = q.Where("co.id in (?)", pg.In(f.ID))
	}
	if len(f.Name) > 0 {
		q = q.Where("lower(co.name) like ?", "%"+strings.ToLower(f.Name)+"%")
	}
	err := q.Relation(relIndustries).
		Relation(relJobPosts).
		Relation(relKeyPersons).
		Returning("*").
		Select()
	return m, err
}

// UpdateCompany updates a Company
func (r *repository) UpdateCompany(ctx context.Context, m *models.Company) (*models.Company, error) {
	if m == nil {
		return nil, errors.New("Company is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(m).WherePK().
		Relation(relIndustries).
		Relation(relJobPosts).
		Relation(relKeyPersons).
		Returning("*").
		Update()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Cannot update company with id %v", m.ID))
	}

	return m, nil
}

// DeleteCompany deletes a Company by ID
func (r *repository) DeleteCompany(ctx context.Context, id uint64) error {
	m := &models.Company{ID: id}
	_, err := r.DB.WithContext(ctx).Model(m).WherePK().Delete()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Cannot delete company with id %v", id))
	}
	return nil
}

/* --------------- Industry --------------- */

// CreateIndustry creates a new Industry
func (r *repository) CreateIndustry(ctx context.Context, m *models.Industry) (*models.Industry, error) {
	if m == nil {
		return nil, errors.New("Input parameter industry is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(m).
		Returning("*").
		Insert()
	if err != nil {
		err = errors.Wrapf(err, "Failed to insert industry %v", m)
		return nil, err
	}

	return m, nil
}

// GetAllIndustries returns all Industries
func (r *repository) GetAllIndustries(ctx context.Context, f models.IndustryFilters) ([]*models.Industry, error) {
	var m []*models.Industry
	q := r.DB.WithContext(ctx).Model(&m)
	if len(f.ID) > 0 {
		q = q.Where("id.id in (?)", pg.In(f.ID))
	}
	if len(f.Name) > 0 {
		q = q.Where("lower(id.name) like ?", "%"+strings.ToLower(f.Name)+"%")
	}
	err := q.Relation(relCompanies).
		Relation(relJobPosts).
		Returning("*").
		Select()
	return m, err
}

// DeleteIndustry deletes a Industry by ID
func (r *repository) DeleteIndustry(ctx context.Context, id uint64) error {
	m := &models.Industry{ID: id}
	_, err := r.DB.WithContext(ctx).Model(m).WherePK().Delete()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Cannot delete industry with id %v", id))
	}
	return nil
}

/* --------------- Job Function --------------- */

// CreateJobFunction creates a new JobFunction
func (r *repository) CreateJobFunction(ctx context.Context, m *models.JobFunction) (*models.JobFunction, error) {
	if m == nil {
		return nil, errors.New("Input parameter job function is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(m).
		Returning("*").
		Insert()
	if err != nil {
		err = errors.Wrapf(err, "Failed to insert job function %v", m)
		return nil, err
	}

	return m, nil
}

// GetAllJobFunctions returns all JobFunctions
func (r *repository) GetAllJobFunctions(ctx context.Context, f models.JobFunctionFilters) ([]*models.JobFunction, error) {
	var m []*models.JobFunction
	q := r.DB.WithContext(ctx).Model(&m)
	if len(f.ID) > 0 {
		q = q.Where("jf.id in (?)", pg.In(f.ID))
	}
	if len(f.Name) > 0 {
		q = q.Where("lower(jf.name) like ?", "%"+strings.ToLower(f.Name)+"%")
	}
	err := q.Returning("*").Select()
	return m, err
}

// DeleteJobFunction deletes a JobFunction by ID
func (r *repository) DeleteJobFunction(ctx context.Context, id uint64) error {
	m := &models.JobFunction{ID: id}
	_, err := r.DB.WithContext(ctx).Model(m).WherePK().Delete()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Cannot delete job function with id %v", id))
	}
	return nil
}

/* --------------- Key Person --------------- */

// CreateKeyPerson creates a new KeyPerson
func (r *repository) CreateKeyPerson(ctx context.Context, m *models.KeyPerson) (*models.KeyPerson, error) {
	if m == nil {
		return nil, errors.New("Input parameter key person is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(m).
		Returning("*").
		Insert()
	if err != nil {
		err = errors.Wrapf(err, "Failed to insert key person %v", m)
		return nil, err
	}

	return m, nil
}

// BulkCreateKeyPerson creates multiple KeyPersons
func (r *repository) BulkCreateKeyPerson(ctx context.Context, m []*models.KeyPerson) ([]*models.KeyPerson, error) {
	empty := []*models.KeyPerson{}
	if len(m) == 0 {
		return empty, errors.New("Input parameter key persons is empty")
	}

	// create all foreign relations first
	var companies []*models.Company
	for _, jp := range m {
		companies = append(companies, jp.Company)
	}

	tx, err := r.DB.BeginContext(ctx)
	defer tx.Close()

	_, err = tx.Model(&companies).OnConflict("(name) DO UPDATE").Set("name = EXCLUDED.name").Insert()
	if err != nil {
		tx.Rollback()
		err = errors.Wrapf(err, "Failed to insert companies %v", companies)
		return empty, err
	}

	for i := range m {
		m[i].CompanyID = companies[i].ID
	}
	_, err = tx.Model(&m).
		Returning("*").
		Insert()
	if err != nil {
		tx.Rollback()
		err = errors.Wrapf(err, "Failed to insert key persons %v", m)
		return empty, err
	}

	if err := tx.Commit(); err != nil {
		return empty, err
	}

	return m, nil
}

// GetAllKeyPersons returns all KeyPersons that match the filters
func (r *repository) GetAllKeyPersons(ctx context.Context, f models.KeyPersonFilters) ([]*models.KeyPerson, error) {
	var m []*models.KeyPerson
	q := r.DB.WithContext(ctx).Model(&m)
	if len(f.ID) > 0 {
		q = q.Where("kp.id in (?)", pg.In(f.ID))
	}
	if len(f.CompanyID) > 0 {
		q = q.Where("kp.company_id in (?)", pg.In(f.CompanyID))
	}
	if len(f.Name) > 0 {
		q = q.Where("lower(kp.name) like ?", "%"+strings.ToLower(f.Name)+"%")
	}
	if len(f.ContactNumber) > 0 {
		q = q.Where("lower(kp.contact_number) like ?", "%"+strings.ToLower(f.ContactNumber)+"%")
	}
	if len(f.Email) > 0 {
		q = q.Where("lower(kp.email) like ?", "%"+strings.ToLower(f.Email)+"%")
	}
	if len(f.JobTitle) > 0 {
		q = q.Where("lower(kp.job_title) like ?", "%"+strings.ToLower(f.JobTitle)+"%")
	}
	err := q.Returning("*").Select()
	return m, err
}

// GetKeyPersonByID finds and returns a KeyPerson by ID
func (r *repository) GetKeyPersonByID(ctx context.Context, id uint64) (*models.KeyPerson, error) {
	m := models.KeyPerson{ID: id}
	err := r.DB.WithContext(ctx).Model(&m).
		Where("id = ?", id).
		Relation(relCompany).
		Returning("*").
		First()
	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return &m, err
}

// UpdateKeyPerson updates a KeyPerson
func (r *repository) UpdateKeyPerson(ctx context.Context, m *models.KeyPerson) (*models.KeyPerson, error) {
	if m == nil {
		return nil, errors.New("Key person is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(m).WherePK().
		Relation(relCompany).
		Returning("*").
		Update()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Cannot update key person with id %v", m.ID))
	}

	return m, nil
}

// DeleteKeyPerson deletes a KeyPerson by ID
func (r *repository) DeleteKeyPerson(ctx context.Context, id uint64) error {
	m := &models.KeyPerson{ID: id}
	_, err := r.DB.WithContext(ctx).Model(m).WherePK().Delete()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Cannot delete key person with id %v", id))
	}
	return nil
}

/* --------------- Job Platform --------------- */

// CreateJobPlatform creates a new JobPlatform
func (r *repository) CreateJobPlatform(ctx context.Context, m *models.JobPlatform) (*models.JobPlatform, error) {
	if m == nil {
		return nil, errors.New("Input parameter job platform is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(m).
		Returning("*").
		Insert()
	if err != nil {
		err = errors.Wrapf(err, "Failed to insert job platform %v", m)
		return nil, err
	}

	return m, nil
}

// GetAllJobPlatforms returns all JobPlatforms
func (r *repository) GetAllJobPlatforms(ctx context.Context, f models.JobPlatformFilters) ([]*models.JobPlatform, error) {
	var m []*models.JobPlatform
	q := r.DB.WithContext(ctx).Model(&m)
	if len(f.ID) > 0 {
		q = q.Where("id.id in (?)", pg.In(f.ID))
	}
	if len(f.Name) > 0 {
		q = q.Where("lower(id.name) like ?", "%"+strings.ToLower(f.Name)+"%")
	}
	err := q.Relation(relJobPosts).
		Returning("*").
		Select()
	return m, err
}

// DeleteJobPlatform deletes a JobPlatform by ID
func (r *repository) DeleteJobPlatform(ctx context.Context, id uint64) error {
	m := &models.JobPlatform{ID: id}
	_, err := r.DB.WithContext(ctx).Model(m).WherePK().Delete()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Cannot delete job platform with id %v", id))
	}
	return nil
}
