package database

import (
	"context"
	"fmt"
	"strconv"

	pg "github.com/go-pg/pg/v10"
	"github.com/pkg/errors"
)

type repository struct {
	DB *pg.DB
}

func NewRepository(db *pg.DB) profile.Repository {
	return &repository{
		DB: db,
	}
}

func (r repository) CreateCandidate(ctx context.Context, candidate *models.Candidate) (models.Candidate, error) {
	if candidate == nil {
		return errors.New("Input parameter candidate is nil")
	}

	result, err := r.DB.WithContext(ctx).Model(candidate).Returning("*").Insert()
	if err != nil {
		return errors.Wrapf(err, "Failed to insert candidate %v", candidate)
	}

	if result != nil {
		if result.RowsAffected() == 0 {
			return errors.New("Failed to insert, affected is 0")
		}
	}

	return result, nil
}

func (r repository) GetAllCandidates(ctx context.Context) ([]models.Candidate, error) {
	var candidates []models.Candidate
	err := r.DB.WithContext(ctx).Model(&candidates).Select()
	if err != nil {
		return nil, err
	}

	if candidates == nil {
		return []models.Candidate{}, ni
	}
	return candidates, nil
}

func (r repository) GetCandidateByID(ctx context.Context, id string) (models.Candidate, error) {
	if len(id) == 0 {
		return models.Candidate{}, errors.New("Candidate id cannot be empty")
	}

	index, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return models.Candidate{}, errors.Wrap(err, "Cannot convert candidate id to integer")
	}
	candidate := &models.Candidate{ID: index}
	err = r.DB.WithContext(ctx).Model(candidate).Where("id = ?", id).Select(candidate)
	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		// if row is empty than return empty model
		return models.Candidate{}, ni
	}
	return candidate, err
}

func (r repository) UpdateCandidate(ctx context.Context, candidate *models.Candidate) error {
	if candidate == nil {
		return errors.New("Candidate is nil")
	}

	result, err := r.DB.WithContext(ctx).Model(candidate).Returning("*").Update()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Cannot update candidate with id %v", candidate.ID))
	}

	return result, nil
}

func (r repository) DeleteCandidate(ctx context.Context, id string) error {
	if len(id) == 0 {
		return errors.New("Candidate id cannot be empty")
	}

	index, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return errors.Wrap(err, "Cannot convert Candidate id")
	}

	candidate := &models.Candidate{ID: index}
	return r.DB.WithContext(ctx).Delete(candidate)
}
