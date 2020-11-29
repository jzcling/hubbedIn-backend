package database

import (
	"context"
	"strings"

	pg "github.com/go-pg/pg/v10"

	"in-backend/services/joblisting"
	"in-backend/services/joblisting/models"
)

// Repository implements the joblisting Repository interface
type repository struct {
	DB *pg.DB
}

// NewRepository declares a new Repository that implements joblisting Repository
func NewRepository(db *pg.DB) joblisting.Repository {
	return &repository{
		DB: db,
	}
}

/* --------------- Joblisting --------------- */

// CreateJoblisting creates a new Joblisting
func (r *repository) CreateJoblisting(ctx context.Context, m *models.Joblisting) (*models.Joblisting, error) {
	if m == nil {
		return nil, nilErr("joblisting")
	}

	_, err := r.DB.WithContext(ctx).Model(m).
		Relation(relJoblistingRating).
		Returning("*").
		Insert()
	if err != nil {
		return nil, failedToInsertErr(err, "joblisting", m)
	}

	return m, nil
}

// GetAllJoblistings returns all Joblistings
func (r *repository) GetAllJoblistings(ctx context.Context, f models.JoblistingFilters) ([]*models.Joblisting, error) {
	var m []*models.Joblisting
	q := r.DB.WithContext(ctx).Model(&m).Relation(relJoblistingRating)

	if len(f.ID) > 0 {
		q = q.Where("p.id in (?)", pg.In(f.ID))
	}
	if f.CandidateID > 0 {
		cp := &models.CandidateJoblisting{CandidateID: f.CandidateID}
		var pids []uint64
		err := r.DB.WithContext(ctx).Model(cp).
			Where("candidate_id = ?", f.CandidateID).
			Returning("joblisting_id").
			Select(&pids)
		if err != nil {
			return nil, candidateIDErr(err, f.CandidateID)
		}

		/*
			Written like this because if pids is empty, a syntax error
			is thrown due to pg.In's handling. In the case where there
			are no joblisting IDs, we should return an empty result and a proxy
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

// GetJoblistingByID returns a Joblisting by ID
func (r *repository) GetJoblistingByID(ctx context.Context, id uint64) (*models.Joblisting, error) {
	m := models.Joblisting{ID: id}
	err := r.DB.WithContext(ctx).Model(&m).
		Where(filJoblistingID, id).
		Relation(relJoblistingRating).
		Returning("*").
		First()
	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return &m, err
}

// UpdateJoblisting updates a Joblisting
func (r *repository) UpdateJoblisting(ctx context.Context, m *models.Joblisting) (*models.Joblisting, error) {
	if m == nil {
		return nil, nilErr("joblisting")
	}

	_, err := r.DB.WithContext(ctx).Model(m).WherePK().
		Relation(relJoblistingRating).
		Returning("*").
		Update()
	if err != nil {
		return nil, updateErr(err, "joblisting", m.ID)
	}

	return m, nil
}

// DeleteJoblisting deletes a Joblisting by ID
func (r *repository) DeleteJoblisting(ctx context.Context, id uint64) error {
	m := &models.Joblisting{ID: id}
	_, err := r.DB.WithContext(ctx).Model(m).WherePK().Delete()
	if err != nil {
		return deleteErr(err, "joblisting", id)
	}
	return nil
}

/* --------------- Candidate Joblisting --------------- */

// CreateCandidateJoblisting creates a new CandidateJoblisting
func (r *repository) CreateCandidateJoblisting(ctx context.Context, m *models.CandidateJoblisting) error {
	if m == nil {
		return nilErr("candidate joblisting")
	}

	_, err := r.DB.WithContext(ctx).Model(m).Insert()
	if err != nil {
		return failedToInsertErr(err, "candidate joblisting", m)
	}

	return nil
}

// DeleteCandidateJoblisting deletes a CandidateJoblisting by ID
func (r *repository) DeleteCandidateJoblisting(ctx context.Context, id uint64) error {
	m := &models.CandidateJoblisting{ID: id}
	_, err := r.DB.WithContext(ctx).Model(m).WherePK().Delete()
	if err != nil {
		return deleteErr(err, "candidate joblisting", id)
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

// DeleteRating deletes a Rating by Candidate ID and Joblisting ID
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
