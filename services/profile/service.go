package profile

import (
	"context"

	"in-backend/services/profile/models"
)

// Service describes the Profile Service.
type Service interface {
	// CreateCandidate a new candidate
	CreateCandidate(ctx context.Context, candidate *models.Candidate) (models.Candidate, error)

	// GetAllCandidates returns all candidates
	GetAllCandidates(ctx context.Context) ([]models.Candidate, error)

	// GetCandidateByID finds and returns a candidate by ID
	GetCandidateByID(ctx context.Context, id uint64) (models.Candidate, error)

	// UpdateCandidate candidate
	UpdateCandidate(ctx context.Context, candidate *models.Candidate) (models.Candidate, error)

	// DeleteCandidate deletes a candidate by ID
	DeleteCandidate(ctx context.Context, id uint64) error
}
