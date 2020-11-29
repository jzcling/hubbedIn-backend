package models

import (
	"in-backend/services/joblisting/models"
	"time"
)

var (
	now time.Time = time.Now()

	// JoblistingNoName is a Joblisting without a name (not null field)
	JoblistingNoName = &models.Joblisting{
		RepoURL:   "repo",
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	// JoblistingValid is a valid Joblisting that should pass tests
	JoblistingValid = &models.Joblisting{
		Name:      "name",
		RepoURL:   "repo",
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	// CandidateJoblistingValid is a valid CandidateJoblisting that should pass tests
	CandidateJoblistingValid = &models.CandidateJoblisting{
		CandidateID:  1,
		JoblistingID: 1,
	}

	// RatingValid is a valid CandidateJoblisting that should pass tests
	RatingValid = &models.Rating{
		JoblistingID:          1,
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
