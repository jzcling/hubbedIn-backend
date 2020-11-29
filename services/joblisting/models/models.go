package models

import (
	"time"
)

// Joblisting declares the model for Joblisting
type Joblisting struct {
	tableName struct{} `pg:"joblistings,alias:jl"`

	ID             uint64      `json:"id"`
	CompanyID      uint64      `json:"company_id" pg:",notnull"`
	Company        *Company    `json:"company,omitempty" pg:"rel:has-one"`
	DepartmentID   uint64      `json:"department_id" pg:",notnull"`
	Department     *Department `json:"department,omitempty" pg:"rel:has-one"`
	Country        string      `json:"country" pg:",notnull"`
	City           string      `json:"city,omitempty"`
	Title          string      `json:"title" pg:",notnull"`
	SalaryCurrency string      `json:"salary_currency,omitempty"`
	MinSalary      uint32      `json:"min_salary,omitempty"`
	MaxSalary      uint32      `json:"max_salary,omitempty"`
	Description    string      `json:"description,omitempty"`
	Tags           []*Tag      `json:"tags,omitempty" pg:"rel:has-many"`
	CreatedAt      *time.Time  `json:"created_at,omitempty" pg:"default:now()"`
	UpdatedAt      *time.Time  `json:"updated_at,omitempty" pg:"default:now()"`
	DeletedAt      *time.Time  `json:"deleted_at,omitempty" pg:",soft_delete"`
}

// Application declares the model for the pivot table between Candidate and Joblisting
type Application struct {
	tableName struct{} `pg:"applications,alias:ap"`

	ID           uint64 `json:"id"`
	CandidateID  uint64 `json:"candidate_id"`
	JoblistingID uint64 `json:"joblisting_id"`
}

// Rating declares the model for Rating
type Rating struct {
	tableName struct{} `pg:"ratings,alias:r"`

	ID                    uint64     `json:"id"`
	JoblistingID          uint64     `json:"joblisting_id" pg:"joblisting_id"`
	ReliabilityRating     int32      `json:"reliability_rating"`
	MaintainabilityRating int32      `json:"maintainability_rating"`
	SecurityRating        int32      `json:"security_rating"`
	SecurityReviewRating  int32      `json:"security_review_rating"`
	Coverage              float32    `json:"coverage" pg:",use_zero"`
	Duplications          float32    `json:"duplications" pg:",use_zero"`
	Lines                 uint64     `json:"lines"`
	CreatedAt             *time.Time `json:"created_at,omitempty" pg:"default:now()"`
}
