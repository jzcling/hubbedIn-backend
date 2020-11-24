package database

import (
	"context"
	"fmt"

	pg "github.com/go-pg/pg/v10"
	"github.com/pkg/errors"

	"in-backend/services/profile"
	"in-backend/services/profile/models"
)

const (
	REL_CANDIDATE_SKILL                string = "Skills"
	REL_CANDIDATE_ACADEMIC             string = "Academics"
	REL_CANDIDATE_ACADEMIC_INSTITUTION string = "Academics.Institution"
	REL_CANDIDATE_ACADEMIC_COURSE      string = "Academics.Course"
	REL_CANDIDATE_JOB                  string = "Jobs"
	REL_CANDIDATE_JOB_COMPANY          string = "Jobs.Company"
	REL_CANDIDATE_JOB_DEPARTMENT       string = "Jobs.Department"

	FILTER_ID string = "id = ?"
)

// Repository implements the profile Repository interface
type repository struct {
	DB *pg.DB
}

// NewRepository declares a new Repository that implements profile Repository
func NewRepository(db *pg.DB) profile.Repository {
	return &repository{
		DB: db,
	}
}

/* --------------- Candidate --------------- */

// CreateCandidate creates a new Candidate
func (r *repository) CreateCandidate(ctx context.Context, c *models.Candidate) (*models.Candidate, error) {
	if c == nil {
		return nil, errors.New("Input parameter candidate is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(c).
		Relation(REL_CANDIDATE_SKILL).
		Relation(REL_CANDIDATE_ACADEMIC).Relation(REL_CANDIDATE_ACADEMIC_INSTITUTION).Relation(REL_CANDIDATE_ACADEMIC_COURSE).
		Relation(REL_CANDIDATE_JOB).Relation(REL_CANDIDATE_JOB_COMPANY).Relation(REL_CANDIDATE_JOB_DEPARTMENT).
		Returning("*").
		Insert()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to insert candidate %v", c)
	}

	return c, nil
}

// GetAllCandidates returns all Candidates
func (r *repository) GetAllCandidates(ctx context.Context) ([]*models.Candidate, error) {
	var c []*models.Candidate
	err := r.DB.WithContext(ctx).Model(&c).
		Relation(REL_CANDIDATE_SKILL).
		Relation(REL_CANDIDATE_ACADEMIC).Relation(REL_CANDIDATE_ACADEMIC_INSTITUTION).Relation(REL_CANDIDATE_ACADEMIC_COURSE).
		Relation(REL_CANDIDATE_JOB).Relation(REL_CANDIDATE_JOB_COMPANY).Relation(REL_CANDIDATE_JOB_DEPARTMENT).
		Returning("*").
		Select()
	return c, err
}

// GetCandidateByID returns a Candidate by ID
func (r *repository) GetCandidateByID(ctx context.Context, id uint64) (*models.Candidate, error) {
	c := models.Candidate{ID: id}
	err := r.DB.WithContext(ctx).Model(&c).
		Where(FILTER_ID, id).
		Relation(REL_CANDIDATE_SKILL).
		Relation(REL_CANDIDATE_ACADEMIC).Relation(REL_CANDIDATE_ACADEMIC_INSTITUTION).Relation(REL_CANDIDATE_ACADEMIC_COURSE).
		Relation(REL_CANDIDATE_JOB).Relation(REL_CANDIDATE_JOB_COMPANY).Relation(REL_CANDIDATE_JOB_DEPARTMENT).
		Returning("*").
		Select()
	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return &c, err
}

// UpdateCandidate updates a Candidate
func (r *repository) UpdateCandidate(ctx context.Context, c *models.Candidate) (*models.Candidate, error) {
	if c == nil {
		return nil, errors.New("Candidate is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(c).WherePK().
		Relation(REL_CANDIDATE_SKILL).
		Relation(REL_CANDIDATE_ACADEMIC).Relation(REL_CANDIDATE_ACADEMIC_INSTITUTION).Relation(REL_CANDIDATE_ACADEMIC_COURSE).
		Relation(REL_CANDIDATE_JOB).Relation(REL_CANDIDATE_JOB_COMPANY).Relation(REL_CANDIDATE_JOB_DEPARTMENT).
		Returning("*").
		Update()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Cannot update candidate with id %v", c.ID))
	}

	return c, nil
}

// DeleteCandidate deletes a Candidate by ID
func (r *repository) DeleteCandidate(ctx context.Context, id uint64) error {
	c := &models.Candidate{ID: id}
	_, err := r.DB.WithContext(ctx).Model(c).WherePK().Delete()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Cannot delete candidate with id %v", id))
	}
	return nil
}

/* --------------- Skill --------------- */

// CreateSkill creates a new Skill
func (r *repository) CreateSkill(ctx context.Context, s *models.Skill) (*models.Skill, error) {
	if s == nil {
		return nil, errors.New("Input parameter skill is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(s).Returning("*").Insert()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to insert skill %v", s)
	}

	return s, nil
}

// GetSkill returns a Skill by ID
func (r *repository) GetSkill(ctx context.Context, id uint64) (*models.Skill, error) {
	s := models.Skill{ID: id}
	err := r.DB.WithContext(ctx).Model(&s).Where(FILTER_ID, id).Select()
	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return &s, err
}

// GetAllSkills returns all Skills
func (r *repository) GetAllSkills(ctx context.Context) ([]*models.Skill, error) {
	var s []*models.Skill
	err := r.DB.WithContext(ctx).Model(&s).Select()
	return s, err
}

/* --------------- User Skill --------------- */

// CreateUserSkill creates a new UserSkill
func (r *repository) CreateUserSkill(ctx context.Context, us *models.UserSkill) (*models.UserSkill, error) {
	if us == nil {
		return nil, errors.New("Input parameter user skill is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(us).Returning("*").Insert()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to insert user skill %v", us)
	}

	return us, nil
}

// DeleteUserSkill deletes a UserSkill by ID
func (r *repository) DeleteUserSkill(ctx context.Context, id uint64) error {
	us := &models.UserSkill{ID: id}
	_, err := r.DB.WithContext(ctx).Model(us).WherePK().Delete()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Cannot delete user skill with id %v", id))
	}
	return nil
}

/* --------------- Institution --------------- */

// CreateInstitution creates a new Institution
func (r *repository) CreateInstitution(ctx context.Context, i *models.Institution) (*models.Institution, error) {
	if i == nil {
		return nil, errors.New("Input parameter institution is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(i).Returning("*").Insert()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to insert institution %v", i)
	}

	return i, nil
}

// GetInstitution returns a Institution by ID
func (r *repository) GetInstitution(ctx context.Context, id uint64) (*models.Institution, error) {
	i := models.Institution{ID: id}
	err := r.DB.WithContext(ctx).Model(&i).Where(FILTER_ID, id).Select()
	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return &i, err
}

// GetAllInstitutions returns all Institutions
func (r *repository) GetAllInstitutions(ctx context.Context) ([]*models.Institution, error) {
	var i []*models.Institution
	err := r.DB.WithContext(ctx).Model(&i).Select()
	return i, err
}

/* --------------- Course --------------- */

// CreateCourse creates a new Course
func (r *repository) CreateCourse(ctx context.Context, c *models.Course) (*models.Course, error) {
	if c == nil {
		return nil, errors.New("Input parameter course is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(c).Returning("*").Insert()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to insert course %v", c)
	}

	return c, nil
}

// GetCourse returns a Course by ID
func (r *repository) GetCourse(ctx context.Context, id uint64) (*models.Course, error) {
	c := models.Course{ID: id}
	err := r.DB.WithContext(ctx).Model(&c).Where(FILTER_ID, id).Select()
	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return &c, err
}

// GetAllCourses returns all Courses
func (r *repository) GetAllCourses(ctx context.Context) ([]*models.Course, error) {
	var c []*models.Course
	err := r.DB.WithContext(ctx).Model(&c).Select()
	return c, err
}

/* --------------- Academic History --------------- */

// CreateAcademicHistory creates a new AcademicHistory
func (r *repository) CreateAcademicHistory(ctx context.Context, a *models.AcademicHistory) (*models.AcademicHistory, error) {
	if a == nil {
		return nil, errors.New("Input parameter academic history is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(a).
		Relation("Institution").
		Relation("Course").
		Returning("*").
		Insert()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to insert academic history %v", a)
	}

	return a, nil
}

// GetAcademicHistory returns a AcademicHistory by ID
func (r *repository) GetAcademicHistory(ctx context.Context, id uint64) (*models.AcademicHistory, error) {
	a := models.AcademicHistory{ID: id}
	err := r.DB.WithContext(ctx).Model(&a).
		Where(FILTER_ID, id).
		Relation("Institution").
		Relation("Course").
		Select()
	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return &a, err
}

// UpdateAcademicHistory updates a AcademicHistory
func (r *repository) UpdateAcademicHistory(ctx context.Context, a *models.AcademicHistory) (*models.AcademicHistory, error) {
	if a == nil {
		return nil, errors.New("AcademicHistory is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(a).WherePK().
		Relation("Institution").
		Relation("Course").
		Returning("*").
		Update()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Cannot update academic history with id %v", a.ID))
	}

	return a, nil
}

// DeleteAcademicHistory deletes a AcademicHistory by ID
func (r *repository) DeleteAcademicHistory(ctx context.Context, id uint64) error {
	a := &models.AcademicHistory{ID: id}
	_, err := r.DB.WithContext(ctx).Model(a).WherePK().Delete()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Cannot delete academic history with id %v", id))
	}
	return nil
}

/* --------------- Company --------------- */

// CreateCompany creates a new Company
func (r *repository) CreateCompany(ctx context.Context, c *models.Company) (*models.Company, error) {
	if c == nil {
		return nil, errors.New("Input parameter company is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(c).Returning("*").Insert()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to insert company %v", c)
	}

	return c, nil
}

// GetCompany returns a Company by ID
func (r *repository) GetCompany(ctx context.Context, id uint64) (*models.Company, error) {
	c := models.Company{ID: id}
	err := r.DB.WithContext(ctx).Model(&c).Where(FILTER_ID, id).Select()
	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return &c, err
}

// GetAllCompanies returns all Companies
func (r *repository) GetAllCompanies(ctx context.Context) ([]*models.Company, error) {
	var c []*models.Company
	err := r.DB.WithContext(ctx).Model(&c).Select()
	return c, err
}

/* --------------- Department --------------- */

// CreateDepartment creates a new Department
func (r *repository) CreateDepartment(ctx context.Context, d *models.Department) (*models.Department, error) {
	if d == nil {
		return nil, errors.New("Input parameter department is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(d).Returning("*").Insert()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to insert department %v", d)
	}

	return d, nil
}

// GetDepartment returns a Department by ID
func (r *repository) GetDepartment(ctx context.Context, id uint64) (*models.Department, error) {
	d := models.Department{ID: id}
	err := r.DB.WithContext(ctx).Model(&d).Where(FILTER_ID, id).Select()
	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return &d, err
}

// GetAllDepartments returns all Departments
func (r *repository) GetAllDepartments(ctx context.Context) ([]*models.Department, error) {
	var d []*models.Department
	err := r.DB.WithContext(ctx).Model(&d).Select()
	return d, err
}

/* --------------- Job History --------------- */

// CreateJobHistory creates a new JobHistory
func (r *repository) CreateJobHistory(ctx context.Context, j *models.JobHistory) (*models.JobHistory, error) {
	if j == nil {
		return nil, errors.New("Input parameter job history is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(j).
		Relation("Company").
		Relation("Department").
		Returning("*").
		Insert()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to insert job history %v", j)
	}

	return j, nil
}

// GetJobHistory returns a JobHistory by ID
func (r *repository) GetJobHistory(ctx context.Context, id uint64) (*models.JobHistory, error) {
	j := models.JobHistory{ID: id}
	err := r.DB.WithContext(ctx).Model(&j).Where(FILTER_ID, id).
		Relation("Company").
		Relation("Department").
		Select()
	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return &j, err
}

// UpdateJobHistory updates a JobHistory
func (r *repository) UpdateJobHistory(ctx context.Context, j *models.JobHistory) (*models.JobHistory, error) {
	if j == nil {
		return nil, errors.New("JobHistory is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(j).WherePK().
		Relation("Company").
		Relation("Department").
		Returning("*").
		Update()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Cannot update job history with id %v", j.ID))
	}

	return j, nil
}

// DeleteJobHistory deletes a JobHistory by ID
func (r *repository) DeleteJobHistory(ctx context.Context, id uint64) error {
	j := &models.JobHistory{ID: id}
	_, err := r.DB.WithContext(ctx).Model(j).WherePK().Delete()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Cannot delete job history with id %v", id))
	}
	return nil
}
