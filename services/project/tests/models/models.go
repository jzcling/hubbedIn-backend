package models

import (
	"in-backend/services/project/models"
	"time"
)

var (
	now time.Time = time.Now()

	// ProjectNoName is a Project without a name (not null field)
	ProjectNoName = &models.Project{
		RepoURL:   "repo",
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	// ProjectValid is a valid Project that should pass tests
	ProjectValid = &models.Project{
		Name:      "name",
		RepoURL:   "repo",
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	// CandidateProjectValid is a valid CandidateProject that should pass tests
	CandidateProjectValid = &models.CandidateProject{
		CandidateID: 1,
		ProjectID:   1,
	}

	// RatingValid is a valid CandidateProject that should pass tests
	RatingValid = &models.Rating{
		ProjectID:             1,
		ReliabilityRating:     1,
		MaintainabilityRating: 1,
		SecurityRating:        1,
		SecurityReviewRating:  1,
		Coverage:              1.0,
		Duplications:          1.0,
		Lines:                 1,
		CreatedAt:             &now,
	}
)
