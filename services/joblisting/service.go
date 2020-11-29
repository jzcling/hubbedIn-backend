package joblisting

import (
	"context"
	"in-backend/services/joblisting/models"
)

// Service describes the Joblisting Service
type Service interface {
	/* --------------- Joblisting --------------- */

	// CreateJoblisting creates a new Joblisting
	CreateJoblisting(ctx context.Context, m *models.Joblisting, cid uint64) (*models.Joblisting, error)

	// GetAllJoblistings returns all Joblistings
	GetAllJoblistings(ctx context.Context, f models.JoblistingFilters) ([]*models.Joblisting, error)

	// GetJoblistingByID finds and returns a Joblisting by ID
	GetJoblistingByID(ctx context.Context, id uint64) (*models.Joblisting, error)

	// UpdateJoblisting updates a Joblisting
	UpdateJoblisting(ctx context.Context, m *models.Joblisting) (*models.Joblisting, error)

	// DeleteJoblisting deletes a Joblisting by ID
	DeleteJoblisting(ctx context.Context, id uint64) error

	// ScanJoblisting scans a Joblisting using sonarqube
	ScanJoblisting(ctx context.Context, id uint64) error

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
