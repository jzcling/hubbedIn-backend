package models

import (
	"in-backend/services/joblisting/pb"
	"testing"

	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/require"
)

func TestJoblistingToORM(t *testing.T) {
	testPbTime := ptypes.TimestampNow()
	testTime, err := ptypes.Timestamp(testPbTime)
	require.NoError(t, err)

	input := &pb.Joblisting{
		Id:      1,
		Name:    "joblisting",
		RepoUrl: "repo",
		Ratings: []*pb.Rating{
			{
				Id:                    1,
				JoblistingId:          1,
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
				JoblistingId:          1,
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

	expect := &Joblisting{
		ID:      1,
		Name:    "joblisting",
		RepoURL: "repo",
		Ratings: []*Rating{
			{
				ID:                    1,
				JoblistingID:          1,
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
				JoblistingID:          1,
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

	got := JoblistingToORM(input)
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

func TestCandidateJoblistingToORM(t *testing.T) {
	input := &pb.CandidateJoblisting{
		Id:           1,
		CandidateId:  1,
		JoblistingId: 1,
	}

	expect := &CandidateJoblisting{
		ID:           1,
		CandidateID:  1,
		JoblistingID: 1,
	}

	got := CandidateJoblistingToORM(input)
	require.EqualValues(t, expect, got)
}
