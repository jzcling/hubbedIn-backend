package project

import (
	"context"
	"in-backend/services/project/models"
)

// Service describes the Profile Service
type Service interface {
	/* --------------- Project --------------- */

	// CreateProject creates a new Project
	CreateProject(ctx context.Context, c *models.Project) (*models.Project, error)

	// GetAllProjects returns all Projects
	GetAllProjects(ctx context.Context) ([]*models.Project, error)

	// GetProjectByID finds and returns a Project by ID
	GetProjectByID(ctx context.Context, id uint64) (*models.Project, error)

	// UpdateProject updates a Project
	UpdateProject(ctx context.Context, c *models.Project) (*models.Project, error)

	// DeleteProject deletes a Project by ID
	DeleteProject(ctx context.Context, id uint64) error

	// ScanProject scans a Project using sonarqube
	ScanProject(ctx context.Context, id uint64) error

	/* --------------- Candidate Project --------------- */

	// CreateCandidateProject creates a new Candidate Project
	CreateCandidateProject(ctx context.Context, c *models.CandidateProject) error

	// DeleteCandidateroject deletes a Candidate Project
	DeleteCandidateProject(ctx context.Context, cid, pid uint64) error

	// GetAllProjectsByCandidate returns all Projects by a Candidate
	GetAllProjectsByCandidate(ctx context.Context, cid uint64) ([]*models.Project, error)

	/* --------------- Rating --------------- */

	// CreateRating creates a new Project Rating
	CreateRating(ctx context.Context, c *models.Rating) error

	// DeleteRating deletes a Project Rating
	DeleteRating(ctx context.Context, id uint64) error
}
