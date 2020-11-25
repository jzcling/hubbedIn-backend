package database

import (
	"context"
	"fmt"

	pg "github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/pkg/errors"

	"in-backend/services/project"
	"in-backend/services/project/models"
)

// Repository implements the project Repository interface
type repository struct {
	DB *pg.DB
}

// NewRepository declares a new Repository that implements project Repository
func NewRepository(db *pg.DB) project.Repository {
	return &repository{
		DB: db,
	}
}

/* --------------- Project --------------- */

// CreateProject creates a new Project
func (r *repository) CreateProject(ctx context.Context, m *models.Project) (*models.Project, error) {
	if m == nil {
		return nil, errors.New("Input parameter project is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(m).
		Relation(relProjectRating).
		Returning("*").
		Insert()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to insert project %v", m)
	}

	return m, nil
}

// GetAllProjects returns all Projects
func (r *repository) GetAllProjects(ctx context.Context) ([]*models.Project, error) {
	var m []*models.Project
	err := r.DB.WithContext(ctx).Model(&m).
		Relation(relProjectRating).
		Returning("*").
		Select()
	return m, err
}

// GetProjectByID returns a Project by ID
func (r *repository) GetProjectByID(ctx context.Context, id uint64) (*models.Project, error) {
	m := models.Project{ID: id}
	err := r.DB.WithContext(ctx).Model(&m).
		Where(filProjectID, id).
		Relation(relProjectRating).
		Returning("*").
		Select()
	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return &m, err
}

// UpdateProject updates a Project
func (r *repository) UpdateProject(ctx context.Context, m *models.Project) (*models.Project, error) {
	if m == nil {
		return nil, errors.New("Project is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(m).WherePK().
		Relation(relProjectRating).
		Returning("*").
		Update()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Cannot update project with id %v", m.ID))
	}

	return m, nil
}

// DeleteProject deletes a Project by ID
func (r *repository) DeleteProject(ctx context.Context, id uint64) error {
	m := &models.Project{ID: id}
	_, err := r.DB.WithContext(ctx).Model(m).WherePK().Delete()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Cannot delete project with id %v", id))
	}
	return nil
}

/* --------------- Candidate Project --------------- */

// CreateCandidateProject creates a new CandidateProject
func (r *repository) CreateCandidateProject(ctx context.Context, m *models.CandidateProject) error {
	if m == nil {
		return errors.New("Input parameter candidate project is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(m).Insert()
	if err != nil {
		return errors.Wrapf(err, "Failed to insert candidate project %v", m)
	}

	return nil
}

// DeleteCandidateProject deletes a CandidateProject by Candidate ID and Project ID
func (r *repository) DeleteCandidateProject(ctx context.Context, cid, pid uint64) error {
	m := &models.CandidateProject{CandidateID: cid, ProjectID: pid}
	_, err := r.DB.WithContext(ctx).Model(m).
		Where("candidate_id = ?", cid).
		Where("project_id = ?", pid).
		Delete()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Cannot delete candidate project for candidate id %v and project id %v", cid, pid))
	}
	return nil
}

// GetAllProjectsByCandidate returns all Projects by a Candidate
func (r *repository) GetAllProjectsByCandidate(ctx context.Context, cid uint64) ([]*models.Project, error) {
	var m []*models.Project
	err := r.DB.WithContext(ctx).Model(&m).
		Relation(relProjectRating, func(q *orm.Query) (*orm.Query, error) {
			return q.Where("candidate_id = ?", cid), nil
		}).
		Returning("*").
		Select()
	return m, err
}

/* --------------- Rating --------------- */

// CreateRating creates a new Rating
func (r *repository) CreateRating(ctx context.Context, m *models.Rating) error {
	if m == nil {
		return errors.New("Input parameter rating is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(m).Insert()
	if err != nil {
		return errors.Wrapf(err, "Failed to insert rating %v", m)
	}

	return nil
}

// DeleteRating deletes a Rating by Candidate ID and Project ID
func (r *repository) DeleteRating(ctx context.Context, id uint64) error {
	m := &models.Rating{ID: id}
	_, err := r.DB.WithContext(ctx).Model(m).
		Where(filRatingID, id).
		Delete()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Cannot delete rating with id %v ", id))
	}
	return nil
}
