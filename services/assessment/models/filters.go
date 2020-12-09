package models

// AssessmentFilters define filters for Assessment model
type AssessmentFilters struct {
	ID         []uint64
	Name       string
	Difficulty []string
	Type       []string

	// relation filters
	CandidateID uint64
	Status      []string
	MinScore    uint8
}

// QuestionFilters define filters for Question model
type QuestionFilters struct {
	ID   []uint64
	Tags []string
}
