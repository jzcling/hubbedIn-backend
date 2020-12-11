package models

import (
	"in-backend/services/assessment/pb"
	"testing"

	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/require"
)

func TestAssessmentToORM(t *testing.T) {
	testPbTime := ptypes.TimestampNow()
	testTime, err := ptypes.Timestamp(testPbTime)
	require.NoError(t, err)

	input := &pb.Assessment{
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
				ImageUrl:  "image",
				Options:   []string{"test", "test2"},
				Answer:    0,
			},
			{
				Id:        2,
				CreatedBy: 1,
				Type:      "Multiple Choice",
				Text:      "text",
				ImageUrl:  "image",
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

	expect := &Assessment{
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
				ImageURL:  "image",
				Options:   []string{"test", "test2"},
				Answer:    0,
			},
			{
				ID:        2,
				CreatedBy: 1,
				Type:      "Multiple Choice",
				Text:      "text",
				ImageURL:  "image",
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

	got := AssessmentToORM(input)
	require.EqualValues(t, expect, got)
}

func TestAssessmentAttemptToORM(t *testing.T) {
	testPbTime := ptypes.TimestampNow()
	testTime, err := ptypes.Timestamp(testPbTime)
	require.NoError(t, err)

	input := &pb.AssessmentAttempt{
		Id:           1,
		AssessmentId: 1,
		CandidateId:  1,
		Status:       "Completed",
		StartedAt:    testPbTime,
		CompletedAt:  testPbTime,
		Score:        5,
	}

	expect := &AssessmentAttempt{
		ID:           1,
		AssessmentID: 1,
		CandidateID:  1,
		Status:       "Completed",
		StartedAt:    &testTime,
		CompletedAt:  &testTime,
		Score:        5,
	}

	got := AssessmentAttemptToORM(input)
	require.EqualValues(t, expect, got)
}

func TestQuestionToORM(t *testing.T) {
	input := &pb.Question{
		Id:        1,
		CreatedBy: 1,
		Type:      "Open",
		Text:      "text",
		ImageUrl:  "image",
		Options:   []string{"test", "test2"},
		Answer:    0,
	}

	expect := &Question{
		ID:        1,
		CreatedBy: 1,
		Type:      "Open",
		Text:      "text",
		ImageURL:  "image",
		Options:   []string{"test", "test2"},
		Answer:    0,
	}

	got := QuestionToORM(input)
	require.EqualValues(t, expect, got)
}

func TestTagToORM(t *testing.T) {
	input := &pb.Tag{
		Id:   1,
		Name: "javascript",
	}

	expect := &Tag{
		ID:   1,
		Name: "javascript",
	}

	got := TagToORM(input)
	require.EqualValues(t, expect, got)
}

func TestQuestionTagToORM(t *testing.T) {
	input := &pb.QuestionTag{
		Id:         1,
		QuestionId: 1,
		TagId:      1,
	}

	expect := &QuestionTag{
		ID:         1,
		QuestionID: 1,
		TagID:      1,
	}

	got := QuestionTagToORM(input)
	require.EqualValues(t, expect, got)
}

func TestResponseToORM(t *testing.T) {
	testPbTime := ptypes.TimestampNow()
	testTime, err := ptypes.Timestamp(testPbTime)
	require.NoError(t, err)

	input := &pb.Response{
		Id:          1,
		QuestionId:  1,
		CandidateId: 1,
		Selection:   0,
		Text:        "text",
		Score:       0,
		TimeTaken:   10,
		CreatedAt:   testPbTime,
	}

	expect := &Response{
		ID:          1,
		QuestionID:  1,
		CandidateID: 1,
		Selection:   0,
		Text:        "text",
		Score:       0,
		TimeTaken:   10,
		CreatedAt:   &testTime,
	}

	got := ResponseToORM(input)
	require.EqualValues(t, expect, got)
}

func TestAssessmentQuestionToORM(t *testing.T) {
	input := &pb.AssessmentQuestion{
		Id:           1,
		AssessmentId: 1,
		QuestionId:   1,
	}

	expect := &AssessmentQuestion{
		ID:           1,
		AssessmentID: 1,
		QuestionID:   1,
	}

	got := AssessmentQuestionToORM(input)
	require.EqualValues(t, expect, got)
}
