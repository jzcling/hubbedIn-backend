package joblisting

import (
	"context"
	"in-backend/services/joblisting/models"
)

// Repository declares the repository for joblistings
type Repository interface {
	/* --------------- Joblisting --------------- */

	// CreateJoblisting creates a new candidate
	CreateJoblisting(ctx context.Context, m *models.Joblisting) (*models.Joblisting, error)

	// GetAllJoblistings returns all candidates
	GetAllJoblistings(ctx context.Context, f models.JoblistingFilters) ([]*models.Joblisting, error)

	// GetJoblistingByID finds and returns a candidate by ID
	GetJoblistingByID(ctx context.Context, id uint64) (*models.Joblisting, error)

	// UpdateJoblisting updates a candidate
	UpdateJoblisting(ctx context.Context, m *models.Joblisting) (*models.Joblisting, error)

	// DeleteJoblisting deletes a candidate by ID
	DeleteJoblisting(ctx context.Context, id uint64) error

	/* --------------- Candidate Joblisting --------------- */

	// CreateCandidateJoblisting creates a new Candidate Joblisting
	CreateCandidateJoblisting(ctx context.Context, m *models.CandidateJoblisting) error

	// DeleteCandidateroject deletes a Candidate Joblisting
	DeleteCandidateJoblisting(ctx context.Context, id uint64) error

	/* --------------- Rating --------------- */

	// CreateRating creates a new Joblisting Rating
	CreateRating(ctx context.Context, m *models.Rating) error

	// DeleteRating deletes a Joblisting Rating
	DeleteRating(ctx context.Context, id uint64) error
}
