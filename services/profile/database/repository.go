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

func (r repository) CreateCandidate(ctx context.Context, candidate *models.Candidate) (models.Candidate, error) {
	var c models.Candidate
	if candidate == nil {
		return c, errors.New("Input parameter candidate is nil")
	}

	result, err := r.DB.WithContext(ctx).Model(&c).Returning("*").Insert()
	if err != nil {
		return c, errors.Wrapf(err, "Failed to insert candidate %v", candidate)
	}

	if result != nil {
		if result.RowsAffected() == 0 {
			return c, errors.New("Failed to insert, affected is 0")
		}
	}

	return c, nil
}

func (r repository) GetAllCandidates(ctx context.Context) ([]models.Candidate, error) {
	var c []models.Candidate
	err := r.DB.WithContext(ctx).Model(&c).Select()
	return c, err
}

func (r repository) GetCandidateByID(ctx context.Context, id uint64) (models.Candidate, error) {
	c := models.Candidate{ID: id}
	err := r.DB.WithContext(ctx).Model(&c).Where("id = ?", id).Select()
	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		// if row is empty than return empty model
		return c, nil
	}
	return c, err
}

func (r repository) UpdateCandidate(ctx context.Context, candidate *models.Candidate) (models.Candidate, error) {
	var c models.Candidate
	if candidate == nil {
		return c, errors.New("Candidate is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(&c).Returning("*").Update()
	if err != nil {
		return c, errors.Wrap(err, fmt.Sprintf("Cannot update candidate with id %v", candidate.ID))
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
