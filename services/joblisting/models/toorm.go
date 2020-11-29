package models

import (
	"in-backend/helpers"
	"in-backend/services/joblisting/pb"
)

// JoblistingToORM maps the proto Joblisting model to the ORM model
func JoblistingToORM(m *pb.Joblisting) *Joblisting {
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

	return &Joblisting{
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
		JoblistingID:          m.JoblistingId,
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

// CandidateJoblistingToORM maps the proto Rating model to the ORM model
func CandidateJoblistingToORM(m *pb.CandidateJoblisting) *CandidateJoblisting {
	if m == nil {
		return nil
	}

	return &CandidateJoblisting{
		ID:           m.Id,
		CandidateID:  m.CandidateId,
		JoblistingID: m.JoblistingId,
	}
}
