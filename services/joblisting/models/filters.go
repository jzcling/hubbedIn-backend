package models

import "time"

// JobPostFilters define filters for JobPost model
type JobPostFilters struct {
	ID                 []uint64
	CompanyID          []uint64
	HRContactID        []uint64
	HiringManagerID    []uint64
	JobPlatformID      []uint64
	SkillID            []uint64
	Title              string
	SeniorityLevel     []string
	MinYearsExperience uint64
	MaxYearsExperience uint64
	EmploymentType     []string
	FunctionID         []uint64
	IndustryID         []uint64
	Remote             bool
	Salary             uint64
	UpdatedAt          *time.Time
	ExpireAt           *time.Time
}

// CompanyFilters define filters for Company model
type CompanyFilters struct {
	ID   []uint64
	Name string
}

// IndustryFilters define filters for Industry model
type IndustryFilters struct {
	ID   []uint64
	Name string
}

// KeyPersonFilters define filters for KeyPerson model
type KeyPersonFilters struct {
	ID            []uint64
	CompanyID     []uint64
	Name          string
	ContactNumber string
	Email         string
	JobTitle      string
}

// JobPlatformFilters define filters for JobPlatform model
type JobPlatformFilters struct {
	ID   []uint64
	Name string
}

// JobFunctionFilters define filters for JobFunction model
type JobFunctionFilters struct {
	ID   []uint64
	Name string
}
