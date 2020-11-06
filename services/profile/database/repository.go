package database

import (
	"context"
	"fmt"

	pg "github.com/go-pg/pg/v10"
	"github.com/pkg/errors"

	"in-backend/services/profile"
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

func (r repository) CreateCandidate(ctx context.Context, c *orm.Candidate) (*orm.Candidate, error) {
	if c == nil {
		return nil, errors.New("Input parameter skill is nil")
	}

	result, err := r.DB.WithContext(ctx).Model(c).Returning("*").Insert()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to insert skill %v", c)
	}

	if result != nil {
		if result.RowsAffected() == 0 {
			return nil, errors.New("Failed to insert, affected is 0")
		}
	}

	return c, nil
}

func (r repository) GetAllCandidates(ctx context.Context) ([]*orm.Candidate, error) {
	var c []*orm.Candidate
	err := r.DB.WithContext(ctx).Model(&c).Select()
	return c, err
}

func (r repository) GetCandidateByID(ctx context.Context, id uint64) (*orm.Candidate, error) {
	c := orm.Candidate{ID: id}
	err := r.DB.WithContext(ctx).Model(&c).Where("id = ?", id).Select()
	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		// if row is empty than return empty model
		return &c, nil
	}
	return &c, err
}

func (r repository) UpdateCandidate(ctx context.Context, c *orm.Candidate) (*orm.Candidate, error) {
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
	c := &orm.Candidate{ID: id}
	_, err := r.DB.WithContext(ctx).Model(c).WherePK().Delete()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Cannot delete candidate with id %v", id))
	}
	return nil
}

/* --------------- Skill --------------- */

func (r repository) CreateSkill(ctx context.Context, s *orm.Skill) (*orm.Skill, error) {
	if s == nil {
		return nil, errors.New("Input parameter skill is nil")
	}

	result, err := r.DB.WithContext(ctx).Model(s).Returning("*").Insert()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to insert skill %v", s)
	}

	if result != nil {
		if result.RowsAffected() == 0 {
			return nil, errors.New("Failed to insert, affected is 0")
		}
	}

	return s, nil
}

func (r repository) GetSkill(ctx context.Context, id uint64) (*orm.Skill, error) {
	s := orm.Skill{ID: id}
	err := r.DB.WithContext(ctx).Model(&s).Where("id = ?", id).Select()
	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		// if row is empty than return empty model
		return &s, nil
	}
	return &s, err
}

func (r repository) GetAllSkills(ctx context.Context) ([]*orm.Skill, error) {
	var s []*orm.Skill
	err := r.DB.WithContext(ctx).Model(&s).Select()
	return s, err
}

/* --------------- Institution --------------- */

func (r repository) CreateInstitution(ctx context.Context, i *orm.Institution) (*orm.Institution, error) {
	if i == nil {
		return nil, errors.New("Input parameter institution is nil")
	}

	result, err := r.DB.WithContext(ctx).Model(i).Returning("*").Insert()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to insert institution %v", i)
	}

	if result != nil {
		if result.RowsAffected() == 0 {
			return nil, errors.New("Failed to insert, affected is 0")
		}
	}

	return i, nil
}

func (r repository) GetInstitution(ctx context.Context, id uint64) (*orm.Institution, error) {
	i := orm.Institution{ID: id}
	err := r.DB.WithContext(ctx).Model(&i).Where("id = ?", id).Select()
	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		// if row is empty than return empty model
		return &i, nil
	}
	return &i, err
}

func (r repository) GetAllInstitutions(ctx context.Context) ([]*orm.Institution, error) {
	var i []*orm.Institution
	err := r.DB.WithContext(ctx).Model(&i).Select()
	return i, err
}

/* --------------- Course --------------- */

func (r repository) CreateCourse(ctx context.Context, c *orm.Course) (*orm.Course, error) {
	if c == nil {
		return nil, errors.New("Input parameter course is nil")
	}

	result, err := r.DB.WithContext(ctx).Model(c).Returning("*").Insert()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to insert course %v", c)
	}

	if result != nil {
		if result.RowsAffected() == 0 {
			return nil, errors.New("Failed to insert, affected is 0")
		}
	}

	return c, nil
}

func (r repository) GetCourse(ctx context.Context, id uint64) (*orm.Course, error) {
	c := orm.Course{ID: id}
	err := r.DB.WithContext(ctx).Model(&c).Where("id = ?", id).Select()
	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		// if row is empty than return empty model
		return &c, nil
	}
	return &c, err
}

func (r repository) GetAllCourses(ctx context.Context) ([]*orm.Course, error) {
	var c []*orm.Course
	err := r.DB.WithContext(ctx).Model(&c).Select()
	return c, err
}

/* --------------- Academic History --------------- */

func (r repository) CreateAcademicHistory(ctx context.Context, a *orm.AcademicHistory) (*orm.AcademicHistory, error) {
	if a == nil {
		return nil, errors.New("Input parameter skill is nil")
	}

	result, err := r.DB.WithContext(ctx).Model(a).Returning("*").Insert()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to insert skill %v", a)
	}

	if result != nil {
		if result.RowsAffected() == 0 {
			return nil, errors.New("Failed to insert, affected is 0")
		}
	}

	return a, nil
}

func (r repository) GetAcademicHistoryByID(ctx context.Context, id uint64) (*orm.AcademicHistory, error) {
	a := orm.AcademicHistory{ID: id}
	err := r.DB.WithContext(ctx).Model(&a).Where("id = ?", id).Select()
	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		// if row is empty than return empty model
		return &a, nil
	}
	return &a, err
}

func (r repository) UpdateAcademicHistory(ctx context.Context, a *orm.AcademicHistory) (*orm.AcademicHistory, error) {
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
	a := &orm.AcademicHistory{ID: id}
	_, err := r.DB.WithContext(ctx).Model(a).WherePK().Delete()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Cannot delete academic history with id %v", id))
	}
	return nil
}

/* --------------- Company --------------- */

func (r repository) CreateCompany(ctx context.Context, c *orm.Company) (*orm.Company, error) {
	if c == nil {
		return nil, errors.New("Input parameter company is nil")
	}

	result, err := r.DB.WithContext(ctx).Model(c).Returning("*").Insert()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to insert company %v", c)
	}

	if result != nil {
		if result.RowsAffected() == 0 {
			return nil, errors.New("Failed to insert, affected is 0")
		}
	}

	return c, nil
}

func (r repository) GetCompany(ctx context.Context, id uint64) (*orm.Company, error) {
	c := orm.Company{ID: id}
	err := r.DB.WithContext(ctx).Model(&c).Where("id = ?", id).Select()
	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		// if row is empty than return empty model
		return &c, nil
	}
	return &c, err
}

func (r repository) GetAllCompanies(ctx context.Context) ([]*orm.Company, error) {
	var c []*orm.Company
	err := r.DB.WithContext(ctx).Model(&c).Select()
	return c, err
}

/* --------------- Department --------------- */

func (r repository) CreateDepartment(ctx context.Context, d *orm.Department) (*orm.Department, error) {
	if d == nil {
		return nil, errors.New("Input parameter department is nil")
	}

	result, err := r.DB.WithContext(ctx).Model(d).Returning("*").Insert()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to insert department %v", d)
	}

	if result != nil {
		if result.RowsAffected() == 0 {
			return nil, errors.New("Failed to insert, affected is 0")
		}
	}

	return d, nil
}

func (r repository) GetDepartment(ctx context.Context, id uint64) (*orm.Department, error) {
	d := orm.Department{ID: id}
	err := r.DB.WithContext(ctx).Model(&d).Where("id = ?", id).Select()
	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		// if row is empty than return empty model
		return &d, nil
	}
	return &d, err
}

func (r repository) GetAllDepartments(ctx context.Context) ([]*orm.Department, error) {
	var d []*orm.Department
	err := r.DB.WithContext(ctx).Model(&d).Select()
	return d, err
}
