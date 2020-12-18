package project

import (
	"context"
	"in-backend/services/project/models"
)

// Repository declares the repository for projects
type Repository interface {
	/* --------------- Project --------------- */

	// CreateProject creates a new candidate
	CreateProject(ctx context.Context, m *models.Project) (*models.Project, error)

	// GetAllProjects returns all candidates
	GetAllProjects(ctx context.Context, f models.ProjectFilters) ([]*models.Project, error)

	// GetProjectByID finds and returns a candidate by ID
	GetProjectByID(ctx context.Context, id uint64) (*models.Project, error)

	// UpdateProject updates a candidate
	UpdateProject(ctx context.Context, m *models.Project) (*models.Project, error)

	// DeleteProject deletes a candidate by ID
	DeleteProject(ctx context.Context, id uint64) error

	/* --------------- Candidate Project --------------- */

	// CreateCandidateProject creates a new Candidate Project
	CreateCandidateProject(ctx context.Context, m *models.CandidateProject) error

	// GetCandidateProjectByID gets a CandidateProject by ID
	GetCandidateProjectByID(ctx context.Context, id uint64) (*models.CandidateProject, error)

	// GetCandidateProject gets a CandidateProject by Candidate ID and Project ID
	GetCandidateProject(ctx context.Context, cid, pid uint64) (*models.CandidateProject, error)

	// DeleteCandidateroject deletes a Candidate Project
	DeleteCandidateProject(ctx context.Context, id uint64) error

	/* --------------- Rating --------------- */

	// CreateRating creates a new Project Rating
	CreateRating(ctx context.Context, m *models.Rating) error

	// DeleteRating deletes a Project Rating
	DeleteRating(ctx context.Context, id uint64) error
}
