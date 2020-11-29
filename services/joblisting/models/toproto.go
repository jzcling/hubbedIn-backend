package models

import (
	"in-backend/helpers"
	"in-backend/services/joblisting/pb"
)

// ToProto maps the ORM Joblisting model to the proto model
func (m *Joblisting) ToProto() *pb.Joblisting {
	if m == nil {
		return nil
	}

	var ratings []*pb.Rating
	r := m.Ratings
	for _, rating := range r {
		ratings = append(ratings, rating.ToProto())
	}

	createdAt := helpers.TimeToProto(m.CreatedAt)
	updatedAt := helpers.TimeToProto(m.UpdatedAt)
	deletedAt := helpers.TimeToProto(m.DeletedAt)

	return &pb.Joblisting{
		Id:        m.ID,
		Name:      m.Name,
		RepoUrl:   m.RepoURL,
		Ratings:   ratings,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}
}

// ToProto maps the ORM Rating model to the proto model
func (m *Rating) ToProto() *pb.Rating {
	if m == nil {
		return nil
	}

	createdAt := helpers.TimeToProto(m.CreatedAt)

	return &pb.Rating{
		Id:                    m.ID,
		JoblistingId:          m.JoblistingID,
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

// ToProto maps the ORM CandidateJoblisting model to the proto model
func (m *CandidateJoblisting) ToProto() *pb.CandidateJoblisting {
	if m == nil {
		return nil
	}

	return &pb.CandidateJoblisting{
		Id:           m.ID,
		CandidateId:  m.CandidateID,
		JoblistingId: m.JoblistingID,
	}
}
