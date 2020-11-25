package models

import (
	"in-backend/helpers"
	"in-backend/services/project/pb"
)

// ProjectToORM maps the proto Project model to the ORM model
func ProjectToORM(m *pb.Project) *Project {
	if m == nil {
		return nil
	}

	var ratings []*Rating
	r := m.Ratings
	for _, rating := range r {
		ratings = append(ratings, RatingToORM(rating))
	}

	createdAt := helpers.ProtoTimeToTime(m.CreatedAt)
	updatedAt := helpers.ProtoTimeToTime(m.UpdatedAt)
	deletedAt := helpers.ProtoTimeToTime(m.DeletedAt)

	return &Project{
		ID:        m.Id,
		Name:      m.Name,
		RepoURL:   m.RepoUrl,
		Ratings:   ratings,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}
}

// RatingToORM maps the proto Rating model to the ORM model
func RatingToORM(m *pb.Rating) *Rating {
	if m == nil {
		return nil
	}

	createdAt := helpers.ProtoTimeToTime(m.CreatedAt)

	return &Rating{
		ID:                    m.Id,
		ProjectID:             m.ProjectId,
		ReliabilityRating:     m.ReliabilityRating,
		MaintainabilityRating: m.MaintainabilityRating,
		SecurityRating:        m.SecurityRating,
		SecurityReviewRating:  m.SecurityReviewRating,
		Coverage:              m.Coverage,
		Duplications:          m.Duplications,
		Lines:                 m.Lines,
		CreatedAt:             createdAt,
	}
}

// CandidateProjectToORM maps the proto Rating model to the ORM model
func CandidateProjectToORM(m *pb.CandidateProject) *CandidateProject {
	if m == nil {
		return nil
	}

	return &CandidateProject{
		ID:          m.Id,
		CandidateID: m.CandidateId,
		ProjectID:   m.ProjectId,
	}
}
