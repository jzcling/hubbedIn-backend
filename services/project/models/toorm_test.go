package models

import (
	"in-backend/services/project/pb"
	"testing"

	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/require"
)

func TestProjectToORM(t *testing.T) {
	testPbTime := ptypes.TimestampNow()
	testTime, err := ptypes.Timestamp(testPbTime)
	require.NoError(t, err)

	input := &pb.Project{
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

	expect := &Project{
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

	got := ProjectToORM(input)
	require.EqualValues(t, expect, got)
}

func TestRatingToORM(t *testing.T) {
	testPbTime := ptypes.TimestampNow()
	testTime, err := ptypes.Timestamp(testPbTime)
	require.NoError(t, err)

	input := &pb.Rating{
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

	expect := &Rating{
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

	got := RatingToORM(input)
	require.EqualValues(t, expect, got)
}

func TestCandidateProjectToORM(t *testing.T) {
	input := &pb.CandidateProject{
		Id:          1,
		CandidateId: 1,
		ProjectId:   1,
	}

	expect := &CandidateProject{
		ID:          1,
		CandidateID: 1,
		ProjectID:   1,
	}

	got := CandidateProjectToORM(input)
	require.EqualValues(t, expect, got)
}
