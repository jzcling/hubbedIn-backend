package service

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"in-backend/services/profile"
	"in-backend/services/profile/models"
)

// service implements the Profile Service
type service struct {
	repository profile.Repository
	logger     log.Logger
}

// New creates and returns a new Profile Service instance
func New(r profile.Repository, l log.Logger) profile.Service {
	return &service{
		repository: r,
		logger:     l,
	}
}

// CreateCandidate creates a new candidate
func (s *service) CreateCandidate(ctx context.Context, candidate *models.Candidate) (models.Candidate, error) {
	logger := log.With(s.logger, "method", "CreateCandidate")

	c, err := s.repository.CreateCandidate(ctx, candidate)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return c, err
}

// GetAllCandidates returns all candidates
func (s *service) GetAllCandidates(ctx context.Context) ([]models.Candidate, error) {
	logger := log.With(s.logger, "method", "GetAllCandidates")

	c, err := s.repository.GetAllCandidates(ctx)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return c, err
}

// GetCandidateByID returns a candidate by ID
func (s *service) GetCandidateByID(ctx context.Context, id uint64) (models.Candidate, error) {
	logger := log.With(s.logger, "method", "GetCandidateByID")

	c, err := s.repository.GetCandidateByID(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return c, err
}

// UpdateCandidate updates a candidate
func (s *service) UpdateCandidate(ctx context.Context, candidate *models.Candidate) (models.Candidate, error) {
	logger := log.With(s.logger, "method", "UpdateCandidate")

	c, err := s.repository.UpdateCandidate(ctx, candidate)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return c, err
}

// DeleteCandidate deletes a candidate by ID
func (s *service) DeleteCandidate(ctx context.Context, id uint64) error {
	logger := log.With(s.logger, "method", "DeleteCandidate")

	err := s.repository.DeleteCandidate(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return err
}
