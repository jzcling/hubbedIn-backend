package database

import (
	"context"
	"strings"

	pg "github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"

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
	q := r.DB.WithContext(ctx).Model(&m).Relation(relProjectRating, func(q *orm.Query) (*orm.Query, error) {
		return q.Order("r.created_at desc"), nil
	})

	if len(f.ID) > 0 {
		q = q.Where("p.id in (?)", pg.In(f.ID))
	}
	if f.CandidateID > 0 {
		cp := &models.CandidateProject{CandidateID: f.CandidateID}
		var pids []uint64
		err := r.DB.WithContext(ctx).Model(cp).
			Where("candidate_id = ?", f.CandidateID).
			Returning("project_id").
			Select(&pids)
		if err != nil {
			return nil, candidateIDErr(err, f.CandidateID)
		}

		/*
			Written like this because if pids is empty, a syntax error
			is thrown due to pg.In's handling. In the case where there
			are no project IDs, we should return an empty result and a proxy
			for that is to condition on id is null
		*/
		if len(pids) > 0 {
			q = q.Where("p.id in (?)", pg.In(pids))
		} else {
			q = q.Where("p.id is null")
		}
	}
	if f.Name != "" {
		q = q.Where("lower(p.name) like ?", "%"+strings.ToLower(f.Name)+"%")
	}
	if f.RepoURL != "" {
		q = q.Where("lower(p.repo_url) like ?", "%"+strings.ToLower(f.RepoURL)+"%")
	}

	err := q.Returning("*").Select()
	return m, err
}

// GetProjectByID returns a Project by ID
func (r *repository) GetProjectByID(ctx context.Context, id uint64) (*models.Project, error) {
	m := &models.Project{ID: id}
	err := r.DB.WithContext(ctx).Model(m).
		Where(filProjectID, id).
		Relation(relProjectRating).
		Returning("*").
		First()
	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return m, err
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

// GetCandidateProjectByID gets a CandidateProject by ID
func (r *repository) GetCandidateProjectByID(ctx context.Context, id uint64) (*models.CandidateProject, error) {
	m := &models.CandidateProject{ID: id}
	err := r.DB.WithContext(ctx).Model(m).
		Where(filID, id).
		Returning("*").
		First()

	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return m, err
}

// GetCandidateProject gets a CandidateProject by Candidate ID and Project ID
func (r *repository) GetCandidateProject(ctx context.Context, cid, pid uint64) (*models.CandidateProject, error) {
	m := &models.CandidateProject{CandidateID: cid, ProjectID: pid}
	err := r.DB.WithContext(ctx).Model(m).
		Where("candidate_id", cid).
		Where("projecT_id", pid).
		Returning("*").
		First()

	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return m, err
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
