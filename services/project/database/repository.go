package database

import (
	"context"
	"fmt"

	pg "github.com/go-pg/pg/v10"

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
		return nil, nilErr("project")
	}

	_, err := r.DB.WithContext(ctx).Model(m).
		Relation(relProjectRating).
		Returning("*").
		Insert()
	if err != nil {
		return nil, failedToInsertErr(err, "project", m)
	}

	return m, nil
}

// GetAllProjects returns all Projects
func (r *repository) GetAllProjects(ctx context.Context, f models.ProjectFilters) ([]*models.Project, error) {
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
		First()
	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return &m, err
}

// UpdateProject updates a Project
func (r *repository) UpdateProject(ctx context.Context, m *models.Project) (*models.Project, error) {
	if m == nil {
		return nil, nilErr("project")
	}

	_, err := r.DB.WithContext(ctx).Model(m).WherePK().
		Relation(relProjectRating).
		Returning("*").
		Update()
	if err != nil {
		return nil, updateErr(err, "project", m.ID)
	}

	return m, nil
}

// DeleteProject deletes a Project by ID
func (r *repository) DeleteProject(ctx context.Context, id uint64) error {
	m := &models.Project{ID: id}
	_, err := r.DB.WithContext(ctx).Model(m).WherePK().Delete()
	if err != nil {
		return deleteErr(err, "project", id)
	}
	return nil
}

/* --------------- Candidate Project --------------- */

// CreateCandidateProject creates a new CandidateProject
func (r *repository) CreateCandidateProject(ctx context.Context, m *models.CandidateProject) error {
	if m == nil {
		return nilErr("candidate project")
	}

	_, err := r.DB.WithContext(ctx).Model(m).Insert()
	if err != nil {
		return failedToInsertErr(err, "candidate project", m)
	}

	return nil
}

// DeleteCandidateProject deletes a CandidateProject by ID
func (r *repository) DeleteCandidateProject(ctx context.Context, id uint64) error {
	m := &models.CandidateProject{ID: id}
	_, err := r.DB.WithContext(ctx).Model(m).WherePK().Delete()
	if err != nil {
		return deleteErr(err, "candidate project", id)
	}
	return nil
}

// GetAllProjectsByCandidate returns all Projects by a Candidate
func (r *repository) GetAllProjectsByCandidate(ctx context.Context, cid uint64) ([]*models.Project, error) {
	cp := &models.CandidateProject{CandidateID: cid}
	fmt.Printf("%v\n", cid)
	var pids []uint64
	err := r.DB.WithContext(ctx).Model(cp).
		Where("candidate_id = ?", cid).
		Returning("project_id").
		Select(&pids)
	if err != nil {
		return nil, candidateIDErr(err, cid)
	}
	fmt.Printf("%v\n", pids)

	var m []*models.Project
	err = r.DB.WithContext(ctx).Model(&m).
		Where("p.id in (?)", pg.In(pids)).
		Relation(relProjectRating).
		Returning("*").
		Select()
	return m, err
}

/* --------------- Rating --------------- */

// CreateRating creates a new Rating
func (r *repository) CreateRating(ctx context.Context, m *models.Rating) error {
	if m == nil {
		return nilErr("rating")
	}

	_, err := r.DB.WithContext(ctx).Model(m).Insert()
	if err != nil {
		return failedToInsertErr(err, "rating", m)
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
		return deleteErr(err, "rating", id)
	}
	return nil
}
