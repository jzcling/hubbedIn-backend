package service

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"in-backend/services/project"
	"in-backend/services/project/models"
)

// Service implements the project Service interface
type service struct {
	repository project.Repository
	logger     log.Logger
}

// New creates and returns a new Service that implements the project Service interface
func New(r project.Repository, l log.Logger) project.Service {
	return &service{
		repository: r,
		logger:     l,
	}
}

/* --------------- Project --------------- */

// CreateProject creates a new Project
func (s *service) CreateProject(ctx context.Context, model *models.Project) (*models.Project, error) {
	logger := log.With(s.logger, "method", "CreateProject")

	m, err := s.repository.CreateProject(ctx, model)
	if err != nil {
		level.Error(logger).Log("err", err)
	}

	return m, err
}

// GetAllProjects returns all Projects
func (s *service) GetAllProjects(ctx context.Context) ([]*models.Project, error) {
	logger := log.With(s.logger, "method", "GetAllProjects")

	m, err := s.repository.GetAllProjects(ctx)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return m, err
}

// GetProjectByID returns a Project by ID
func (s *service) GetProjectByID(ctx context.Context, id uint64) (*models.Project, error) {
	logger := log.With(s.logger, "method", "GetProjectByID")

	m, err := s.repository.GetProjectByID(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return m, err
}

// UpdateProject updates a Project
func (s *service) UpdateProject(ctx context.Context, model *models.Project) (*models.Project, error) {
	logger := log.With(s.logger, "method", "UpdateProject")

	m, err := s.repository.UpdateProject(ctx, model)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return m, err
}

// DeleteProject deletes a Project by ID
func (s *service) DeleteProject(ctx context.Context, id uint64) error {
	logger := log.With(s.logger, "method", "DeleteProject")

	err := s.repository.DeleteProject(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return err
}

// ScanProject scans a Project using sonarqube
func (s *service) ScanProject(ctx context.Context, id uint64) error {
	_ = log.With(s.logger, "method", "ScanProject")

	// TODO: clone repo
	// TODO: add properties file
	// TODO: run scanner
	return nil
}

/* --------------- Candidate Project --------------- */

// CreateCandidateProject creates a new CandidateProject
func (s *service) CreateCandidateProject(ctx context.Context, m *models.CandidateProject) error {
	logger := log.With(s.logger, "method", "CreateCandidateProject")

	err := s.repository.CreateCandidateProject(ctx, m)
	if err != nil {
		level.Error(logger).Log("err", err)
	}

	return err
}

// DeleteCandidateProject deletes a CandidateProject by Candidate ID and Project ID
func (s *service) DeleteCandidateProject(ctx context.Context, cid, pid uint64) error {
	logger := log.With(s.logger, "method", "DeleteCandidateProject")

	err := s.repository.DeleteCandidateProject(ctx, cid, pid)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return err
}

// GetAllProjectsByCandidate returns all Projects by a Candidate
func (s *service) GetAllProjectsByCandidate(ctx context.Context, cid uint64) ([]*models.Project, error) {
	logger := log.With(s.logger, "method", "GetAllProjectsByCandidate")

	c, err := s.repository.GetAllProjectsByCandidate(ctx, cid)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return c, err
}

/* --------------- Rating --------------- */

// CreateRating creates a new Rating
func (s *service) CreateRating(ctx context.Context, m *models.Rating) error {
	logger := log.With(s.logger, "method", "CreateRating")

	err := s.repository.CreateRating(ctx, m)
	if err != nil {
		level.Error(logger).Log("err", err)
	}

	return err
}

// DeleteRating deletes a Rating by Candidate ID and Project ID
func (s *service) DeleteRating(ctx context.Context, id uint64) error {
	logger := log.With(s.logger, "method", "DeleteRating")

	err := s.repository.DeleteRating(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return err
}
