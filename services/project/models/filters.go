package models

// ProjectFilters define filters for Project model
type ProjectFilters struct {
	ID      []uint64
	Name    string
	RepoURL string
}

// RatingFilters define filters for Rating model
type RatingFilters struct {
	ProjectID             []uint64
	ReliabilityRating     []int32
	MaintainabilityRating []int32
	SecurityRating        []int32
	SecurityReviewRating  []int32
}
