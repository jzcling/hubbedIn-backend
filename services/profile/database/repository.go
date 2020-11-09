package database

import (
	"context"
	"fmt"

	pg "github.com/go-pg/pg/v10"
	"github.com/pkg/errors"

	"in-backend/services/profile"
	"in-backend/services/profile/models"
)

type repository struct {
	DB *pg.DB
}

// NewRepository declares a new profile repository
func NewRepository(db *pg.DB) profile.Repository {
	return &repository{
		DB: db,
	}
}

/* --------------- Candidate --------------- */

func (r repository) CreateCandidate(ctx context.Context, c *models.Candidate) (*models.Candidate, error) {
	if c == nil {
		return nil, errors.New("Input parameter candidate is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(c).Returning("*").Insert()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to insert candidate %v", c)
	}

	return c, nil
}

func (r repository) GetAllCandidates(ctx context.Context) ([]*models.Candidate, error) {
	var c []*models.Candidate
	err := r.DB.WithContext(ctx).Model(&c).Select()
	return c, err
}

func (r repository) GetCandidateByID(ctx context.Context, id uint64) (*models.Candidate, error) {
	c := models.Candidate{ID: id}
	err := r.DB.WithContext(ctx).Model(&c).Where("id = ?", id).Select()
	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		return &c, nil
	}
	return &c, err
}

func (r repository) UpdateCandidate(ctx context.Context, c *models.Candidate) (*models.Candidate, error) {
	if c == nil {
		return nil, errors.New("Candidate is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(c).Returning("*").Update()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Cannot update candidate with id %v", c.ID))
	}

	return c, nil
}

func (r repository) DeleteCandidate(ctx context.Context, id uint64) error {
	c := &models.Candidate{ID: id}
	_, err := r.DB.WithContext(ctx).Model(c).WherePK().Delete()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Cannot delete candidate with id %v", id))
	}
	return nil
}

/* --------------- Skill --------------- */

func (r repository) CreateSkill(ctx context.Context, s *models.Skill) (*models.Skill, error) {
	if s == nil {
		return nil, errors.New("Input parameter skill is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(s).Returning("*").Insert()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to insert skill %v", s)
	}

	return s, nil
}

func (r repository) GetSkill(ctx context.Context, id uint64) (*models.Skill, error) {
	s := models.Skill{ID: id}
	err := r.DB.WithContext(ctx).Model(&s).Where("id = ?", id).Select()
	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		return &s, nil
	}
	return &s, err
}

func (r repository) GetAllSkills(ctx context.Context) ([]*models.Skill, error) {
	var s []*models.Skill
	err := r.DB.WithContext(ctx).Model(&s).Select()
	return s, err
}

/* --------------- Institution --------------- */

func (r repository) CreateInstitution(ctx context.Context, i *models.Institution) (*models.Institution, error) {
	if i == nil {
		return nil, errors.New("Input parameter institution is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(i).Returning("*").Insert()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to insert institution %v", i)
	}

	return i, nil
}

func (r repository) GetInstitution(ctx context.Context, id uint64) (*models.Institution, error) {
	i := models.Institution{ID: id}
	err := r.DB.WithContext(ctx).Model(&i).Where("id = ?", id).Select()
	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		return &i, nil
	}
	return &i, err
}

func (r repository) GetAllInstitutions(ctx context.Context) ([]*models.Institution, error) {
	var i []*models.Institution
	err := r.DB.WithContext(ctx).Model(&i).Select()
	return i, err
}

/* --------------- Course --------------- */

func (r repository) CreateCourse(ctx context.Context, c *models.Course) (*models.Course, error) {
	if c == nil {
		return nil, errors.New("Input parameter course is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(c).Returning("*").Insert()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to insert course %v", c)
	}

	return c, nil
}

func (r repository) GetCourse(ctx context.Context, id uint64) (*models.Course, error) {
	c := models.Course{ID: id}
	err := r.DB.WithContext(ctx).Model(&c).Where("id = ?", id).Select()
	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		return &c, nil
	}
	return &c, err
}

func (r repository) GetAllCourses(ctx context.Context) ([]*models.Course, error) {
	var c []*models.Course
	err := r.DB.WithContext(ctx).Model(&c).Select()
	return c, err
}

/* --------------- Academic History --------------- */

func (r repository) CreateAcademicHistory(ctx context.Context, a *models.AcademicHistory) (*models.AcademicHistory, error) {
	if a == nil {
		return nil, errors.New("Input parameter academic history is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(a).Returning("*").Insert()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to insert academic history %v", a)
	}

	return a, nil
}

func (r repository) GetAcademicHistory(ctx context.Context, id uint64) (*models.AcademicHistory, error) {
	a := models.AcademicHistory{ID: id}
	err := r.DB.WithContext(ctx).Model(&a).Where("id = ?", id).Select()
	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		return &a, nil
	}
	return &a, err
}

func (r repository) UpdateAcademicHistory(ctx context.Context, a *models.AcademicHistory) (*models.AcademicHistory, error) {
	if a == nil {
		return nil, errors.New("AcademicHistory is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(a).Returning("*").Update()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Cannot update academic history with id %v", a.ID))
	}

	return a, nil
}

func (r repository) DeleteAcademicHistory(ctx context.Context, id uint64) error {
	a := &models.AcademicHistory{ID: id}
	_, err := r.DB.WithContext(ctx).Model(a).WherePK().Delete()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Cannot delete academic history with id %v", id))
	}
	return nil
}

/* --------------- Company --------------- */

func (r repository) CreateCompany(ctx context.Context, c *models.Company) (*models.Company, error) {
	if c == nil {
		return nil, errors.New("Input parameter company is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(c).Returning("*").Insert()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to insert company %v", c)
	}

	return c, nil
}

func (r repository) GetCompany(ctx context.Context, id uint64) (*models.Company, error) {
	c := models.Company{ID: id}
	err := r.DB.WithContext(ctx).Model(&c).Where("id = ?", id).Select()
	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		return &c, nil
	}
	return &c, err
}

func (r repository) GetAllCompanies(ctx context.Context) ([]*models.Company, error) {
	var c []*models.Company
	err := r.DB.WithContext(ctx).Model(&c).Select()
	return c, err
}

/* --------------- Department --------------- */

func (r repository) CreateDepartment(ctx context.Context, d *models.Department) (*models.Department, error) {
	if d == nil {
		return nil, errors.New("Input parameter department is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(d).Returning("*").Insert()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to insert department %v", d)
	}

	return d, nil
}

func (r repository) GetDepartment(ctx context.Context, id uint64) (*models.Department, error) {
	d := models.Department{ID: id}
	err := r.DB.WithContext(ctx).Model(&d).Where("id = ?", id).Select()
	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		return &d, nil
	}
	return &d, err
}

func (r repository) GetAllDepartments(ctx context.Context) ([]*models.Department, error) {
	var d []*models.Department
	err := r.DB.WithContext(ctx).Model(&d).Select()
	return d, err
}

/* --------------- Job History --------------- */

func (r repository) CreateJobHistory(ctx context.Context, j *models.JobHistory) (*models.JobHistory, error) {
	if j == nil {
		return nil, errors.New("Input parameter job history is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(j).Returning("*").Insert()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to insert job history %v", j)
	}

	return j, nil
}

func (r repository) GetJobHistory(ctx context.Context, id uint64) (*models.JobHistory, error) {
	j := models.JobHistory{ID: id}
	err := r.DB.WithContext(ctx).Model(&j).Where("id = ?", id).Select()
	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		return &j, nil
	}
	return &j, err
}

func (r repository) UpdateJobHistory(ctx context.Context, j *models.JobHistory) (*models.JobHistory, error) {
	if j == nil {
		return nil, errors.New("JobHistory is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(j).Returning("*").Update()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Cannot update Job history with id %v", j.ID))
	}

	return j, nil
}

func (r repository) DeleteJobHistory(ctx context.Context, id uint64) error {
	j := &models.JobHistory{ID: id}
	_, err := r.DB.WithContext(ctx).Model(j).WherePK().Delete()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Cannot delete Job history with id %v", id))
	}
	return nil
}
