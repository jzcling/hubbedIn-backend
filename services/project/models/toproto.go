package models

import (
	"in-backend/helpers"
	"in-backend/services/project/pb"
)

// ToProto maps the ORM Project model to the proto model
func (m *Project) ToProto() *pb.Project {
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

	return &pb.Project{
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
		ProjectId:             m.ProjectID,
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

// ToProto maps the ORM CandidateProject model to the proto model
func (m *CandidateProject) ToProto() *pb.CandidateProject {
	if m == nil {
		return nil
	}

	return &pb.CandidateProject{
		Id:          m.ID,
		CandidateId: m.CandidateID,
		ProjectId:   m.ProjectID,
	}
}
