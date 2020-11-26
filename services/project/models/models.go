package models

import (
	"time"
)

// Project declares the model for Project
type Project struct {
	tableName struct{} `pg:"projects,alias:p"`

	ID        uint64     `json:"id"`
	Name      string     `json:"name" pg:",notnull"`
	RepoURL   string     `json:"repo_url,omitempty" pg:"repo_url,notnull,unique"`
	Ratings   []*Rating  `json:"ratings,omitempty" pg:"rel:has-many"`
	CreatedAt *time.Time `json:"created_at,omitempty" pg:"default:now()"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" pg:"default:now()"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" pg:",soft_delete"`
}

// CandidateProject declares the model for the pivot table between Candidate and Project
type CandidateProject struct {
	tableName struct{} `pg:"candidates_projects,alias:cp"`

	ID          uint64 `json:"id"`
	CandidateID uint64 `json:"candidate_id"`
	ProjectID   uint64 `json:"project_id"`
}

// Rating declares the model for Rating
type Rating struct {
	tableName struct{} `pg:"ratings,alias:r"`

	ID                    uint64     `json:"id"`
	ProjectID             uint64     `json:"project_id" pg:"project_id"`
	ReliabilityRating     int32      `json:"reliability_rating"`
	MaintainabilityRating int32      `json:"maintainability_rating"`
	SecurityRating        int32      `json:"security_rating"`
	SecurityReviewRating  int32      `json:"security_review_rating"`
	Coverage              float32    `json:"coverage" pg:",use_zero"`
	Duplications          float32    `json:"duplications" pg:",use_zero"`
	Lines                 uint64     `json:"lines"`
	CreatedAt             *time.Time `json:"created_at,omitempty" pg:"default:now()"`
}
