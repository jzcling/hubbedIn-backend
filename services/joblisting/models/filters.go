package models

// JoblistingFilters define filters for Joblisting model
type JoblistingFilters struct {
	ID          []uint64
	CandidateID uint64
	Name        string
	RepoURL     string
}

// RatingFilters define filters for Rating model
type RatingFilters struct {
	JoblistingID          []uint64
	ReliabilityRating     []int32
	MaintainabilityRating []int32
	SecurityRating        []int32
	SecurityReviewRating  []int32
}
