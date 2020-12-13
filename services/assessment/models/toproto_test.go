package models

import (
	"in-backend/services/assessment/pb"
	"testing"

	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/require"
)

func TestAssessmentToProto(t *testing.T) {
	testPbTime := ptypes.TimestampNow()
	testTime, err := ptypes.Timestamp(testPbTime)
	require.NoError(t, err)

	input := &Assessment{
		ID:           1,
		Name:         "Javascript",
		Description:  "JS Test",
		Notes:        "Notes",
		ImageURL:     "image",
		Difficulty:   "Easy",
		TimeAllowed:  3600,
		Type:         "Multiple Choice",
		Randomise:    true,
		NumQuestions: 10,
		Questions: []*Question{
			{
				ID:        1,
				CreatedBy: 1,
				Type:      "Open",
				Text:      "text",
				MediaURL:  "image",
				Code:      "code",
				Options:   []string{"test", "test2"},
				Answer:    0,
			},
			{
				ID:        2,
				CreatedBy: 1,
				Type:      "Multiple Choice",
				Text:      "text",
				MediaURL:  "image",
				Code:      "code",
				Options:   []string{"test", "test2"},
				Answer:    0,
			},
		},
		Attempts: []*AssessmentAttempt{
			{
				ID:           1,
				AssessmentID: 1,
				CandidateID:  1,
				Status:       "Completed",
				StartedAt:    &testTime,
				CompletedAt:  &testTime,
				Score:        5,
			},
			{
				ID:           2,
				AssessmentID: 1,
				CandidateID:  1,
				Status:       "Completed",
				StartedAt:    &testTime,
				CompletedAt:  &testTime,
				Score:        5,
			},
		},
	}

	expect := &pb.Assessment{
		Id:           1,
		Name:         "Javascript",
		Description:  "JS Test",
		Notes:        "Notes",
		ImageUrl:     "image",
		Difficulty:   "Easy",
		TimeAllowed:  3600,
		Type:         "Multiple Choice",
		Randomise:    true,
		NumQuestions: 10,
		Questions: []*pb.Question{
			{
				Id:        1,
				CreatedBy: 1,
				Type:      "Open",
				Text:      "text",
				MediaUrl:  "image",
				Code:      "code",
				Options:   []string{"test", "test2"},
				Answer:    0,
			},
			{
				Id:        2,
				CreatedBy: 1,
				Type:      "Multiple Choice",
				Text:      "text",
				MediaUrl:  "image",
				Code:      "code",
				Options:   []string{"test", "test2"},
				Answer:    0,
			},
		},
		Attempts: []*pb.AssessmentAttempt{
			{
				Id:           1,
				AssessmentId: 1,
				CandidateId:  1,
				Status:       "Completed",
				StartedAt:    testPbTime,
				CompletedAt:  testPbTime,
				Score:        5,
			},
			{
				Id:           2,
				AssessmentId: 1,
				CandidateId:  1,
				Status:       "Completed",
				StartedAt:    testPbTime,
				CompletedAt:  testPbTime,
				Score:        5,
			},
		},
	}

	got := input.ToProto()
	require.EqualValues(t, expect, got)
}

func TestAssessmentAttemptToProto(t *testing.T) {
	testPbTime := ptypes.TimestampNow()
	testTime, err := ptypes.Timestamp(testPbTime)
	require.NoError(t, err)

	input := &AssessmentAttempt{
		ID:           1,
		AssessmentID: 1,
		CandidateID:  1,
		Status:       "Completed",
		StartedAt:    &testTime,
		CompletedAt:  &testTime,
		Score:        5,
	}

	expect := &pb.AssessmentAttempt{
		Id:           1,
		AssessmentId: 1,
		CandidateId:  1,
		Status:       "Completed",
		StartedAt:    testPbTime,
		CompletedAt:  testPbTime,
		Score:        5,
	}

	got := input.ToProto()
	require.EqualValues(t, expect, got)
}

func TestQuestionToProto(t *testing.T) {
	input := &Question{
		ID:        1,
		CreatedBy: 1,
		Type:      "Open",
		Text:      "text",
		MediaURL:  "image",
		Code:      "code",
		Options:   []string{"test", "test2"},
		Answer:    0,
	}

	expect := &pb.Question{
		Id:        1,
		CreatedBy: 1,
		Type:      "Open",
		Text:      "text",
		MediaUrl:  "image",
		Code:      "code",
		Options:   []string{"test", "test2"},
		Answer:    0,
	}

	got := input.ToProto()
	require.EqualValues(t, expect, got)
}

func TestTagToProto(t *testing.T) {
	input := &Tag{
		ID:   1,
		Name: "javascript",
	}

	expect := &pb.Tag{
		Id:   1,
		Name: "javascript",
	}

	got := input.ToProto()
	require.EqualValues(t, expect, got)
}

func TestQuestionTagToProto(t *testing.T) {
	input := &QuestionTag{
		ID:         1,
		QuestionID: 1,
		TagID:      1,
	}

	expect := &pb.QuestionTag{
		Id:         1,
		QuestionId: 1,
		TagId:      1,
	}

	got := input.ToProto()
	require.EqualValues(t, expect, got)
}

func TestAttemptQuestionToProto(t *testing.T) {
	testPbTime := ptypes.TimestampNow()
	testTime, err := ptypes.Timestamp(testPbTime)
	require.NoError(t, err)

	input := &AttemptQuestion{
		ID:          1,
		AttemptID:   1,
		QuestionID:  1,
		CandidateID: 1,
		Selection:   0,
		Text:        "text",
		CMMode:      "text/javascript",
		Score:       0,
		TimeTaken:   10,
		CreatedAt:   &testTime,
		UpdatedAt:   &testTime,
	}

	expect := &pb.AttemptQuestion{
		Id:          1,
		AttemptId:   1,
		QuestionId:  1,
		CandidateId: 1,
		Selection:   0,
		Text:        "text",
		CmMode:      "text/javascript",
		Score:       0,
		TimeTaken:   10,
		CreatedAt:   testPbTime,
		UpdatedAt:   testPbTime,
	}

	got := input.ToProto()
	require.EqualValues(t, expect, got)
}

func TestAssessmentQuestionToProto(t *testing.T) {
	input := &AssessmentQuestion{
		ID:           1,
		AssessmentID: 1,
		QuestionID:   1,
	}

	expect := &pb.AssessmentQuestion{
		Id:           1,
		AssessmentId: 1,
		QuestionId:   1,
	}

	got := input.ToProto()
	require.EqualValues(t, expect, got)
}
