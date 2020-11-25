package models

import (
	"in-backend/services/project/pb"
	"testing"

	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/require"
)

func TestProjectToProto(t *testing.T) {
	testPbTime := ptypes.TimestampNow()
	testTime, err := ptypes.Timestamp(testPbTime)
	require.NoError(t, err)

	input := &Project{
		ID:      1,
		Name:    "project",
		RepoURL: "repo",
		Ratings: []*Rating{
			{
				ID:                    1,
				ProjectID:             1,
				ReliabilityRating:     1,
				MaintainabilityRating: 1,
				SecurityRating:        1,
				SecurityReviewRating:  1,
				Coverage:              1.0,
				Duplications:          1.0,
				Lines:                 1,
				CreatedAt:             &testTime,
			},
			{
				ID:                    2,
				ProjectID:             1,
				ReliabilityRating:     1,
				MaintainabilityRating: 1,
				SecurityRating:        1,
				SecurityReviewRating:  1,
				Coverage:              1.0,
				Duplications:          1.0,
				Lines:                 1,
				CreatedAt:             &testTime,
			},
		},
		CreatedAt: &testTime,
		UpdatedAt: &testTime,
		DeletedAt: &testTime,
	}

	expect := &pb.Project{
		Id:      1,
		Name:    "project",
		RepoUrl: "repo",
		Ratings: []*pb.Rating{
			{
				Id:                    1,
				ProjectId:             1,
				ReliabilityRating:     1,
				MaintainabilityRating: 1,
				SecurityRating:        1,
				SecurityReviewRating:  1,
				Coverage:              1.0,
				Duplications:          1.0,
				Lines:                 1,
				CreatedAt:             testPbTime,
			},
			{
				Id:                    2,
				ProjectId:             1,
				ReliabilityRating:     1,
				MaintainabilityRating: 1,
				SecurityRating:        1,
				SecurityReviewRating:  1,
				Coverage:              1.0,
				Duplications:          1.0,
				Lines:                 1,
				CreatedAt:             testPbTime,
			},
		},
		CreatedAt: testPbTime,
		UpdatedAt: testPbTime,
		DeletedAt: testPbTime,
	}

	got := input.ToProto()
	require.EqualValues(t, expect, got)
}

func TestRatingToProto(t *testing.T) {
	testPbTime := ptypes.TimestampNow()
	testTime, err := ptypes.Timestamp(testPbTime)
	require.NoError(t, err)

	input := &Rating{
		ID:                    1,
		ReliabilityRating:     1,
		MaintainabilityRating: 1,
		SecurityRating:        1,
		SecurityReviewRating:  1,
		Coverage:              1.0,
		Duplications:          1.0,
		Lines:                 1,
		CreatedAt:             &testTime,
	}

	expect := &pb.Rating{
		Id:                    1,
		ReliabilityRating:     1,
		MaintainabilityRating: 1,
		SecurityRating:        1,
		SecurityReviewRating:  1,
		Coverage:              1.0,
		Duplications:          1.0,
		Lines:                 1,
		CreatedAt:             testPbTime,
	}

	got := input.ToProto()
	require.EqualValues(t, expect, got)
}

func TestCandidateProjectToProto(t *testing.T) {
	input := &CandidateProject{
		ID:          1,
		CandidateID: 1,
		ProjectID:   1,
	}

	expect := &pb.CandidateProject{
		Id:          1,
		CandidateId: 1,
		ProjectId:   1,
	}

	got := input.ToProto()
	require.EqualValues(t, expect, got)
}
